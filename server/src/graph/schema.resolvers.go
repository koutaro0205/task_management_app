package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"server/db"
	"server/graph/generated"
	"server/graph/model"
	"server/graph/service"
)

// Mutation
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.CreateTodoInput) (*model.Todo, error) {
	return service.CreateTodo(db.DB, input)
}

// UpdateTodo is the resolver for the updateTodo field.
func (r *mutationResolver) UpdateTodo(ctx context.Context, input model.UpdateTodoInput) (*model.Todo, error) {
	return service.UpdateTodo(db.DB, input)
}

// DeleteTodo is the resolver for the deleteTodo field.
func (r *mutationResolver) DeleteTodo(ctx context.Context, input model.DeleteTodoInput) (*model.Todo, error) {
	return service.DeleteTodo(db.DB, input.ID)
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	return service.CreateUser(db.DB, input)
}

// Query
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return service.GetAllTodos(db.DB)
}

// Todo is the resolver for the todo field.
func (r *queryResolver) Todo(ctx context.Context, input *model.FetchTodoInput) (*model.Todo, error) {
	return service.GetTodo(db.DB, input.ID)
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return service.GetAllUsers(db.DB)
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, input *model.FetchUserInput) (*model.User, error) {
	return service.GetUser(db.DB, input.ID)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }