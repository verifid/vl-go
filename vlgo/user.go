package vlgo

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Country      string `json:"country"`
	DateOfBirth  string `json:"dateOfBirth"`
	Gender       string `json:"gender"`
	Name         string `json:"name"`
	PlaceOfBirth string `json:"placeOfBirth"`
	Surname      string `json:"surname"`
}

func UserToJson(user User) string {
	json, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(json)
}
