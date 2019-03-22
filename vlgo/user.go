package vlgo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

const (
	baseURL = "https://api.verifid.app/v1"
)

type Client struct {
	httpClient *http.Client
}

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

func (client *Client) SendUserData(user User) (*UserResponse, *http.Response, error) {
	b := UserToJson(user)
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/user/sendData", baseURL), bytes.NewBufferString(string(b)))
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to build request")
	}
	resp, err := client.httpClient.Do(req)
	if err != nil {
		return nil, resp, errors.Wrap(err, "request failed")
	}
	userResponse := new(UserResponse)
	if err := json.NewDecoder(resp.Body).Decode(&userResponse); err != nil {
		return nil, resp, errors.Wrap(err, "unmarshaling failed")
	}
	return userResponse, resp, nil
}
