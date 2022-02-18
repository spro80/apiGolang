package controllers

import (
	"fmt"
)

type OrderUseCase interface {
	Process() (bool, error)
}

type orderUseCase struct {
}

func NewOrderUseCase() *orderUseCase {
	return &orderUseCase{}
}

func (s *orderUseCase) Process() (bool, error) {
	fmt.Print("[controllers-getHealthCheck] Init in healtchCheck")
	return true, nil
}
