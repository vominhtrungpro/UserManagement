package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"fmt"
	"strconv"
	"vominhtrungpro/usermanagement/graph/model"
	"vominhtrungpro/usermanagement/internal/controller/user"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented: CreateTodo - createTodo"))
}

// Createuser is the resolver for the createuser field.
func (r *mutationResolver) Createuser(ctx context.Context, input model.NewUser) (*model.User, error) {
	var result model.User
	var request user.CreateInput
	request.Username = input.Username
	request.Password = input.Password
	request.Email = input.Email
	request.Age = int64(input.Age)
	user, err := r.UserCtrl.CreateUser(ctx, request)
	if err != nil {
		return &result, err
	}
	s := strconv.Itoa(int(user.ID))
	result.ID = s
	result.Username = user.Username
	result.Password = user.Password
	result.Email = user.Email
	result.Age = int(user.Age)
	return &result, nil
}

// Updateuser is the resolver for the updateuser field.
func (r *mutationResolver) Updateuser(ctx context.Context, input model.UpdateUser) (*model.User, error) {
	var result model.User
	var request user.UpdateInput
	userid, err := strconv.ParseInt(input.ID, 10, 64)
	if err != nil {
		panic(err)
	}
	request.Id = userid
	request.Username = input.Username
	request.Password = input.Password
	request.Email = input.Email
	request.Age = int64(input.Age)
	user, err := r.UserCtrl.UpdateUser(ctx, request)
	if err != nil {
		return &result, err
	}
	s := strconv.Itoa(int(user.ID))
	result.ID = s
	result.Username = user.Username
	result.Password = user.Password
	result.Email = user.Email
	result.Age = int(user.Age)
	return &result, nil
}

// Deleteuser is the resolver for the deleteuser field.
func (r *mutationResolver) Deleteuser(ctx context.Context, id string) (string, error) {
	userid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		panic(err)
	}
	errDelete := r.UserCtrl.DeleteUser(ctx, userid)
	if errDelete != nil {
		return "failed", errDelete
	}
	return "success", nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented: Todos - todos"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	var result []*model.User
	users, err := r.UserCtrl.GetAllUser(ctx)
	if err != nil {
		return result, err
	}
	for _, element := range users {
		var user model.User
		s := strconv.Itoa(int(element.ID))
		user.ID = s
		user.Username = element.Username
		user.Password = element.Password
		user.Email = element.Email
		user.Age = int(element.Age)
		result = append(result, &user)
	}
	return result, nil
}

// Getuserbyid is the resolver for the getuserbyid field.
func (r *queryResolver) Getuserbyid(ctx context.Context, id string) (*model.User, error) {
	var result model.User
	userid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		panic(err)
	}
	user, err := r.UserCtrl.GetUserByID(ctx, userid)
	if err != nil {
		return &result, err
	}
	s := strconv.Itoa(int(user.ID))
	result.ID = s
	result.Username = user.Username
	result.Password = user.Password
	result.Email = user.Email
	result.Age = int(user.Age)
	return &result, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
