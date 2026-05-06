package handlers

import (
	"time"

	taskdomain "example.com/taskservice/internal/domain/task"
)

type taskMutationDTO struct {//новые мутации
    Title              string              `json:"title"`
    Description        string              `json:"description"`
    Status             taskdomain.Status   `json:"status"`
    RecurrenceType     taskdomain.RecurrenceType `json:"recurrence_type"`
    RecurrenceInterval *int                `json:"recurrence_interval,omitempty"`
    RecurrenceDays     []int               `json:"recurrence_days,omitempty"`
    RecurrenceDates    []string            `json:"recurrence_dates,omitempty"`
    RecurrenceParity   *string             `json:"recurrence_parity,omitempty"`
    RecurrenceStart    *string             `json:"recurrence_start,omitempty"`
	RecurrenceEnd      *string             `json:"recurrence_end,omitempty"`
}

type taskDTO struct {
    ID                 int64               `json:"id"`
    Title              string              `json:"title"`
    Description        string              `json:"description"`
    Status             taskdomain.Status   `json:"status"`
    CreatedAt          time.Time           `json:"created_at"`
    UpdatedAt          time.Time           `json:"updated_at"`
    RecurrenceType     taskdomain.RecurrenceType `json:"recurrence_type"`
    RecurrenceInterval *int                `json:"recurrence_interval,omitempty"`
    RecurrenceDays     []int               `json:"recurrence_days,omitempty"`
    RecurrenceDates    []string            `json:"recurrence_dates,omitempty"`
    RecurrenceParity   *string             `json:"recurrence_parity,omitempty"`
    RecurrenceStart    *string             `json:"recurrence_start,omitempty"`
	RecurrenceEnd      *string             `json:"recurrence_end,omitempty"`
}

func newTaskDTO(task *taskdomain.Task) taskDTO {
    var start, end *string
    if task.RecurrenceStart != nil {
        s := task.RecurrenceStart.Format("2006-01-02")
        start = &s
    }
    if task.RecurrenceEnd != nil {
        s := task.RecurrenceEnd.Format("2006-01-02")
        end = &s
    }
    dates := make([]string, len(task.RecurrenceDates))
    for i, d := range task.RecurrenceDates {
        dates[i] = d.Format("2006-01-02")
    }
    return taskDTO{
        ID:                 task.ID,
        Title:              task.Title,
        Description:        task.Description,
        Status:             task.Status,
        CreatedAt:          task.CreatedAt,
        UpdatedAt:          task.UpdatedAt,
        RecurrenceType:     task.RecurrenceType,
        RecurrenceInterval: task.RecurrenceInterval,
        RecurrenceDays:     task.RecurrenceDays,
        RecurrenceDates:    dates,
        RecurrenceParity:   task.RecurrenceParity,
        RecurrenceStart:    start,
        RecurrenceEnd:      end,
    }
}