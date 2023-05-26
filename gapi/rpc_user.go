package gapi

import (
	"context"

	"github.com/azusaanson/invest-api/domain"
	"github.com/azusaanson/invest-api/proto/pb"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	if violations := validateCreateUserRequest(req); violations != nil {
		return nil, invalidArgumentError(violations)
	}

	name, _ := domain.NewUserName(req.GetName())
	password, _ := domain.NewPassword(req.GetPassword())
	role, _ := domain.NewUserRole(req.GetRole())

	userExist, err := server.store.GetUserByName(ctx, name)
	if err != nil {
		return nil, errorWithCode(codes.Internal, err)
	}
	if userExist != nil {
		return nil, errorWithCode(codes.AlreadyExists, ErrDuplicateUserName)
	}

	user, err := domain.NewUser(name, password.Hash(), role)
	if err != nil {
		return nil, errorWithCode(codes.Internal, err)
	}

	if err := server.store.CreateUser(ctx, user); err != nil {
		return nil, errorWithCode(codes.Internal, err)
	}

	res := &pb.CreateUserResponse{
		User: toUserResponse(user),
	}
	return res, nil
}

func validateCreateUserRequest(req *pb.CreateUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if req.GetName() == "" {
		violations = append(violations, fieldViolation("name", ErrValidationUserNameRequired))
	} else if _, err := domain.NewUserName(req.GetName()); err != nil {
		violations = append(violations, fieldViolation("name", err))
	}

	if req.GetPassword() == "" {
		violations = append(violations, fieldViolation("password", ErrValidationUserPasswordRequired))
	} else if _, err := domain.NewPassword(req.GetPassword()); err != nil {
		violations = append(violations, fieldViolation("password", err))
	}

	if req.GetRole() == "" {
		violations = append(violations, fieldViolation("role", ErrValidationUserRoleRequired))
	} else if _, err := domain.NewUserRole(req.GetRole()); err != nil {
		violations = append(violations, fieldViolation("role", err))
	}

	return violations
}

func (server *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	res := &pb.LoginResponse{}
	return res, nil
}

func toUserResponse(user *domain.User) *pb.User {
	return &pb.User{
		Name: string(user.Name()),
		Role: string(user.Role()),
	}
}
