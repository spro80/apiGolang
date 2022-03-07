package orderUseCase

import (
	"fmt"
)

type SaveResponse struct {
	Status string `json:status`
}

//func OrderUseCase() (bool, error) {
func OrderUseCase() (bool, error) {
	fmt.Println("[orderUseCase] OrderUseCase Init in orderUseCase")
	/*var response := SaveResponse{
		Status: "Ok from use casse"
	} */

	var id int = 100
	var name string = "David"

	status := SaveInDatabase(id, name)
	/*if err != nil {
		fmt.Println("[orderUseCase] OrderUseCase Error al guardar en base de datos.")
	}*/
	fmt.Println("[orderUseCase] OrderUseCase status de saveInDatabase")
	fmt.Println(status)

	return true, nil
}

func SaveInDatabase(id int, name string) bool {
	fmt.Println("[orderUseCase] saveInDatabase")
	fmt.Println("[orderUseCase] saveInDatabase id:")
	fmt.Println(id)
	fmt.Println("[orderUseCase] saveInDatabase name:")
	fmt.Println(name)
	return true
}
