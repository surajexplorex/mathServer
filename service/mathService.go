package service

import (
	"context"
	"fmt"
	"log"
	"mathOperation/constants"
	"mathOperation/entities"
	protos "mathOperation/proto/mathServer"
	"mathOperation/repository"
)

type MathService struct {
	mathRepository *repository.MathOperationRepository
	protos.UnimplementedMathOperationsServer
}

// Add will perform the addition operation
func (t *MathService) Add(ctx context.Context, i *protos.MathInput) (*protos.MathOutput, error) {

	log.Printf(`Adding two numbers %f and %f`, i.GetFirstNumber(), i.GetSecondNumber())
	result := i.GetFirstNumber() + i.GetSecondNumber()
	tra := &protos.MathOutput{
		Output:  result,
		Message: fmt.Sprintf("Multiplication result of %f and %f", i.GetFirstNumber(), i.GetSecondNumber()),
	}
	mathOperationEntity := entities.APILogs{
		Request:       i.String(),
		Response:      tra.String(),
		OperationType: constants.ADD,
	}
	err := t.mathRepository.InsertRequestAndResponse(mathOperationEntity)
	if err != nil {
		fmt.Errorf("error while inserting request response data into DB. Err - %w", err)
		return tra, err
	}
	return tra, nil
}

// Subtract will perform the subtraction operation
func (t *MathService) Subtract(ctx context.Context, i *protos.MathInput) (*protos.MathOutput, error) {

	log.Printf(`Subtracting two numbers %f and %f`, i.GetFirstNumber(), i.GetSecondNumber())
	result := i.GetFirstNumber() - i.GetSecondNumber()

	tra := &protos.MathOutput{
		Output:  result,
		Message: fmt.Sprintf("Subtraction result of %f and %f", i.GetFirstNumber(), i.GetSecondNumber()),
	}
	log.Printf(`Subtract operation response %s`, tra.String())
	mathOperationEntity := entities.APILogs{
		Request:       i.String(),
		Response:      tra.String(),
		OperationType: constants.SUBTRACT,
	}
	err := t.mathRepository.InsertRequestAndResponse(mathOperationEntity)
	if err != nil {
		fmt.Errorf("error while inserting request response data into DB. Err - %w", err)
		return tra, err
	}
	return tra, nil
}

// Multiply will perform the multiplication operation
func (t *MathService) Multiply(ctx context.Context, i *protos.MathInput) (*protos.MathOutput, error) {

	log.Printf(`Multiplying two numbers %f with %f`, i.GetFirstNumber(), i.GetSecondNumber())
	result := i.GetFirstNumber() * i.GetSecondNumber()
	tra := &protos.MathOutput{
		Output:  result,
		Message: fmt.Sprintf("Multiplication result of %f and %f", i.GetFirstNumber(), i.GetSecondNumber()),
	}
	mathOperationEntity := entities.APILogs{
		Request:       i.String(),
		Response:      tra.String(),
		OperationType: constants.MULTIPLY,
	}
	err := t.mathRepository.InsertRequestAndResponse(mathOperationEntity)
	if err != nil {
		fmt.Errorf("error while inserting request response data into DB. Err - %w", err)
		return tra, err
	}
	return tra, nil
}

// Divide will perform the multiplication operation
func (t *MathService) Divide(ctx context.Context, i *protos.MathInput) (*protos.MathOutput, error) {

	log.Printf(`Dividing two numbers %f with %f`, i.GetFirstNumber(), i.GetSecondNumber())
	if i.GetSecondNumber() == 0 {
		return &protos.MathOutput{
			Output:  0,
			Message: "second number can't be zero",
		}, nil
	}
	result := i.GetFirstNumber() / i.GetSecondNumber()
	tra := &protos.MathOutput{
		Output:  result,
		Message: fmt.Sprintf("Division result of %f and %f", i.GetFirstNumber(), i.GetSecondNumber()),
	}
	mathOperationEntity := entities.APILogs{
		Request:       i.String(),
		Response:      tra.String(),
		OperationType: constants.DIVIDE,
	}
	err := t.mathRepository.InsertRequestAndResponse(mathOperationEntity)
	if err != nil {
		fmt.Errorf("error while inserting request response data into DB. Err - %w", err)
		return tra, err
	}
	return tra, nil
}

func NewMathOperationServer() *MathService {
	mathServer := new(MathService)
	mathServer.mathRepository = repository.NewMathOperationRepository()
	return mathServer
}
