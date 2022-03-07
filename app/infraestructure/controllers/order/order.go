package order

import (
	"fmt"

	"github.com/spro80/apiGolang/app/application/orderUseCase"
)

//	"github.com/spro80/apiGolang/app/application/orderUseCase"

type OrderControllerInterface interface {
	Process() (bool, error)
}

type orderController struct {
	responseController int
}

func NewOrderController() *orderController {
	return &orderController{}
}

func (s *orderController) Process() (bool, error) {
	fmt.Println("[controllers-getHealthCheck] Init in healtchCheck from controller")
	//var sum int = CalculateSum(99, 2)
	//fmt.Println(sum)

	fmt.Println("[controllers-getHealthCheck] Calling to use case Order")
	var status, err = orderUseCase.OrderUseCase()
	if err != nil {
		fmt.Println("[controllers-getHealthCheck] Error in called to use case")
	}
	fmt.Println("[controllers-getHealthCheck] Use case order was called succesfully")
	fmt.Println(status)

	return status, err
}

func CalculateSum(n1 int, n2 int) int {
	return n1 + n2
}
