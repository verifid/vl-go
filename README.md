# vl-go [![Build Status](https://travis-ci.org/verifid/vl-go.svg?branch=master)](https://travis-ci.org/verifid/vl-go) [![GoDoc](https://godoc.org/github.com/verifid/vl-go/vlgo?status.svg)](https://godoc.org/github.com/verifid/vl-go/vlgo) [![Go Report Card](https://goreportcard.com/badge/verifid/vl-go)](https://goreportcard.com/report/verifid/vl-go)


**vlgo** is a Go HTTP client library around the VerifID identity verification layer API. It's a complete wrapper contains all endpoints available on [Verification Layer](https://github.com/verifid/vl). Use of this client and API
is enough to verify someone's identity.

## Features

* VerifID `vl` REST API:
    * User
    * Image

## Install

    go get github.com/verifid/vl-go/vlgo

## Usage

### REST API

The `vlgo` package provides a `Client` for accessing the VerifID [Verification Layer](https://github.com/verifid/vl) API. You need to follow 4 basic steps to verify a person and his identity.

Steps of user verification

> 1. Send user's personal data
> 2. Upload photos of identity card of passport
> 3. Upload profile photo
> 4. Call verify user

Here are some example requests.

```go
import (
	"fmt"
	"net/http"
	"path"
	"time"

	"github.com/verifid/vl-go/vlgo"
)

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
```
