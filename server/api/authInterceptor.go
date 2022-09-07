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
	jwtManager            *manager.JWTManager
	accessibleAbonnements map[string][]string
}

func NewAuthInterceptor(jwtManager *manager.JWTManager) *AuthInterceptor {
	return &AuthInterceptor{jwtManager, accessibleAbonnements()}
}

// Filter access by abonnement
func accessibleAbonnements() map[string][]string {

	// restricted api
	return map[string][]string{
		apiServicePath + "CreateUser": {"admin"},
		apiServicePath + "GetMe":      {"user", "admin"},

		apiServicePath + "GetCurrency":                 {},
		apiServicePath + "GetBeneficiaryName":          {},
		apiServicePath + "GetApplicantName":            {},
		apiServicePath + "GetThirdParty":               {},
		apiServicePath + "GetCustomerLimit":            {},
		apiServicePath + "GetTaskMappingFile":          {},
		apiServicePath + "GetTaskMapping":              {},
		apiServicePath + "GetTaskMappingDetail":        {},
		apiServicePath + "CreateTaskMapping":           {"data_entry:maker"},
		apiServicePath + "GetTaskMappingDigitalFile":   {"download_report:-"},
		apiServicePath + "GetTaskMappingDigital":       {},
		apiServicePath + "GetTaskMappingDigitalDetail": {},
		apiServicePath + "CreateTaskMappingDigital":    {"data_entry:maker"},
		apiServicePath + "GetTransactionAttachment":    {},
		apiServicePath + "GetTransactionFile":          {},
		apiServicePath + "GetTransaction":              {},
		apiServicePath + "GetTransactionDetail":        {},
		apiServicePath + "CreateTransaction":           {"data_entry:maker"},
		apiServicePath + "DeleteTransaction":           {"delete:maker"},
		apiServicePath + "GetTaskIssuing":              {},
		apiServicePath + "GetTaskIssuingDetail":        {},
		apiServicePath + "GetTaskIssuingFile":          {"download_report:-"},
		apiServicePath + "CreateTaskIssuing":           {"data_entry:maker"},
		apiServicePath + "TaskAction":                  {"approve:signer"},
		apiServicePath + "CheckIssuingStatus":          {},
		apiServicePath + "FileUpload":                  {"upload_file:maker"},
		apiServicePath + "CheckIndividualLimit":        {},
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

		err = interceptor.authorize(claims, info.FullMethod)
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

		err = interceptor.authorize(claims, info.FullMethod)
		if err != nil {
			return err
		}

		return handler(srv, stream)
	}
}

func (interceptor *AuthInterceptor) isRestricted(method string) bool {
	_, restricted := interceptor.accessibleAbonnements[method]
	return restricted
}

func (interceptor *AuthInterceptor) authorize(claims *manager.UserClaims, method string) error {
	// fmt.Println(md)
	featureRoles := []string{}
	for _, v := range claims.ProductRoles {
		if v.ProductName == "Subscription" {
			featureRoles = v.Authorities
			break
		}
	}

	accessibleRoles, ok := interceptor.accessibleAbonnements[method]
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

	return status.Error(codes.PermissionDenied, "no permission to access this RPC")
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
	return claims, nil
}
