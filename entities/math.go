package entities

import "time"

type APILogs struct {
	BaseEntity
	Request       string `json:"request" type:"string"`
	Response      string `json:"response" type:"string"`
	OperationType string `json:"operation_type" type:"string"`
}

type BaseEntity struct {
	Id        uint64     `json:"id" gorm:"primaryKey;AUTO_INCREMENT" mapstructure:"id"`
	CreatedAt *time.Time `json:"created_at" type:"date"`
	CreatedBy string     `json:"created_by" type:"string"`
	IsDeleted int        `json:"status" type:"integer"`
}
