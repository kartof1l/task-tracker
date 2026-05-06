package task

import (
	"context"
	"fmt"
	"strings"
	"time"

	taskdomain "example.com/taskservice/internal/domain/task"
)

type Service struct {
	repo Repository
	now  func() time.Time
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
		now:  func() time.Time { return time.Now().UTC() },
	}
}

func (s *Service) Create(ctx context.Context, input CreateInput) (*taskdomain.Task, error) {
	normalized, err := validateCreateInput(input)//метод изменен в   соответствии с новой валидацией
	if err != nil {
		return nil, err
	}
	model := &taskdomain.Task{//обновил модель соответственно input
		Title:              normalized.Title,
		Description:        normalized.Description,
		Status:             normalized.Status,
		RecurrenceType:     normalized.RecurrenceType,
		RecurrenceInterval: normalized.RecurrenceInterval,
		RecurrenceDays:     normalized.RecurrenceDays,
		RecurrenceParity:   normalized.RecurrenceParity,

	}
	var dates []time.Time
	for _, d := range normalized.RecurrenceDates {
		t, err := time.Parse("2006-01-02", d)
		if err != nil {
			return nil, fmt.Errorf("%w: неверный формат даты в recurrence_dates", ErrInvalidInput)
		}
		dates = append(dates, t)
	}
	model.RecurrenceDates = dates
	if normalized.RecurrenceStart != nil {
    t, err := time.Parse("2006-01-02", *normalized.RecurrenceStart)
    if err != nil {
        return nil, fmt.Errorf("%w: неверный формат recurrence_start", ErrInvalidInput)
    }
    model.RecurrenceStart = &t
	}
	if normalized.RecurrenceEnd != nil {
		t, err := time.Parse("2006-01-02", *normalized.RecurrenceEnd)
		if err != nil {
			return nil, fmt.Errorf("%w: неверный формат recurrence_end", ErrInvalidInput)
		}
		model.RecurrenceEnd = &t
	}
	if err := model.ValidateRecurrence();err != nil{
		return nil, fmt.Errorf("%w: %v", ErrInvalidInput, err)
	}
	now := s.now()
	model.CreatedAt = now
	model.UpdatedAt = now

	created, err := s.repo.Create(ctx, model)
	if err != nil {
		return nil, err
	}

	return created, nil
}

func (s *Service) GetByID(ctx context.Context, id int64) (*taskdomain.Task, error) {
	if id <= 0 {
		return nil, fmt.Errorf("%w: id must be positive", ErrInvalidInput)
	}

	return s.repo.GetByID(ctx, id)
}

func (s *Service) Update(ctx context.Context, id int64, input UpdateInput) (*taskdomain.Task, error) {
	if id <= 0 {
		return nil, fmt.Errorf("%w: id must be positive", ErrInvalidInput)
	}

	normalized, err := validateUpdateInput(input)
	if err != nil {
		return nil, err
	}

	model := &taskdomain.Task{
		Title:              normalized.Title,
		Description:        normalized.Description,
		Status:             normalized.Status,
		RecurrenceType:     normalized.RecurrenceType,
		RecurrenceInterval: normalized.RecurrenceInterval,
		RecurrenceDays:     normalized.RecurrenceDays,
		RecurrenceParity:   normalized.RecurrenceParity,
	}
	var dates []time.Time//опять же новая  переменная для конвертации, инициировал здесь для наглядности
	for _, d := range normalized.RecurrenceDates {
		t, err := time.Parse("2006-01-02", d)
		if err != nil {
			return nil, fmt.Errorf("%w: неверный формат даты в recurrence_dates", ErrInvalidInput)
		}
		dates = append(dates, t)
	}
	model.RecurrenceDates = dates
	if normalized.RecurrenceStart != nil {
    t, err := time.Parse("2006-01-02", *normalized.RecurrenceStart)
    if err != nil {
        return nil, fmt.Errorf("%w: неверный формат recurrence_start", ErrInvalidInput)
    }
    model.RecurrenceStart = &t
	}
	if err := model.ValidateRecurrence();err != nil{
		return nil, fmt.Errorf("%w: %v", ErrInvalidInput, err)
	}
	if normalized.RecurrenceEnd != nil {
		t, err := time.Parse("2006-01-02", *normalized.RecurrenceEnd)
		if err != nil {
			return nil, fmt.Errorf("%w: неверный формат recurrence_end", ErrInvalidInput)
		}
		model.RecurrenceEnd = &t
	}
	updated, err := s.repo.Update(ctx, model)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

func (s *Service) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return fmt.Errorf("%w: id must be positive", ErrInvalidInput)
	}

	return s.repo.Delete(ctx, id)
}

func (s *Service) List(ctx context.Context) ([]taskdomain.Task, error) {
	return s.repo.List(ctx)
}

func validateCreateInput(input CreateInput) (CreateInput, error) {
	input.Title = strings.TrimSpace(input.Title)
	input.Description = strings.TrimSpace(input.Description)
	if input.RecurrenceType == ""{
		input.RecurrenceType = "none"
	}
	if !input.RecurrenceType.Valid() {
    return CreateInput{}, fmt.Errorf("%w: неверный тип периодичности", ErrInvalidInput)
}
	if input.Title == "" {
		return CreateInput{}, fmt.Errorf("%w: title is required", ErrInvalidInput)
	}

	if input.Status == "" {
		input.Status = taskdomain.StatusNew
	}

	if !input.Status.Valid() {
		return CreateInput{}, fmt.Errorf("%w: invalid status", ErrInvalidInput)
	}

	return input, nil
}

func validateUpdateInput(input UpdateInput) (UpdateInput, error) {
	input.Title = strings.TrimSpace(input.Title)
	input.Description = strings.TrimSpace(input.Description)
	if input.RecurrenceType == ""{
		input.RecurrenceType = "none"
	}
	if !input.RecurrenceType.Valid() {
    return UpdateInput{}, fmt.Errorf("%w: неверный тип периодичности", ErrInvalidInput)
}
	if input.Title == "" {
		return UpdateInput{}, fmt.Errorf("%w: title is required", ErrInvalidInput)
	}

	if !input.Status.Valid() {
		return UpdateInput{}, fmt.Errorf("%w: invalid status", ErrInvalidInput)
	}

	return input, nil
}
