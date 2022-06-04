package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type UserServiceGetAllInterface interface {
	ProcessGetAllUsersServices() ([]UserResponseGetAllService, error)
}

type UserResponseGetAllService struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

/*
type CreateTagsResponse struct {
	Tags []Response
}

func (ctr *CreateTagsResponse) unmarshalSingle(b []byte) error {
	var t Tag
	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}
	ctr.Tags = []Tag{t}
	return nil
}*/

/*
func (ctr *CreateTagsResponse) unmarshalMany(b []byte) error {
	var tags []Response
	err := json.Unmarshal(b, &tags)
	if err != nil {
		return err
	}
	ctr.Tags = tags
	return nil
}*/

type GetAllUsersServicesHandler struct {
}

func NewGetUsers() *GetAllUsersServicesHandler {
	return &GetAllUsersServicesHandler{}
}

/***********************************************
url: "https://jsonplaceholder.typicode.com/todos"
method: "GET"
response:
[
  {
    "userId": 1,
    "id": 1,
    "title": "delectus aut autem",
    "completed": false
  },
  {
    "userId": 1,
    "id": 2,
    "title": "quis ut nam facilis et officia qui",
    "completed": false
  }
]
************************************************/
func (g GetAllUsersServicesHandler) ProcessGetAllUsersServices() ([]UserResponseGetAllService, error) {

	var tags []UserResponseGetAllService
	var tagsArray []UserResponseGetAllService
	var tagsJson UserResponseGetAllService

	fmt.Println("[services][get_all_users] Calling GetAllUsers")
	response, err := http.Get("https://jsonplaceholder.typicode.com/todos")
	if err != nil {
		fmt.Println("[services][get_all_users] err distinto de nil")
		fmt.Print(err.Error())
		return tags, errors.New("[services][get_all_users] Error in Calling Get")
	}

	fmt.Println("[services][get_all_users] Calling ReadAll")
	body, errReadAll := ioutil.ReadAll(response.Body)
	if errReadAll != nil {
		fmt.Println("[services][get_all_users] Error in Called to ReadAll")
		return tags, errors.New("[services][get_all_users] Error in ReadAll")
	}
	fmt.Println("[services][get_all_users] ReadAll was called successfully")

	resAnalize, errAnalize := analizeResponse(body)
	if errAnalize != nil {
		fmt.Println("[services][get_all_users] Error captured in AnalizeResponse")
		return tags, errAnalize
	}

	if resAnalize == "array" {
		fmt.Println("se generara response del array")
		errUnmarshal := json.Unmarshal(body, &tagsArray)
		if errUnmarshal != nil {
			fmt.Println("error in unmarshal del array")
			return tags, errors.New("Error in unmarshal array")
		}
		tags = tagsArray
	}
	if resAnalize == "json" {
		fmt.Println("se generara response del json como array")
		errUnmarshal := json.Unmarshal(body, &tagsJson)
		if errUnmarshal != nil {
			fmt.Println("error in unmarshal del json")
			return tags, errors.New("Error in unmarshal json")
		}
		tags = []UserResponseGetAllService{tagsJson}
	}

	return tags, nil
}

func analizeResponse(body []byte) (string, error) {

	if len(body) == 0 {
		return "nodata", errors.New("The Response Body don't have elements")
	}

	switch body[0] {
	case '{':
		return "json", nil
	case '[':
		return "array", nil
	default:
		return "default", errors.New("The format of the response is distinct to Json and Array")
	}
}
