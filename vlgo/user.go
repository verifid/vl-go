package vlgo

import (
	"net/http"

	"github.com/dghubble/sling"
)

// UserService provides an interface for user endpoints.
type UserService struct {
	sling *sling.Sling
}

// newUserService returns a new UserService.
func newUserService(sling *sling.Sling) *UserService {
	return &UserService{
		sling: sling,
	}
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

// SendUserData sends user data using with Client.
// Takes user as a parameter.
// Returns user response, http response and error.
func (s *UserService) SendUserData(user User) (*UserResponse, *http.Response, error) {
	userResponse := new(UserResponse)
	resp, err := s.sling.New().Post("/user/sendUserData").QueryStruct(user).ReceiveSuccess(userResponse)
	return userResponse, resp, err
}

// VerifyUser verifies user with given user id and language.
// func (client *Client) VerifyUser(verifyUser VerifyUser) (*UserVerificationResponse, *http.Response, error) {
func (s *UserService) VerifyUser(verifyUser VerifyUser) (*UserVerificationResponse, *http.Response, error) {
	userVerificationResponse := new(UserVerificationResponse)
	resp, err := s.sling.New().Post("/user/verify").QueryStruct(verifyUser).ReceiveSuccess(userVerificationResponse)
	return userVerificationResponse, resp, err
}
