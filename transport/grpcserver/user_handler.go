package grpcserver

import (
	"context"

	loansv1 "loans-item-go/gen/loans/v1"
	"loans-item-go/model"
	usersvc "loans-item-go/service/user"
)

type UserHandler struct {
	loansv1.UnimplementedUserServiceServer
	svc usersvc.Service
}

func NewUserHandler(svc usersvc.Service) *UserHandler {
	return &UserHandler{svc: svc}
}

func (h *UserHandler) Register(ctx context.Context, req *loansv1.RegisterRequest) (*loansv1.RegisterResponse, error) {
	user := h.svc.Register(ctx, model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})
	return &loansv1.RegisterResponse{User: toProtoUser(user)}, nil
}

func (h *UserHandler) Login(ctx context.Context, req *loansv1.LoginRequest) (*loansv1.LoginResponse, error) {
	user, token := h.svc.Login(ctx, req.Email, req.Password)
	return &loansv1.LoginResponse{User: toProtoUser(user), Token: token}, nil
}

func (h *UserHandler) GetUser(ctx context.Context, req *loansv1.GetUserRequest) (*loansv1.GetUserResponse, error) {
	user := h.svc.FindById(ctx, int(req.Id))
	return &loansv1.GetUserResponse{User: toProtoUser(user)}, nil
}

func (h *UserHandler) ListUsers(ctx context.Context, req *loansv1.ListUsersRequest) (*loansv1.ListUsersResponse, error) {
	page, pageSize := pageOrDefault(req.Page)
	users, total := h.svc.FindAll(ctx, page, pageSize)
	return &loansv1.ListUsersResponse{
		Users:    toProtoUsers(users),
		Page:     int32(page),
		PageSize: int32(pageSize),
		Total:    total,
	}, nil
}

func (h *UserHandler) UpdateUser(ctx context.Context, req *loansv1.UpdateUserRequest) (*loansv1.UpdateUserResponse, error) {
	user := h.svc.Update(ctx, model.User{
		Id:    req.Id,
		Name:  req.Name,
		Email: req.Email,
	})
	return &loansv1.UpdateUserResponse{User: toProtoUser(user)}, nil
}

func (h *UserHandler) DeleteUser(ctx context.Context, req *loansv1.DeleteUserRequest) (*loansv1.DeleteUserResponse, error) {
	h.svc.Delete(ctx, model.User{Id: req.Id})
	return &loansv1.DeleteUserResponse{Message: "user deleted"}, nil
}
