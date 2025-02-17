package format

import (
	"github.com/davidchristie/tasks/internal/app/gateway/entity"
	"github.com/davidchristie/tasks/internal/app/gateway/graphql/model"
)

func Task(task *entity.Task) *model.Task {
	if task == nil {
		return nil
	}
	return &model.Task{
		CreatedAt:       task.CreatedAt,
		CreatedByUserID: task.CreatedByUserID,
		Done:						 task.CompletedAt != nil,
		ID:              task.ID.String(),
		Text:            task.Text,
	}
}

func Tasks(tasks []*entity.Task) []*model.Task {
	output := []*model.Task{}
	for _, task := range tasks {
		output = append(output, Task(task))
	}
	return output
}
