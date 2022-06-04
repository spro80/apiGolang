package routes

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/spro80/apiGolang/app/infraestructure/controllers/userControllerGetAll"
)

type RouterApiGetAllUsersHandler struct {
	userControllerGetAll userControllerGetAll.UserControllerGetAllInterface
}

func NewRouterApiGetAllUsersHandler(e *echo.Echo, userControllerGetAll userControllerGetAll.UserControllerGetAllInterface) *RouterApiGetAllUsersHandler {
	RouterApiGetAllUsersHandler := &RouterApiGetAllUsersHandler{
		userControllerGetAll: userControllerGetAll,
	}
	e.GET("/apiGetAllUsers", RouterApiGetAllUsersHandler.Process)
	return RouterApiGetAllUsersHandler
}

func (t RouterApiGetAllUsersHandler) Process(c echo.Context) error {
	fmt.Println("[routes][getAllUsersRouter] Calling Controller GetAllUsers")
	response, err := t.userControllerGetAll.Process()
	if err != nil {
		fmt.Println("[routes] error != nil")
	}
	fmt.Println("[routes][getAllUsersRouter] Controller GetAllUsers was called successfuly")
	return c.JSON(http.StatusCreated, response)
}
