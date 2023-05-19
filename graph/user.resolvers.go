package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/dan6erbond/go-gqlgen-fx-template/graph/generated"
	"github.com/dan6erbond/go-gqlgen-fx-template/graph/model"

	// "github.com/dan6erbond/go-gqlgen-fx-template/pkg"
	"github.com/dan6erbond/go-gqlgen-fx-template/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

// SignUp is the resolver for the signUp field.
func (r *mutationResolver) SignUp(ctx context.Context, input model.RegisterInput) (*model.AuthResponse, error) {
	var getuser models.User
	err := r.db.First(&getuser, "email = ?", input.Email).Error
	if err == nil {
		return nil, errors.New("email already in used")
	}

	// hashing
	bytePassword := []byte(input.Password)
	passwordHash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("can't hashed pass something went wrong")
	}

	hashedpassword := string(passwordHash)
	//
	user := &models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: hashedpassword,
	}

	err = r.db.Create(&user).Error
	if err != nil {
		return nil, errors.New("can't crate the user")
	}
	// generate token
	token, err := GenToken(user)
	if err != nil {
		log.Printf("error while generating the token: %v", err)
		return nil, errors.New("something went wrong")
	}
	//
	return &model.AuthResponse{
		// AuthToken: token,
		Authtoken: token,
		User:      user,
	}, nil

}

// SignIn is the resolver for the signIn field.
func (r *mutationResolver) SignIn(ctx context.Context, input model.LoginInput) (*model.AuthResponse, error) {
	// panic(fmt.Errorf("not implemented: SignIn - signIn"))
	var user models.User
	err := r.db.First(&user, "email = ?", input.Email).Error
	if err != nil {
		return nil, errors.New("email not found")
	}

	// compear password
	bytePassword := []byte(input.Password)
	byteHashedPassword := []byte(user.Password)
	err = bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
	if err != nil {
		return nil, errors.New("email/user not correct")
	}

	// generate token
	token, err := GenToken(&user)
	if err != nil {
		log.Printf("error while generating the token: %v", err)
		return nil, errors.New("something went wrong")
	}
	//
	return &model.AuthResponse{
		// AuthToken: token,
		Authtoken: token,
		User:      &user,
	}, nil

}

// Me is the resolver for the me field.
func (r *queryResolver) Me(ctx context.Context) (*models.User, error) {
	panic(fmt.Errorf("not implemented: Me - me"))
	// e, err := pkg.GetCurrentUserFromCTX(ctx)
	// if err != nil {
	// 	return nil, errors.New("not authenticated fine")
	// }
	// return e, nil
}

// ID is the resolver for the id field.
func (r *userResolver) ID(ctx context.Context, obj *models.User) (string, error) {
	// panic(fmt.Errorf("not implemented: ID - id"))
	return fmt.Sprint(obj.ID), nil
}

// Todos is the resolver for the todos field.
func (r *userResolver) Todos(ctx context.Context, obj *models.User) ([]*models.Todo, error) {
	// panic(fmt.Errorf("not implemented: Todos - todos"))
	// TODO DATALOADER
	var todos []*models.Todo
	r.db.Find(&todos, "user_id = ?", obj.ID).Select("password")
	// err := r.db.First(&todos, "id = ?", obj.ID).Error
	return todos, nil
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
