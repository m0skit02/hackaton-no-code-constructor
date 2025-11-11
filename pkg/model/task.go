package models

import (
	"log"
	"time"

	"github.com/lib/pq"

	"github.com/google/uuid"
)

type TaskStatus string

const (
	TaskPending    TaskStatus = "pending"
	TaskInProgress TaskStatus = "in_progress"
	TaskDone       TaskStatus = "done"
	TaskError      TaskStatus = "error"
)

type Task struct {
	ID          uuid.UUID      `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime"`
	Status      TaskStatus     `json:"status" gorm:"type:task_status;default:'pending'"`
	Target      string         `json:"target" gorm:"not null"`
	Checks      pq.StringArray `json:"checks" gorm:"type:check_type[]"`
	AgentID     *uuid.UUID     `json:"agent_id,omitempty"`
	RequestedBy *string        `json:"requested_by,omitempty"`
}

func MustParseUUID(s string) uuid.UUID {
	id, err := uuid.Parse(s)
	if err != nil {
		log.Fatalf("invalid UUID: %v", err)
	}
	return id
}
