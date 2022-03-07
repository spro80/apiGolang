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
	fmt.Println("")
	processOrderUseCase
	response := useCaseHealth{
		Status: "UP useCase One",
	}

	return c.JSON(http.StatusOK, response)
}*/

type SaveResponse struct {
	OrderId     string `json:"order_id"`
	Description string `json:"description"`
}

type saveOrderHandler struct {
	processOrderController order.OrderControllerInterface
}

func NewSaveOrderHandler(e *echo.Echo, processOrderController order.OrderControllerInterface) *saveOrderHandler {
	saveOrderHandler := &saveOrderHandler{
		processOrderController: processOrderController,
	}
	//e.POST("/order/save", saveOrderHandler.Save)
	e.GET("/order", saveOrderHandler.Save)
	return saveOrderHandler
}

func (s *saveOrderHandler) Save(c echo.Context) error {

	fmt.Print("[web-routes] save_oder.go Init")

	response, err := s.processOrderController.Process()
	if err != nil {
		return err
	}
	fmt.Print("[web-routes] save_oder.go Response:")
	fmt.Print(response)

	return c.JSON(http.StatusCreated, SaveResponse{
		OrderId:     "123456789",
		Description: "message save successfully",
	})

	/*if !config.GetBool("feature.flags.save") {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{Description: "save order disabled"})
	}

	saveOrderRequest := new(models.SaveOrderRequest)

	if err := c.Bind(saveOrderRequest); err != nil {
		err := errors.New("error getting order payload")
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{Description: err.Error()})
	}

	if err := c.Validate(saveOrderRequest); err != nil {
		var msgError string
		var split string
		for _, e := range err.(validator.ValidationErrors) {
			msgError = fmt.Sprintf("%s%s%s", msgError, split, e)
			split = ", "
		}
		err := errors.New(fmt.Sprintf("error validating data structure: %s", msgError))
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{Description: err.Error()})
	}

	order := &entity.Order{Id: saveOrderRequest.OrderID, Channel: saveOrderRequest.Channel}

	orderValue, err := s.processOrderUseCase.Process(order)
	if err != nil {
		customErr, ok := err.(*custom_errors.RequestError)
		if ok {
			switch customErr.Kind() {
			case custom_errors.DataBaseError:
				return c.JSON(http.StatusFailedDependency, models.ErrorResponse{
					Kind:        customErr.Kind(),
					Description: err.Error(),
				})
			default:
				return c.JSON(http.StatusInternalServerError, models.ErrorResponse{Kind: custom_errors.Unknown, Description: err.Error()})
			}
		}
	}

	return c.JSON(http.StatusCreated, models.SaveResponse{
		OrderId:     orderValue.Id,
		Description: "message save successfully",
	})*/
}
