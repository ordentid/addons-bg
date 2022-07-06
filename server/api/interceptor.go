package api

import (
	"os"

	jwt "github.com/dgrijalva/jwt-go"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"context"
	"time"

	"bitbucket.bri.co.id/scm/addons/addons-bg-service/server/db"
)

// AuthTokenKey is the key used within our metadata to store a JWT
const AuthTokenKey = "authentication"

var authFilteredModule = []string{
	apiServicePath + "GetUser",
}

var (
	errGrpcUnauthenticated = status.Errorf(codes.Unauthenticated, "missing authentication token")
)

type UserCtxKey struct{}

// Interceptors implements the grpc.UnaryServerInteceptor function to add
// interceptors around all gRPC unary calls
func UnaryInterceptors(
	authI *AuthInterceptor,
) grpc.UnaryServerInterceptor {
	return grpc_middleware.ChainUnaryServer(
		LoggingInterceptor,
		ErrorsInterceptor,
		authI.Unary(),
	)
}

// Interceptors implements the grpc.StreamServerInteceptor function to add
// interceptors around all gRPC stream calls
func StreamInterceptors(
	authI *AuthInterceptor,
) grpc.StreamServerInterceptor {
	return grpc_middleware.ChainStreamServer(
		authI.Stream(),
	)
}

// ErrorsInterceptor adds error type checking to see if there are any known types
// what we return different grpc error codes for, for example: NotFound resources.
func ErrorsInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (out interface{}, err error) {
	out, err = handler(ctx, req)

	switch tErr := err.(type) {
	case db.NotFoundErr:
		return out, status.Errorf(codes.NotFound, tErr.Error())
	}

	return out, err
}

// LoggingInterceptor adds logging around every gRPC call. It includes the method name and timing information.
// if the given handler raises an error, it also appends that to a key.
func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (out interface{}, err error) {
	if info.FullMethod == "/grpc.health.v1.Health/Check" {
		out, err = handler(ctx, req)

		return out, err
	} else {
		entry := logrus.WithField("method", info.FullMethod)
		start := time.Now()
		out, err = handler(ctx, req)
		duration := time.Since(start)

		if err != nil {
			entry = entry.WithError(err)
		}

		entry.WithField("duration", duration.String()).Info("finished RPC")
		return out, err
	}
}

// AuthenticationInterceptor validates a JWT token and appends the username to the
// context that is passed to the handler
func AuthenticationInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (out interface{}, err error) {

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return handler(ctx, req)
	}

	if restricted := contains(authFilteredModule, info.FullMethod); !restricted {
		return handler(ctx, req)
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errGrpcUnauthenticated
	}

	tokenString, ok := md[AuthTokenKey]
	if !ok || len(tokenString) < 1 {
		return nil, errGrpcUnauthenticated
	}

	token, err := jwt.Parse(tokenString[0], func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		ctx = context.WithValue(ctx, UserCtxKey{}, claims["un"])
	} else {
		return nil, errGrpcUnauthenticated
	}

	return handler(ctx, req)
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
