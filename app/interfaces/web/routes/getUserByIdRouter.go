package routes

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/spro80/apiGolang/app/infraestructure/controllers/userControllerGetById"
)

type User struct {
	UserID    string `json:"userId"`
	ID        int    `json:"id" query:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type RouterGetUserByIdResponse struct {
	Status string                  `json:"status"`
	Data   []RouterGetUserByIdData `json:"data"`
	Error  string                  `json:"error"`
}

type RouterGetUserByIdData struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id" query:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type RouterApiGetUserByIdHandler struct {
	userControllerGetById userControllerGetById.UserControllerGetByIdInterface
}

func NewRouteApiGetUserByIdHandler(e *echo.Echo, userControllerGetById userControllerGetById.UserControllerGetByIdInterface) *RouterApiGetUserByIdHandler {
	RouterApiGetUserByIdHandler := &RouterApiGetUserByIdHandler{
		userControllerGetById: userControllerGetById,
	}
	e.GET("/getUserById/:id", RouterApiGetUserByIdHandler.Process)
	return RouterApiGetUserByIdHandler
}

func (t RouterApiGetUserByIdHandler) Process(c echo.Context) error {
	//https://jsonplaceholder.typicode.com/todos/4
	/*{
		"userId": 1,
		"id": 4,
		"title": "et porro tempora",
		"completed": true
	}*/

	res := RouterGetUserByIdResponse{}

	id := c.Param("id")
	fmt.Println(reflect.TypeOf(id))
	idValue, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("[routes][getUserByIdRouter] Error in paramneter id")
		fmt.Println(err.Error())

		res.Status = "nok"
		res.Error = err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}
	fmt.Println(idValue)

	fmt.Println("[routes][getUserByIdRouter] Calling to controller apiGetUsersByIdController")
	response, err := t.userControllerGetById.Process(id)
	fmt.Println("[routes][getUserByIdRouter] Controller apiGetUsersByIdController was called succesfully")
	if err != nil {
		fmt.Println("[routes][getUserByIdRouter] error")
		fmt.Println(err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	return c.JSON(http.StatusCreated, response)
}
