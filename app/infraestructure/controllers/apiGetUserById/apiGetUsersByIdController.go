package apiGetUserById

import (
	"fmt"
	"reflect"

	"github.com/spro80/apiGolang/app/infraestructure/services"
)

type ControllerApiGetUserByIdInterface interface {
	Process(userId string) (services.ServiceGetUserByIdResponse, error)
}

type ControllerApiGetUserByIdHandler struct {
	service services.ServiceGetUserByIdInterface
}

func NewControllerApiGetUserBydId(service services.ServiceGetUserByIdInterface) *ControllerApiGetUserByIdHandler {
	return &ControllerApiGetUserByIdHandler{
		service: service,
	}
}

func (t *ControllerApiGetUserByIdHandler) Process(id string) (services.ServiceGetUserByIdResponse, error) {
	fmt.Println("[controller][apiGetUsersByIdController] Calling service GetUserById")
	response, err := t.service.ProcessServicesGetUsersById(id)
	if err != nil {
		fmt.Println("Error in call Process Services Get Users By Id")
	}
	fmt.Println("[controller][apiGetUsersByIdController] Service GetUserById was called successfully")
	fmt.Println(reflect.TypeOf(response))
	fmt.Println(response)

	return response, nil
}
