package order

import (
	"fmt"

	"github.com/spro80/apiGolang/app/application/orderUseCase"
)

type OrderControllerInterface interface {
	Process() (Response, error)
}
type orderControllerHandler struct {
	orderUseCase orderUseCase.OrderUseCaseInterface
}

type Response struct {
	Status            bool   `json:"status"`
	StatusCode        int    `json:"statusCode"`
	StatusDescription string `json:"statusDescription"`
	StatusError       string `json:"statusError"`
}

func NewOrderController(orderUseCase orderUseCase.OrderUseCaseInterface) *orderControllerHandler {
	orderController := &orderControllerHandler{
		orderUseCase: orderUseCase,
	}
	return orderController
}

func (s *orderControllerHandler) Process() (Response, error) {
	fileName := "controllers-order"
	fmt.Printf("\n[%s] Init in healtchCheck from controller", fileName)
	fmt.Printf("\n[%s] Calling to use case Order", fileName)

	var res, err = s.orderUseCase.Process()
	if err != nil {
		fmt.Printf("\n[%s] Error calling use case Order", fileName)
	}

	fmt.Printf("\n[%s] Use case order was called succesfully", fileName)
	fmt.Println(res)

	var status = res.Status
	var statusCode = res.StatusCode
	var statusDes = res.StatusDescription
	var statusErr = res.StatusError
	responseController := Response{
		Status:            status,
		StatusCode:        statusCode,
		StatusDescription: statusDes,
		StatusError:       statusErr,
	}
	PresenterTopic("Present to message in topic kafka.")
	PresenterSlack("Present to message in slack.")

	return responseController, err

}

func PresenterTopic(message string) (string, error) {
	fileName := "controllers-order"
	fmt.Printf("\n[%s] Init in PresenterTopic", fileName)
	return message, nil
}

func PresenterSlack(message string) (string, error) {
	fileName := "controllers-order"
	fmt.Printf("\n[%s] Init in PresenterSlack", fileName)
	return message, nil
}

func CalculateSum(n1 int, n2 int) int {
	return n1 + n2
}
