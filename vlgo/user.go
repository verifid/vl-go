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

// Client contains a http client where we use for all requests.
type Client struct {
	HTTPClient *http.Client
}

// User is request body for sending user data.
type User struct {
	Country     string `json:"country"`
	DateOfBirth string `json:"dateOfBirth"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
}

// UserResponse is the response model of send user data request.
type UserResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Type    string `json:"type"`
	UserID  string `json:"userId,omitempty"`
}

// VerifyUser is request model of user verification.
type VerifyUser struct {
	UserID   string `json:"userId"`
	Language string `json:"language"`
}

// UserVerificationResponse is response model of user verification request.
type UserVerificationResponse struct {
	Code             int `json:"code"`
	VerificationRate int `json:"verificationRate"`
}

// UserToJSON marshalls user struct.
func UserToJSON(user User) []byte {
	b, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return b
}

// VerifyUserToJSON marshalls verfiy user struct.
func VerifyUserToJSON(verifyUser VerifyUser) []byte {
	b, err := json.Marshal(verifyUser)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return b
}

// SendUserData sends user data using with Client.
// Takes user as a parameter.
// Returns user response, http response and error.
func (client *Client) SendUserData(user User) (*UserResponse, *http.Response, error) {
	b := UserToJSON(user)
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/user/sendUserData", baseURL), bytes.NewBufferString(string(b)))
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to build request")
	}
	resp, err := client.HTTPClient.Do(req)
	if err != nil {
		return nil, resp, errors.Wrap(err, "request failed")
	}
	userResponse := new(UserResponse)
	if err := json.NewDecoder(resp.Body).Decode(&userResponse); err != nil {
		return nil, resp, errors.Wrap(err, "unmarshaling failed")
	}
	return userResponse, resp, nil
}

// VerifyUser verifies user with given user id and language.
func (client *Client) VerifyUser(verifyUser VerifyUser) (*UserVerificationResponse, *http.Response, error) {
	b := VerifyUserToJSON(verifyUser)
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/user/verify", baseURL), bytes.NewBufferString(string(b)))
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to build request")
	}
	resp, err := client.HTTPClient.Do(req)
	if err != nil {
		return nil, resp, errors.Wrap(err, "request failed")
	}
	userVerificationResponse := new(UserVerificationResponse)
	if err := json.NewDecoder(resp.Body).Decode(&userVerificationResponse); err != nil {
		return nil, resp, errors.Wrap(err, "unmarshaling failed")
	}
	return userVerificationResponse, resp, nil
}
