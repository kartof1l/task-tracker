package task

import (
	"errors"
	"fmt"
	"time"
)

type Status string
type RecurrenceType string //новый тип как в миграции
const (
	StatusNew        Status = "new"
	StatusInProgress Status = "in_progress"
	StatusDone       Status = "done"
	RecurrenceNone     RecurrenceType = "none" //новые константы, соответственно видам задачи
    RecurrenceDaily    RecurrenceType = "daily"
    RecurrenceMonthly  RecurrenceType = "monthly"
    RecurrenceSpecific RecurrenceType = "specific"
    RecurrenceParity   RecurrenceType = "parity"
)

type Task struct {
	ID                 int64           `json:"id"`
	Title              string          `json:"title"`
	Description        string          `json:"description"`
	Status             Status          `json:"status"`
	CreatedAt          time.Time       `json:"created_at"`
	UpdatedAt          time.Time       `json:"updated_at"`
	RecurrenceType     RecurrenceType  `json:"recurrence_type"`
	RecurrenceInterval *int            `json:"recurrence_interval,omitempty"`
	RecurrenceDays     []int           `json:"recurrence_days,omitempty"`
/*насколько я знаю могут возникнуть проблемы если сделать слайс типа time.Time,
	поэтому во избежание проблем с  []time.Time и DATE[] будут строки*/
	RecurrenceDates    []time.Time     `json:"recurrence_dates,omitempty"`
	RecurrenceParity   *string         `json:"recurrence_parity,omitempty"`
	RecurrenceStart    *time.Time      `json:"recurrence_start,omitempty"`
	RecurrenceEnd      *time.Time      `json:"recurrence_end,omitempty"`
}
func (s Status) Valid() bool {
	switch s {
	case StatusNew, StatusInProgress, StatusDone:
		return true
	default:
		return false
	}
}
func (r RecurrenceType) Valid() bool {//для валидации новых полей как в статусе
	switch r {
	case RecurrenceNone, RecurrenceDaily, RecurrenceMonthly, RecurrenceSpecific, RecurrenceParity:
		return true
	default:
		return false
	}
}
func (t Task) ValidateRecurrence() error {//новый метод для валидации сочетания полей
	if !t.RecurrenceType.Valid() {
		return fmt.Errorf("неизвестный тип периодичности: %s", t.RecurrenceType)
	}

	switch t.RecurrenceType {
	case RecurrenceNone:
		return nil
	case RecurrenceDaily:
		if t.RecurrenceInterval == nil || *t.RecurrenceInterval < 1 {
			return errors.New("для ежедневной периодичности интервал должен быть не меньше 1")
		}
	case RecurrenceMonthly:
		if len(t.RecurrenceDays) == 0 {
			return errors.New("для ежемесячной периодичности нужно указать хотя бы одно число месяца")
		}
		for _, day := range t.RecurrenceDays {
			if day < 1 || day > 30 {//лучше затронуть большую часть месяцев, нежели уделять много внимания феврали, туда дату лучше вписать вручную, как и на 31 числа
				return errors.New("числа месяца должны быть от 1 до 30")
			}
		}
	case RecurrenceSpecific:
		if len(t.RecurrenceDates) == 0 {
			return errors.New("для периодичности по конкретным датам нужно указать хотя бы одну дату")
		}
	case RecurrenceParity:
		if t.RecurrenceParity == nil {
			return errors.New("для периодичности по чётности нужно указать 'even' или 'odd'")
		}
		if *t.RecurrenceParity != "even" && *t.RecurrenceParity != "odd" {
			return errors.New("чётность должна быть 'even' (чётные) или 'odd' (нечётные)")
		}
	}

	return nil
}
