package task

import (
	"context"

	taskdomain "example.com/taskservice/internal/domain/task"
)

type Repository interface {
	Create(ctx context.Context, task *taskdomain.Task) (*taskdomain.Task, error)
	GetByID(ctx context.Context, id int64) (*taskdomain.Task, error)
	Update(ctx context.Context, task *taskdomain.Task) (*taskdomain.Task, error)
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context) ([]taskdomain.Task, error)
}

type Usecase interface {
	Create(ctx context.Context, input CreateInput) (*taskdomain.Task, error)
	GetByID(ctx context.Context, id int64) (*taskdomain.Task, error)
	Update(ctx context.Context, id int64, input UpdateInput) (*taskdomain.Task, error)
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context) ([]taskdomain.Task, error)
}

type CreateInput struct {// добавление новых полей дял периодичных тасок
	Title              string
	Description        string
	Status             taskdomain.Status
	RecurrenceType     taskdomain.RecurrenceType
	RecurrenceInterval *int
	RecurrenceDays     []int
	RecurrenceDates    []string
	RecurrenceParity   *string
	RecurrenceStart    *string
	RecurrenceEnd      *string
}

type UpdateInput struct {// такая же история
	Title              string
	Description        string
	Status             taskdomain.Status
	RecurrenceType     taskdomain.RecurrenceType
	RecurrenceInterval *int
	RecurrenceDays     []int
	RecurrenceDates    []string
	RecurrenceParity   *string
	RecurrenceStart    *string
	RecurrenceEnd      *string
}
