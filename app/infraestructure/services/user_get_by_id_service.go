package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

type UserServiceGetByIdInterface interface {
	ProcessServicesGetUsersById(userId string) (UserServiceGetByIdResponse, error)
}

/*
type ServiceGetUserByIdResponse struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}*/

/*
type ServiceGetUserByIdResponse struct {
	Status string `json:"status"`
	Data   struct {
		UserID    string `json:"userId"`
		ID        string `json:"id"`
		Title     string `json:"title"`
		Completed string `json:"completed"`
	} `json:"data"`
	//Error string `json:"error"`
}*/

/*
type ServiceGetUserByIdResponse struct {
	Status string `json:"status"`
	Data   []struct {
		UserID    string `json:"userId"`
		ID        string `json:"id"`
		Title     string `json:"title"`
		Completed string `json:"completed"`
	} `json:"data"`
}
*/

type UserServiceGetByIdResponse struct {
	Status string                   `json:"status"`
	Data   []UserServiceGetByIdData `json:"data"`
	Error  string                   `json:"error"`
}

type UserServiceGetByIdData struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id" query:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type UserServiceGetByIdHandler struct {
}

func NewGetUsersById() *UserServiceGetByIdHandler {
	return &UserServiceGetByIdHandler{}
}

/***********************************************
url: "https://jsonplaceholder.typicode.com/todos/1"
method: "GET"
response:
  {
    "userId": 1,
    "id": 1,
    "title": "delectus aut autem",
    "completed": false
  }
************************************************/
func (g UserServiceGetByIdHandler) ProcessServicesGetUsersById(id string) (UserServiceGetByIdResponse, error) {

	var responseInfo UserServiceGetByIdResponse
	var tags []UserServiceGetByIdData
	//var tagsArray []ServiceGetUserByIdData
	var tagsJson UserServiceGetByIdData

	urlBase := "https://jsonplaceholder.typicode.com/"
	//path := "todos/"
	pathAll := "todos/"
	urlBasePath := urlBase + pathAll

	u, err := url.Parse(urlBase)
	u.Path = path.Join(u.Path, pathAll+id)
	urlFinal := u.String() // prints http://foo/bar.html
	fmt.Println("urlFinal:")
	fmt.Println(urlFinal)
	fmt.Println("urlBasePath:")
	fmt.Println(urlBasePath)

	//url := "https://jsonplaceholder.typicode.com/todos/"
	/*
		v := url.Values{}
		v.Set("id", id)
		perform := urlBasePath + "?" + v.Encode()
		fmt.Println("perform:::")
		fmt.Println(perform)
	*/

	//fmt.Println("[services][get_users_by_id] Calling Get Users By Id")
	response, err := http.Get(urlFinal)
	if err != nil {
		fmt.Println("[services][get_users_by_id] Error in call Get Users By Id")
		fmt.Print(err.Error())
		return responseInfo, errors.New("[services][get_users_by_id] Error in Calling Get Users By Id")
	}

	fmt.Println("[services][get_users_by_id] Calling ReadAll")
	fmt.Println(response.Body)
	body, errReadAll := ioutil.ReadAll(response.Body)
	if errReadAll != nil {
		fmt.Println("[services][get_users_by_id] Error in Called to ReadAll")
		fmt.Println(errReadAll.Error())
		return responseInfo, errors.New("[services][get_users_by_id] Error in ReadAll")
	}
	fmt.Println("[services][get_users_by_id] ReadAll was called successfully")

	resAnalize, errAnalize := analizeResponseUsersById(body)
	if errAnalize != nil {
		fmt.Println("[services][get_users_by_id] Error captured in AnalizeResponse")
		fmt.Println(errAnalize.Error())
		return responseInfo, errAnalize
	}

	/*if resAnalize == "array" {
		fmt.Println("se generara response del array")
		fmt.Println(body)
		fmt.Println(string(body))
		errUnmarshal := json.Unmarshal(body, &tagsArray)
		if errUnmarshal != nil {
			fmt.Println("[services][get_users_by_id] Error in unmarshal of array response")
			fmt.Println(errUnmarshal)
			return responseInfo, errors.New("[services][get_users_by_id] Error in unmarshal of array response")
		}
		tags = tagsArray
		responseInfo = ServiceGetUserByIdResponse{
			Status: "ok",
			Data:   tags,
			Error:  "",
		}
		fmt.Println(responseInfo)
	}*/

	if resAnalize == "json" {
		fmt.Println("Se generara response del json como array")
		errUnmarshal := json.Unmarshal(body, &tagsJson)
		if errUnmarshal != nil {
			fmt.Println("[services][get_users_by_id] Error in unmarshal of json response")
			fmt.Println(errUnmarshal.Error())
			return responseInfo, errors.New("[services][get_users_by_id] Error in unmarshal of json response")
		}
		tags = []UserServiceGetByIdData{tagsJson}
		responseInfo = UserServiceGetByIdResponse{
			Status: "ok",
			Data:   tags,
			Error:  "",
		}
		fmt.Println(responseInfo)
	}

	fmt.Println("[services][get_users_by_id] Returning successfully Users with id")
	return responseInfo, nil
}

func analizeResponseUsersById(body []byte) (string, error) {

	if len(body) == 0 {
		return "nodata", errors.New("[services][get_users_by_id] Errors: The Response Body don't have elements")
	}

	switch body[0] {
	case '{':
		return "json", nil
	case '[':
		return "array", nil
	default:
		return "default", errors.New("[services][get_users_by_id] Errors: The format of the response is distinct to Json and Array")
	}
}
