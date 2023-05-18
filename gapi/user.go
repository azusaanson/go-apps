package gapi

import (
	"context"

	"github.com/azusaanson/invest-api/domain"
	"github.com/azusaanson/invest-api/proto/pb"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	if err := validateCreateUserRequest(req); err != nil {
		return nil, err
	}

	name, _ := domain.NewUserName(req.GetName())
	password, _ := domain.NewPassword(req.GetPassword())
	role, _ := domain.NewUserRole(req.GetRole())

	user, err := domain.NewUser(name, password.Hash(), role)
	if err != nil {
		return nil, err
	}

	if err := server.store.CreateUser(ctx, user); err != nil {
		return nil, err
	}

	res := &pb.CreateUserResponse{
		User: toUserResponse(user),
	}
	return res, nil
}

func validateCreateUserRequest(req *pb.CreateUserRequest) error {
	if req.GetName() == "" {
		return ErrValidationUserNameRequired
	} else if _, err := domain.NewUserName(req.GetName()); err != nil {
		return err
	}

	if req.GetPassword() == "" {
		return ErrValidationUserPasswordRequired
	} else if _, err := domain.NewPassword(req.GetPassword()); err != nil {
		return err
	}

	if req.GetRole() == "" {
		return ErrValidationUserRoleRequired
	} else if _, err := domain.NewUserRole(req.GetRole()); err != nil {
		return err
	}

	return nil
}

func toUserResponse(user *domain.User) *pb.User {
	return &pb.User{
		Name: string(user.Name()),
		Role: string(user.Role()),
	}
}
