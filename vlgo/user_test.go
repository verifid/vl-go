package vlgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitStruct(t *testing.T) {
	user := User{
		Country:      "United States",
		DateOfBirth:  "10.04.1980",
		Gender:       "Male",
		Name:         "Tony",
		PlaceOfBirth: "New York",
		Surname:      "Stark"}
	assert.Equal(t, user.Country, "United States")
	assert.Equal(t, user.DateOfBirth, "10.04.1980")
	assert.Equal(t, user.Gender, "Male")
	assert.Equal(t, user.Name, "Tony")
	assert.Equal(t, user.PlaceOfBirth, "New York")
	assert.Equal(t, user.Surname, "Stark")
}

func TestUserToJson(t *testing.T) {
	user := User{
		Country:      "United States",
		DateOfBirth:  "10.04.1980",
		Gender:       "Male",
		Name:         "Tony",
		PlaceOfBirth: "New York",
		Surname:      "Stark"}
	json := UserToJson(user)
	assert.Equal(t, json, "{\"country\":\"United States\",\"dateOfBirth\":\"10.04.1980\",\"gender\":\"Male\",\"name\":\"Tony\",\"placeOfBirth\":\"New York\",\"surname\":\"Stark\"}")
}
