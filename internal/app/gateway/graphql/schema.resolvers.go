package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

	"github.com/davidchristie/tasks/internal/app/gateway/auth"
	"github.com/davidchristie/tasks/internal/app/gateway/graphql/generated"
	"github.com/davidchristie/tasks/internal/app/gateway/graphql/model"
)

func (r *mutationResolver) CreateTask(ctx context.Context, input model.CreateTask) (*model.Task, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) LoggedInUser(ctx context.Context) (*model.User, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return nil, nil
	}
	return &model.User{
		Avatar: user.AvatarURL,
		ID:     strconv.Itoa(user.ID),
		Name:   user.Name,
	}, nil
}

func (r *queryResolver) Tasks(ctx context.Context) ([]*model.Task, error) {
	return []*model.Task{}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
