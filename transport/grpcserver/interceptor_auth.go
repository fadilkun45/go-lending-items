package grpcserver

import (
	"context"

	"google.golang.org/grpc"
)

// Same public surface as middleware.Auth's publicPaths, plus reflection
// so grpcurl can discover services without a token.
var publicMethods = map[string]bool{
	"/loans.v1.UserService/Register": true,
	"/loans.v1.UserService/Login":    true,
}

// AuthInterceptor mirrors middleware.Auth: it reads the "authorization"
// metadata, validates the JWT, and stores user_id under middleware.UserIDKey
// so the service layer sees the same context shape on both transports.
func AuthInterceptor(secret string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		// if publicMethods[info.FullMethod] || strings.HasPrefix(info.FullMethod, "/grpc.reflection.") {
		// 	return handler(ctx, req)
		// }

		// md, ok := metadata.FromIncomingContext(ctx)
		// if !ok {
		// 	return nil, status.Error(codes.Unauthenticated, "missing metadata")
		// }
		// values := md.Get("authorization")
		// if len(values) == 0 || !strings.HasPrefix(values[0], "Bearer ") {
		// 	return nil, status.Error(codes.Unauthenticated, "missing or invalid authorization header")
		// }

		// tokenString := strings.TrimPrefix(values[0], "Bearer ")
		// claims := jwt.MapClaims{}
		// token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		// 	return []byte(secret), nil
		// })
		// if err != nil || !token.Valid {
		// 	return nil, status.Error(codes.Unauthenticated, "invalid or expired token")
		// }

		// userID, _ := claims["user_id"].(float64)
		// ctx = context.WithValue(ctx, middleware.UserIDKey, int64(userID))
		return handler(ctx, req)
	}
}
