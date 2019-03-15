package vlgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitStruct(t *testing.T) {
	user := User{country: "United States", dateOfBirth: "10.04.1980", gender: "Male", name: "Tony", placeOfBirth: "New York", surname: "Stark"}
	assert.NotNil(t, user)
}
