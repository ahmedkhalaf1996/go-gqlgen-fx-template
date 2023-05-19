package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

	"github.com/dan6erbond/go-gqlgen-fx-template/graph/generated"
	"github.com/dan6erbond/go-gqlgen-fx-template/graph/model"
	"github.com/dan6erbond/go-gqlgen-fx-template/pkg/models"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*models.Todo, error) {
	// panic(fmt.Errorf("not implemented: CreateTodo - createTodo"))

	if len(input.Text) < 3 {
		return nil, errors.New("text not long enough")
	}

	if input.UserID == "" {
		return nil, errors.New("missing provide creator id")
	}

	todo := &models.Todo{
		Text:   input.Text,
		Done:   false,
		UserID: input.UserID,
	}

	r.db.Create(&todo)

	return todo, nil
}

// MarkComplete is the resolver for the markComplete field.
func (r *mutationResolver) MarkComplete(ctx context.Context, todoID string) (*models.Todo, error) {
	if todoID == "" {
		return nil, errors.New("missing provide todo id")
	}

	todo := models.Todo{}
	err := r.db.First(&todo, "id = ?", todoID).Error

	if err != nil {
		return nil, errors.New("todo not found")
	}

	newTodo := todo
	newTodo.Done = true
	r.db.Model(&todo).Updates(&newTodo)

	return &todo, nil
}

// DeleteTodo is the resolver for the deleteTodo field.
func (r *mutationResolver) DeleteTodo(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteTodo - deleteTodo"))
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*models.Todo, error) {
	// _, err := middleware.GetCurrentUserFromCTX(ctx)
	// if err != nil {
	// 	return nil, errors.New("not authenticated fine")
	// }
	var todos []*models.Todo
	r.db.Find(&todos)
	return todos, nil
}

// ID is the resolver for the id field.
func (r *todoResolver) ID(ctx context.Context, obj *models.Todo) (string, error) {
	// panic(fmt.Errorf("not implemented: ID - id"))
	return fmt.Sprint(obj.ID), nil
}

// User is the resolver for the user field.
func (r *todoResolver) User(ctx context.Context, obj *models.Todo) (*models.User, error) {
	// TODO DATALOADER
	var user models.User
	err := r.db.First(&user, "id = ?", obj.UserID).Error
	return &user, err
}

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type todoResolver struct{ *Resolver }
