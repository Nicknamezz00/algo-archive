package model

import (
	"gorm.io/plugin/soft_delete"
)

// Model Common model
type Model struct {
	ID        int64                 `gorm:"primaryKey" json:"id"`
	CreatedAt int64                 `json:"created_at"`
	UpdatedAt int64                 `json:"updated_at"`
	DeletedAt int64                 `json:"deleted_at"`
	IsDeleted soft_delete.DeletedAt `gorm:"softDelete:flag" json:"is_deleted"`
}

type ConditionsT map[string]interface{}
