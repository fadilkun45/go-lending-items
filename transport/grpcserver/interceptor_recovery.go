package grpcserver

import (
	"context"
	"loans-item-go/helper"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// RecoveryInterceptor is the gRPC counterpart of helper.RecoverError:
// services panic with helper.AppError, so this must be the outermost
// interceptor in the chain to translate panics into status codes.
func RecoveryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		defer func() {
			if r := recover(); r != nil {
				switch v := r.(type) {
				case helper.AppError:
					err = status.Error(httpToGRPCCode(v.Code), v.Message)
				case error:
					err = status.Error(codes.Internal, v.Error())
				case string:
					err = status.Error(codes.Internal, v)
				default:
					err = status.Error(codes.Internal, "internal server error")
				}
			}
		}()
		return handler(ctx, req)
	}
}

func httpToGRPCCode(code int) codes.Code {
	switch code {
	case http.StatusBadRequest:
		return codes.InvalidArgument
	case http.StatusUnauthorized:
		return codes.Unauthenticated
	case http.StatusNotFound:
		return codes.NotFound
	default:
		return codes.Internal
	}
}
