package vlgo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type User struct {
	Country      string `json:"country"`
	DateOfBirth  string `json:"dateOfBirth"`
	Gender       string `json:"gender"`
	Name         string `json:"name"`
	PlaceOfBirth string `json:"placeOfBirth"`
	Surname      string `json:"surname"`
}

type UserResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Type    string `json:"type"`
	UserId  string `json:"user_id"`
}

func UserToJson(user User) []byte {
	b, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return b
}

func SendUserData(user User) (*UserResponse, *http.Response, error) {
	b := UserToJson(user)
	resp, err := http.Post("https://api.verifid.app/v1/user/sendData", "application/json", bytes.NewBuffer(b))
	if err != nil {
		return nil, resp, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	responseString := string(body)
	fmt.Println(responseString)
	if err != nil {
		return nil, resp, err
	}
	userResponse := new(UserResponse)
	err = json.Unmarshal(body, &userResponse)
	if err != nil {
		return nil, resp, err
	}
	return userResponse, resp, nil
}
