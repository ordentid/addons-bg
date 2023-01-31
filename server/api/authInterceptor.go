package api

import (
	"context"
	"strings"

	manager "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/lib/jwt"
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
		// ==== Example =====
		// apiServicePath + <api function name>: {<authority>, <productName>},

		// HTTP - API
		apiServicePath + "CreateTaskMapping":        {"data_entry:maker", "BG Mapping"},
		apiServicePath + "CreateTaskMappingDigital": {"data_entry:maker", "BG Mapping Digital"},
		// apiServicePath + "CreateTransaction":        {"data_entry:maker"},
		// apiServicePath + "DeleteTransaction":        {"data_entry:maker"},
		apiServicePath + "CreateTaskIssuing":  {"data_entry:maker", "BG Issuing"},
		apiServicePath + "GetTaskIssuingFile": {"download_report:-", "BG Issuing"},
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
	accessibleRoles, ok := interceptor.accessibleRoles[method]
	if !ok {
		// everyone can access
		return nil
	}
	log.Infoln("[interceptor] Accessible Roles:", accessibleRoles)

	if len(accessibleRoles) < 1 {
		return nil
	}

	allowedAccess := accessibleRoles[0]
	currentProduct := accessibleRoles[1]

	featureRoles := []string{}
	for _, v := range claims.ProductRoles {
		if currentProduct == v.ProductName {
			featureRoles = v.Authorities
			break
		}
	}
	log.Infoln("[interceptor] Feature Roles:", featureRoles)

	for _, exist := range featureRoles {
		if allowedAccess == exist {
			return nil
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
