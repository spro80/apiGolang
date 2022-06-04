package userControllerGetAll

import (
	"fmt"
	"reflect"

	"github.com/spro80/apiGolang/app/infraestructure/services"
)

type UserControllerGetAllInterface interface {
	Process() ([]services.UserResponseGetAllService, error)
}

type UserControllerGetAllHandler struct {
	service services.UserServiceGetAllInterface
}

type UserControllerResponseApiGetAll struct {
	StatusCode        int    `json:"statusCode"`
	StatusDescription string `json:"statusDescription"`
	StatusError       string `json:"statusError"`
	Data              string `json:"data"`
}

func NewUserControllerGetAll(service services.UserServiceGetAllInterface) *UserControllerGetAllHandler {
	return &UserControllerGetAllHandler{
		service: service,
	}
}

func (t *UserControllerGetAllHandler) Process() ([]services.UserResponseGetAllService, error) {
	fmt.Println("[controller][] Init in method Process of controller")
	response, err := t.service.ProcessGetAllUsersServices()
	if err != nil {
		fmt.Println("Error in call Process Services Get All Users")
	}
	fmt.Println("[controller][GetAllServices] services was called succesffullys")
	fmt.Println(reflect.TypeOf(response))
	fmt.Println(response)

	return response, nil
	/*
		res1 := services.Response{
			UserId:    1,
			Id:        1,
			Title:     "title1",
			Completed: true,
		}
		res2 := services.Response{
			UserId:    2,
			Id:        2,
			Title:     "title2",
			Completed: true,
		}
		responseController := []services.Response{res1, res2}
		return responseController, nil
	*/
}
