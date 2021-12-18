package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	controller "github.com/fadhiilrachman/e-bpm/controllers"
	"github.com/fadhiilrachman/e-bpm/graph/generated"
	"github.com/fadhiilrachman/e-bpm/graph/model"
	auth "github.com/fadhiilrachman/e-bpm/middleware"
	"github.com/fadhiilrachman/e-bpm/utils"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	hashed, _ := utils.HashPassword(input.Password)

	user := &model.User{
		ID:          uuid.New().String(),
		Username:    input.Username,
		Password:    hashed,
		Fullname:    input.Fullname,
		Role:        input.Role,
		Phone:       input.Phone,
		Email:       input.Email,
		CreatedAt:   utils.TimeNow(),
		UpdatedAt:   utils.TimeNow(),
		LastLoginAt: utils.TimeNow(),
	}

	err := controller.CreateNewUser(user)

	if err != nil {
		return &model.User{}, fmt.Errorf("User already exist")
	}

	return user, nil
}

func (r *mutationResolver) LoginUser(ctx context.Context, input model.LoginUser) (string, error) {
	data, check := controller.AuthUser(input.Username, input.Password)
	if !check {
		return "", fmt.Errorf("Wrong username or password")
	}

	token, err := utils.GenerateToken(data.Role, data.Username)
	if err != nil {
		panic(err)
	}

	return token, nil
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenData) (string, error) {
	data, err := utils.ParseToken(input.Token)
	if err != nil {
		panic(err)
	}

	token, err := utils.GenerateToken(data.Role, data.Username)
	if err != nil {
		panic(err)
	}

	return token, nil
}

func (r *queryResolver) ParseTokenData(ctx context.Context) (*model.TokenData, error) {
	user, ok := auth.ForContext(ctx)

	if !ok {
		return &model.TokenData{}, fmt.Errorf("Forbidden")
	}

	data := &model.TokenData{
		Role:     user.Role,
		Username: user.Username,
	}

	return data, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
