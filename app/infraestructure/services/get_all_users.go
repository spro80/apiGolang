package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
)

type GetAllUsersServicesInterface interface {
	ProcessServicesGetAllUsers() (Response, error)
}

type Response struct {
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

type GetAllUsersServices struct {
}

func NewGetUsers() *GetAllUsersServices {
	return &GetAllUsersServices{}
}

func (g GetAllUsersServices) ProcessServicesGetAllUsers() ([]Response, error) {

	//urlComplete := "https://jsonplaceholder.typicode.com/todos"
	//url := "https://jsonplaceholder.typicode.com/"
	//path := "todos"

	fmt.Println("[services][get_all_users] Calling GetAllUsers 1")
	response, err := http.Get("https://jsonplaceholder.typicode.com/todos")
	if err != nil {
		fmt.Println("[services][get_all_users] err distinto de nil")
		fmt.Print(err.Error())
	}
	fmt.Println(reflect.TypeOf(response))
	fmt.Println(reflect.TypeOf(response.Body))

	fmt.Println(response)
	fmt.Println("[services][get_all_users] Calling GetAllUsers 1.5")
	fmt.Println(response.Body)
	fmt.Println("[services][get_all_users] Calling GetAllUsers 2")
	body, err0 := ioutil.ReadAll(response.Body)
	if err0 != nil {
		fmt.Println("[services][get_all_users] err0 fatal")
		log.Fatal(err0)
	}
	fmt.Println(reflect.TypeOf(body))
	fmt.Println("[services][get_all_users] Calling GetAllUsers3")

	if len(body) == 0 {
		fmt.Errorf("no bytes to unmarshal")
	}

	fmt.Println(body)

	fmt.Println(body)
	fmt.Println(len(body))
	fmt.Println("body[0]:::")
	fmt.Println(body[0])
	fmt.Println("body[1]:::")
	fmt.Println(body[1])

	sBody := string(body)
	//fmt.Println("sBody:::")
	//fmt.Println(sBody)
	fmt.Println("sBody0:::")
	fmt.Println(sBody[0])

	switch body[0] {
	case '{':
		fmt.Println("response single")
	case '[':
		fmt.Println("response array")
	}

	var tags []Response
	err3 := json.Unmarshal(body, &tags)
	if err3 != nil {
		fmt.Println("error in unmarshal")
	}
	//g.Tags = tags
	fmt.Println("g.Tags:::")
	fmt.Println(tags)
	fmt.Println(reflect.TypeOf(tags))

	/*
			{
		    "userId": 2,
		    "id": 30,
		    "title": "nemo perspiciatis repellat ut dolor libero commodi blanditiis omnis",
		    "completed": true
		  },
		  {
		    "userId": 2,
		    "id": 31,
		    "title": "repudiandae totam in est sint facere fuga",
		    "completed": false
		  },*/

	/*
		fmt.Println(reflect.TypeOf(sBody))
		var jsonMaps2 Response
		err3 := json.Unmarshal([]byte(sBody), &jsonMaps2)
		if err3 != nil {
			fmt.Println("[services][get_all_users] err3 distinto de nil")
			fmt.Printf(err3.Error())
		}

		fmt.Println(reflect.TypeOf(jsonMaps2))
		fmt.Println("[services][get_all_users] return GetAllUsers")
	*/

	/*
		responseController := Response{
			UserId:    10,
			Id:        1,
			Title:     "title",
			Completed: true,
		}

		return responseController, nil*/

	return tags, nil

}
