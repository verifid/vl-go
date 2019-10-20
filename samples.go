package main

import (
	"fmt"
	"net/http"
	"path"
	"time"

	"github.com/verifid/vl-go/vlgo"
)

func main() {
	var httpClient = &http.Client{
		Timeout: time.Second * 15,
	}

	// Send User Data
	user := vlgo.User{
		Country:     "United States",
		DateOfBirth: "10.04.1980",
		Name:        "Tony",
		Surname:     "Stark"}

	userClient := vlgo.NewUserService(httpClient)
	userResponse, resp, err := userClient.User.SendUserData(user)
	fmt.Println(userResponse)
	fmt.Println(resp)
	fmt.Println(err)

	// Upload Identity Image
	imageClient := vlgo.NewImageService(httpClient)

	identityImagePath := path.Dir("/resources/2.png")
	base64Str := imageClient.Image.ImageFileToBase64(identityImagePath)
	imageUpload := vlgo.ImageUpload{Image: base64Str, UserID: "userId"}

	uploadResponse, resp, err := imageClient.Image.UploadIdentity(imageUpload)
	fmt.Println(uploadResponse)
	fmt.Println(resp)
	fmt.Println(err)

	profileImagePath := path.Dir("/resources/2.png")
	base64Str = imageClient.Image.ImageFileToBase64(profileImagePath)
	imageUpload = vlgo.ImageUpload{Image: base64Str, UserID: "userId"}

	// Upload Profile Image
	uploadResponse, resp, err = imageClient.Image.UploadProfile(imageUpload)
	fmt.Println(uploadResponse)
	fmt.Println(resp)
	fmt.Println(err)

	// Verify User
	verifyUser := vlgo.VerifyUser{
		UserID:   "userId",
		Language: "en_core_web_sm"}
	userVerificationResponse, httResponse, err := userClient.User.VerifyUser(verifyUser)
	fmt.Println(userVerificationResponse)
	fmt.Println(httResponse)
	fmt.Println(err)
}
