package userControllerGetById

import (
	"fmt"

	"github.com/spro80/apiGolang/app/infraestructure/services"
)

type UserControlerGetByIdInterface interface {
	Process(userId string) (services.UserServiceGetByIdResponse, error)
}

type UserControllerGetByIdHandler struct {
	service services.UserServiceGetByIdInterface
}

func NewUserControllerGetBydId(service services.UserServiceGetByIdInterface) *UserControllerGetByIdHandler {
	return &UserControllerGetByIdHandler{
		service: service,
	}
}

func (t *UserControllerGetByIdHandler) Process(id string) (services.UserServiceGetByIdResponse, error) {
	fmt.Println("[controller][apiGetUsersByIdController] Calling service user_get_by_id")
	response, err := t.service.ProcessServicesGetUsersById(id)
	if err != nil {
		fmt.Println("Error in call to Services user_get_by_id")
		fmt.Println(err.Error())
		return response, err
	}
	fmt.Println("[controller][apiGetUsersByIdController] Service user_get_by_id was called successfully")
	return response, nil
}
