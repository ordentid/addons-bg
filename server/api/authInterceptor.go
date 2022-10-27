package api

import (
	"context"
	"fmt"
	"strings"

	manager "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/lib/jwt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type AuthInterceptor struct {
	jwtManager      *manager.JWTManager
	accessibleRoles map[string][]string
}

func NewAuthInterceptor(jwtManager *manager.JWTManager) *AuthInterceptor {
	return &AuthInterceptor{jwtManager, accessibleRoles()}
}

// Filter access by role
func accessibleRoles() map[string][]string {

	// restricted api
	return map[string][]string{
		// HTTP - API
		apiServicePath + "CreateTaskMapping":        {"data_entry:maker"},
		apiServicePath + "CreateTaskMappingDigital": {"data_entry:maker"},
		// apiServicePath + "CreateTransaction":        {"data_entry:maker"},
		// apiServicePath + "DeleteTransaction":        {"data_entry:maker"},
		apiServicePath + "CreateTaskIssuing":        {"data_entry:maker"},
		apiServicePath + "GetTaskIssuingFile":       {"download_report:-"},
	}
}

func (interceptor *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {

		if !interceptor.isRestricted(info.FullMethod) {
			return handler(ctx, req)
		}

		claims, err := interceptor.claimsToken(ctx)
		if err != nil {
			return nil, err
		}

		// ctx, err = interceptor.getUserData(ctx)
		// if err != nil {
		// 	return nil, err
		// }

		err = interceptor.authorize(ctx, claims, info.FullMethod)
		if err != nil {
			return nil, err
		}

		return handler(ctx, req)
	}
}

func (interceptor *AuthInterceptor) Stream() grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		stream grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {

		if !interceptor.isRestricted(info.FullMethod) {
			return handler(srv, stream)
		}

		claims, err := interceptor.claimsToken(stream.Context())
		if err != nil {
			return err
		}

		err = interceptor.authorize(stream.Context(), claims, info.FullMethod)
		if err != nil {
			return err
		}

		return handler(srv, stream)
	}
}

func (interceptor *AuthInterceptor) isRestricted(method string) bool {
	_, restricted := interceptor.accessibleRoles[method]
	return restricted
}

func (interceptor *AuthInterceptor) authorize(ctx context.Context, claims *manager.UserClaims, method string) error {
	// fmt.Println(md)
	featureRoles := []string{}
	for _, v := range claims.ProductRoles {
		if contains([]string{"BG Mapping", "BG Mapping Digital", "BG Monitoring", "BG Issuing"}, v.ProductName) {
			featureRoles = v.Authorities
			break
		}
	}
	logrus.Infoln("[interceptor] Feature Roles:", featureRoles)

	accessibleRoles, ok := interceptor.accessibleRoles[method]
	if !ok {
		// everyone can access
		return nil
	}

	if len(accessibleRoles) < 1 {
		return nil
	}

	for _, role := range accessibleRoles {
		for _, exist := range featureRoles {
			if role == exist {
				return nil
			}
		}
	}

	return status.Error(codes.PermissionDenied, "Access denied")
}

func (interceptor *AuthInterceptor) claimsToken(ctx context.Context) (*manager.UserClaims, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	split := strings.Split(values[0], " ")
	accessToken := split[0]
	if len(split) > 1 {
		accessToken = split[1]
	}
	claims, err := interceptor.jwtManager.Verify(accessToken)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	getUser, err := interceptor.jwtManager.GetMeFromAuthService(ctx, accessToken)
	if err != nil {
		return nil, err
	}
	if getUser.IsExpired && !getUser.IsValid {
		return nil, status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	return claims, nil
}

func (interceptor *AuthInterceptor) getUserData(ctx context.Context) (context.Context, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	split := strings.Split(values[0], " ")
	accessToken := split[0]
	if len(split) > 1 {
		accessToken = split[1]
	}

	userType := md["auth-usertype"]
	username := md["auth-username"]
	userid := md["auth-userid"]

	if len(userType) == 0 || len(username) == 0 || len(userid) == 0 {
		getUser, err := interceptor.jwtManager.GetMeFromAuthService(ctx, accessToken)
		if err != nil {
			return nil, err
		}
		if getUser.IsExpired && !getUser.IsValid {
			return nil, status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
		}

		for _, v := range getUser.ProductRoles {
			grpc.SendHeader(ctx, metadata.Pairs("auth-role-"+v.ProductName, strings.Join(v.Authorities, "|")))
		}

		grpc.SendHeader(ctx, metadata.Pairs("auth-usertype", getUser.UserType))
		grpc.SendHeader(ctx, metadata.Pairs("auth-username", getUser.Username))
		grpc.SendHeader(ctx, metadata.Pairs("auth-userid", fmt.Sprintf("%v", getUser.UserID)))

		fmt.Println("")
		fmt.Println("=====>")
		fmt.Println(getUser.Username)

	}

	return ctx, nil
}
