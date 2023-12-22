package services

import (
	"context"
	"study-kafka-ddb/controllers/dtos"
	"study-kafka-ddb/domains"
	"study-kafka-ddb/domains/interfaces"
	"study-kafka-ddb/utils/deftype"
)

type UserService struct {
	userRepo interfaces.UserRepository
}

func (r *UserService) SignUp(ctx context.Context, req *dtos.SignUpReq) (*dtos.SignUpResp, deftype.Error) {
	user := domains.NewUser()

	if err := user.SignUp(req.Name, req.Email, req.Password); err != nil {
		return nil, err
	}

	if err := r.userRepo.Save(ctx, user); err != nil {
		return nil, err
	}

	return &dtos.SignUpResp{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		Email:     user.Email,
	}, nil
}

func (r *UserService) SignIn(ctx context.Context, req *dtos.SignUpReq) (*dtos.SignInResp, deftype.Error) {
	if !domains.ValidateUserInput(req.Email, req.Password) {
		return nil, deftype.ErrInvalidRequestData
	}

	user, err := r.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if err = user.SignIn(req.Email, req.Password); err != nil {
		return nil, err
	}

	return &dtos.SignInResp{
		Success: true,
	}, nil
}
