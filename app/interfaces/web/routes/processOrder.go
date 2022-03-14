package routes

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/spro80/apiGolang/app/infraestructure/controllers/order"
)

//"github.com/spro80/apiGolang/app/infraestructure/controllers/useCaseOne"
/*
type useCaseHealthCheckHandler struct {
	processOrderUseCase useCaseOne.OrderUseCase
}

type useCaseHealth struct {
	Status string `json:"status"`
}

func NewUseCaseHealthCheckHandler(e *echo.Echo, processOrderUseCase useCaseOne.OrderUseCase) {
	h := &useCaseHealthCheckHandler{}
	e.GET("/healthUseCaseOne", h.UseCaseOne)
}

//func (p *useCaseHealthCheckHandler) UseCaseOne(c echo.Context) error {
func (p *useCaseHealthCheckHandler) UseCaseOne(c echo.Context, processOrderUseCase useCaseOne.OrderUseCase) error {
	processOrderUseCase
	response := useCaseHealth{
		Status: "UP useCase One",
	}

	return c.JSON(http.StatusOK, response)
}*/

type processOrderResponse struct {
}

type processOrderHandler struct {
	processOrderController order.OrderControllerInterface
}

func NewProcessOrderHandler(e *echo.Echo, processOrderController order.OrderControllerInterface) *processOrderHandler {
	processOrderHandler := &processOrderHandler{
		processOrderController: processOrderController,
	}
	e.GET("/processOrder", processOrderHandler.ProcessOrder)
	return processOrderHandler
}

func (s *processOrderHandler) ProcessOrder(c echo.Context) error {

	fmt.Println("[web-routes] processOrder.go Init")

	responseController, err := s.processOrderController.Process()
	if err != nil {
		return err
	}
	fmt.Println("[web-routes] processOrder.go Response:")
	fmt.Println(responseController)

	return c.JSON(http.StatusCreated, responseController)
}
