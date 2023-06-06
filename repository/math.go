package repository

import (
	"fmt"
	"gorm.io/gorm"
	"mathOperation/db"
	"mathOperation/entities"
)

// MathOperationRepository repo for math operation.
type MathOperationRepository struct {
	DB *gorm.DB
}

// InsertRequestAndResponse insert the request and response of the API requested
func (m *MathOperationRepository) InsertRequestAndResponse(apiLogEntity entities.APILogs) error {
	db := m.DB
	result := db.Create(&apiLogEntity)
	if result.Error != nil {
		return fmt.Errorf("error while inserting data into api logs table. Err %w", result.Error)
	}
	return nil
}

// NewMathOperationRepository return instance of the type.
func NewMathOperationRepository() *MathOperationRepository {
	mathRepo := new(MathOperationRepository)
	mathRepo.DB = db.GetDB()
	return mathRepo
}
