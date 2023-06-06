package server

import (
	"context"
	"fmt"
	"log"
	"mathOperation/constants"
	"mathOperation/entities"
	protos "mathOperation/protos/mathServer"
	"mathOperation/repository"
)

type MathOperationServer struct {
	mathRepository *repository.MathOperationRepository
	protos.UnimplementedMathOperationsServer
}

// Add will perform the addition operation
func (t *MathOperationServer) Add(ctx context.Context, i *protos.MathInput) (*protos.MathOutput, error) {

	log.Printf(`Adding two numbers %f and %f`, i.GetFirstNumber(), i.GetSecondNumber())
	result := i.GetFirstNumber() + i.GetSecondNumber()
	tra := &protos.MathOutput{
		Output: result,
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
func (t *MathOperationServer) Subtract(ctx context.Context, i *protos.MathInput) (*protos.MathOutput, error) {

	log.Printf(`Subtracting two numbers %f and %f`, i.GetFirstNumber(), i.GetSecondNumber())
	result := i.GetFirstNumber() - i.GetSecondNumber()
	if i.GetFirstNumber() < i.GetSecondNumber() {
		result = i.GetSecondNumber() - i.GetFirstNumber()
	}
	tra := &protos.MathOutput{
		Output: result,
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
func (t *MathOperationServer) Multiply(ctx context.Context, i *protos.MathInput) (*protos.MathOutput, error) {

	log.Printf(`Multiplying two numbers %f with %f`, i.GetFirstNumber(), i.GetSecondNumber())
	result := i.GetFirstNumber() * i.GetSecondNumber()
	tra := &protos.MathOutput{
		Output: result,
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

func NewMathOperationServer() *MathOperationServer {
	mathServer := new(MathOperationServer)
	mathServer.mathRepository = repository.NewMathOperationRepository()
	return mathServer
}
