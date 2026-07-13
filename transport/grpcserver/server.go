package grpcserver

import (
	loansv1 "loans-item-go/gen/loans/v1"
	loansvc "loans-item-go/service/loan"
	usersvc "loans-item-go/service/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func New(jwtSecret string, userSvc usersvc.Service, loanSvc loansvc.Service) *grpc.Server {
	server := grpc.NewServer(
		// Recovery must be outermost: services panic on error and the
		// auth interceptor runs inside it.
		grpc.ChainUnaryInterceptor(
			RecoveryInterceptor(),
			AuthInterceptor(jwtSecret),
		),
	)

	loansv1.RegisterUserServiceServer(server, NewUserHandler(userSvc))
	loansv1.RegisterLoanServiceServer(server, NewLoanHandler(loanSvc))

	// Reflection lets grpcurl/Postman discover services without .proto files.
	reflection.Register(server)

	return server
}
