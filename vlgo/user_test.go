package vlgo

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testingHTTPClient(handler http.Handler) (*http.Client, func()) {
	s := httptest.NewTLSServer(handler)
	cli := &http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, network, _ string) (net.Conn, error) {
				return net.Dial(network, s.Listener.Addr().String())
			},
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	return cli, s.Close
}

const (
	okResponse = `{
		"code": 200, "message": "User created with received values.", "type": "success", "user_id": "37924286-c6c1-4f7a-a164-db8158171152"
	}`

	verifyUserOkResponse = `{
		"code": 200,
		"verificationRate": 85
	}`
)

func TestInitStruct(t *testing.T) {
	user := User{
		Country:     "United States",
		DateOfBirth: "10.04.1980",
		Name:        "Tony",
		Surname:     "Stark"}
	assert.Equal(t, user.Country, "United States")
	assert.Equal(t, user.DateOfBirth, "10.04.1980")
	assert.Equal(t, user.Name, "Tony")
	assert.Equal(t, user.Surname, "Stark")
}

func TestUserToJson(t *testing.T) {
	user := User{
		Country:     "United States",
		DateOfBirth: "10.04.1980",
		Name:        "Tony",
		Surname:     "Stark"}
	json := UserToJSON(user)
	assert.Equal(t, json, []byte("{\"country\":\"United States\",\"dateOfBirth\":\"10.04.1980\",\"name\":\"Tony\",\"surname\":\"Stark\"}"))
}

func TestSendUserData(t *testing.T) {
	user := User{
		Country:     "United States",
		DateOfBirth: "10.04.1980",
		Name:        "Tony",
		Surname:     "Stark"}
	userResponse := new(UserResponse)

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(okResponse))
	})
	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	client := new(Client)
	client.HTTPClient = httpClient

	userResponse, resp, err := client.SendUserData(user)
	assert.Nil(t, err)
	assert.Equal(t, userResponse.Code, 200)
	assert.NotNil(t, resp)
}

func TestVerifyUser(t *testing.T) {
	verifyUser := VerifyUser{
		UserID:   "userId",
		Language: "en_core_web_sm"}
	userVerificationResponse := new(UserVerificationResponse)
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(verifyUserOkResponse))
	})
	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	client := new(Client)
	client.HTTPClient = httpClient

	userVerificationResponse, resp, err := client.VerifyUser(verifyUser)
	assert.Nil(t, err)
	assert.Equal(t, userVerificationResponse.Code, 200)
	assert.NotNil(t, resp)
}
