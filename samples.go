package main

import (
	"fmt"
	"net/http"
	"path"
	"time"

	"github.com/verifid/vl-go/vlgo"
)

func main() {
	// vlgo client
	client := new(vlgo.Client)
	var httpClient = &http.Client{
		Timeout: time.Second * 15,
	}
	client.HTTPClient = httpClient

	// Send User Data
	user := vlgo.User{
		Country:     "United States",
		DateOfBirth: "10.04.1980",
		Name:        "Tony",
		Surname:     "Stark"}
	userResponse, resp, err := client.SendUserData(user)
	fmt.Println(userResponse)
	fmt.Println(resp)
	fmt.Println(err)

	path := path.Dir("/resources/2.png")
	base64Str := client.ImageFileToBase64(path)
	imageUpload := vlgo.ImageUpload{Image: base64Str, UserID: "userId"}
	uploadResponse, resp, err := client.UploadIdentity(imageUpload, vlgo.ImageType.Identity)
	fmt.Println(uploadResponse)
	fmt.Println(resp)
	fmt.Println(err)

	verifyUser := vlgo.VerifyUser{
		UserID:   "userId",
		Language: "en_core_web_sm"}
	userVerificationResponse, httResponse, err := client.VerifyUser(verifyUser)
	fmt.Println(userVerificationResponse)
	fmt.Println(httResponse)
	fmt.Println(err)
}
