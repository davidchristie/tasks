package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/davidchristie/tasks/internal/app/gateway/auth"
	"github.com/davidchristie/tasks/internal/app/gateway/entity"
	"github.com/davidchristie/tasks/internal/app/gateway/graphql/format"
	"github.com/davidchristie/tasks/internal/app/gateway/graphql/generated"
	"github.com/davidchristie/tasks/internal/app/gateway/graphql/model"
	"github.com/google/uuid"
)

func (r *mutationResolver) CreateTask(ctx context.Context, input model.CreateTask) (*model.Task, error) {
	loggedInUser := auth.ForContext(ctx)
	if loggedInUser == nil {
		return nil, ErrMustBeLoggedIn
	}

	task := entity.Task{
		CreatedAt:       time.Now().Format(time.RFC3339),
		CreatedByUserID: loggedInUser.ID,
		ID:              uuid.New(),
		Text:            input.Text,
	}

	if task.Text == "" {
		return nil, errors.New("text cannot be empty")
	}

	err := r.Database.InsertTask(ctx, &task)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return nil, errors.New("error creating task")
	}

	return format.Task(&task), nil
}

func (r *mutationResolver) DeleteTask(ctx context.Context, id string) (*model.Task, error) {
	loggedInUser := auth.ForContext(ctx)
	if loggedInUser == nil {
		return nil, ErrMustBeLoggedIn
	}

	taskID, err := uuid.Parse(id)
	if err != nil {
		return nil, ErrNotFound
	}

	task, err := r.Database.FindTaskByID(taskID)

	if task.CreatedByUserID != loggedInUser.ID {
		return nil, ErrNotFound
	}

	err = r.Database.DeleteTask(task.ID)
	if err != nil {
		return nil, errors.New("error deleting task")
	}

	return format.Task(task), nil
}

func (r *queryResolver) LoggedInUser(ctx context.Context) (*model.User, error) {
	loggedInUser := auth.ForContext(ctx)
	if loggedInUser == nil {
		return nil, nil
	}
	return &model.User{
		Avatar: loggedInUser.AvatarURL,
		ID:     loggedInUser.ID.String(),
		Name:   loggedInUser.Name,
	}, nil
}

func (r *queryResolver) Tasks(ctx context.Context) ([]*model.Task, error) {
	const limit = 25 // TODO: Add this as an argument.

	loggedInUser := auth.ForContext(ctx)
	if loggedInUser == nil {
		return nil, ErrMustBeLoggedIn
	}

	tasks, err := r.Database.FindTasksCreatedByUserID(loggedInUser.ID, limit)
	if err != nil {
		return nil, err
	}

	return format.Tasks(tasks), nil
}

func (r *taskResolver) CreatedBy(ctx context.Context, obj *model.Task) (*model.User, error) {
	user, err := r.Database.FindUserByID(obj.CreatedByUserID)
	if err != nil {
		return nil, ErrNotFound
	}
	return &model.User{
		Avatar: user.AvatarURL,
		ID:     user.ID.String(),
		Name:   user.Name,
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Task returns generated.TaskResolver implementation.
func (r *Resolver) Task() generated.TaskResolver { return &taskResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type taskResolver struct{ *Resolver }
