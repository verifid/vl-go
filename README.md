# vl-go [![Build Status](https://travis-ci.org/verifid/vl-go.svg?branch=master)](https://travis-ci.org/verifid/vl-go) [![GoDoc](https://godoc.org/github.com/verifid/vl-go/vlgo?status.svg)](https://godoc.org/github.com/verifid/vl-go/vlgo)


**vlgo** is a http wrapper for identity verification layer. For now it has 4 main functions which is enough to verify user's identity.

### Features

* VerifID `vl` REST API:
    * User
    * Image

## Install

    go get github.com/verifid/vl-go/vlgo

## Usage

### REST API

The `vlgo` package provides a `Client` for accessing the VerifID vl API. Here are some example requests.

```go
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
    Surname:     "Stark"
    }
userResponse, resp, err := client.SendUserData(user)

// Upload User Identity Image
path := path.Dir("/resources/2.png")
base64Str := client.ImageFileToBase64(path)
imageUpload := vlgo.ImageUpload{Image: base64Str, UserID: "userId"}
uploadResponse, resp, err := client.UploadIdentity(imageUpload, vlgo.ImageType.Identity)

// Upload User Profile Image
path := path.Dir("/resources/2.png")
base64Str := client.ImageFileToBase64(path)
imageUpload := vlgo.ImageUpload{Image: base64Str, UserID: "userId"}
uploadResponse, resp, err := client.UploadProfile(imageUpload, vlgo.ImageType.Profile)
```
