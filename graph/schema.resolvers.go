package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"graphql_datnd/graph/generated"
	"graphql_datnd/graph/model"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		ID:   fmt.Sprint(len(r.TodosList) + 1),
		Text: input.Text,
		Done: input.Done,
		User: &model.User{
			ID:   input.UserID,
			Name: fmt.Sprintf("Hi%s", fmt.Sprint(len(r.TodosList)+1)),
		},
	}
	r.TodosList = append(r.TodosList, todo)
	return todo, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context, done bool) ([]*model.Todo, error) {
	return r.TodosList, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
