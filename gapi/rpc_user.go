package gapi

import (
	"context"
	"time"

	"github.com/azusaanson/invest-api/domain"
	"github.com/azusaanson/invest-api/proto/pb"
	"github.com/pkg/errors"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	if violations := validateLoginRequest(req); violations != nil {
		return nil, invalidArgumentError(violations)
	}

	name, _ := domain.NewUserName(req.GetName())
	password, _ := domain.NewPassword(req.GetPassword())

	user, err := server.store.GetUserByName(ctx, name)
	if err != nil {
		return nil, serverError(err)
	}
	if user == nil {
		return nil, clientError(codes.NotFound, ErrNotFoundUser)
	}

	if err := user.HashedPassword().Verify(password); err != nil {
		return nil, clientError(codes.Unauthenticated, ErrValidationUserPasswordInvalid)
	}

	accessToken, accessPayload, err := server.tokenMaker.CreateToken(
		user.ID(),
		server.config.AccessTokenDuration,
	)
	if err != nil {
		return nil, serverError(errors.Wrap(ErrCreateAccessToken, err.Error()))
	}

	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(
		user.ID(),
		server.config.RefreshTokenDuration,
	)
	if err != nil {
		return nil, serverError(errors.Wrap(ErrCreateRefreshToken, err.Error()))
	}

	userMetaData, err := server.extractMetadata(ctx)
	if err != nil {
		return nil, serverError(err)
	}

	session, err := domain.NewSession(refreshPayload.ID, user.ID(), refreshToken, refreshPayload.ExpiresAt, userMetaData)
	if err != nil {
		return nil, serverError(err)
	}

	if err = server.store.CreateSession(ctx, session); err != nil {
		return nil, serverError(err)
	}

	res := &pb.LoginResponse{
		User:                  toUserResponse(user),
		SessionId:             session.UUID().ToString(),
		AccessToken:           string(accessToken),
		RefreshToken:          string(refreshToken),
		AccessTokenExpiresAt:  timestamppb.New(time.Time(accessPayload.ExpiresAt)),
		RefreshTokenExpiresAt: timestamppb.New(time.Time(refreshPayload.ExpiresAt)),
	}
	return res, nil
}

func validateLoginRequest(req *pb.LoginRequest) (violations []*errdetails.BadRequest_FieldViolation) {
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

	return violations
}

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	if violations := validateCreateUserRequest(req); violations != nil {
		return nil, invalidArgumentError(violations)
	}

	name, _ := domain.NewUserName(req.GetName())
	password, _ := domain.NewPassword(req.GetPassword())
	role, _ := domain.NewUserRole(req.GetRole())

	userExist, err := server.store.GetUserByName(ctx, name)
	if err != nil {
		return nil, serverError(err)
	}
	if userExist != nil {
		return nil, clientError(codes.AlreadyExists, ErrDuplicateUserName)
	}

	user, err := domain.NewUser(name, password.Hash(), role)
	if err != nil {
		return nil, serverError(err)
	}

	if err = server.store.ExecTx(ctx, func(ctx context.Context) error {
		if err := server.store.CreateUser(ctx, user); err != nil {
			return errors.WithStack(err)
		}

		return nil
	}); err != nil {
		return nil, serverError(err)
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

func toUserResponse(user *domain.User) *pb.User {
	return &pb.User{
		Name: string(user.Name()),
		Role: string(user.Role()),
	}
}
