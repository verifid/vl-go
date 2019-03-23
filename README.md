# vl-go [![Build Status](https://travis-ci.org/verifid/vlgo.svg?branch=master)](https://travis-ci.org/verifid/vlgo)

**vlgo** is a http wrapper for identity verification layer. For now it has 4 main functions which is enough to verify user's identity.

### Features

* VerifID `vl` REST API:
    * User
    * Image

## Install

    go get github.com/verifid/vlgo/vlgo

## Usage

### REST API

The `vlgo` package provides a `Client` for accessing the VerifID vl API. Here are some example requests.

```go
client := new(vlgo.Client)
var httpClient = &http.Client{
    Timeout: time.Second * 15,
}
client.HttpClient = httpClient

user := User{
    Country:      "United States",
    DateOfBirth:  "10.04.1980",
    Gender:       "Male",
    Name:         "Tony",
    PlaceOfBirth: "New York",
    Surname:      "Stark"}
userResponse := new(UserResponse)
userResponse, resp, err := client.SendUserData(user)
```
