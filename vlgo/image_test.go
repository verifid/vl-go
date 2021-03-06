package vlgo

import (
	"net/http"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUploadIdentity(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		okResponse := `{
			"code": 200, "message": "Image file received.", "type": "success"
		}`
		w.Write([]byte(okResponse))
	})
	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	client := NewImageService(httpClient)

	path := path.Dir("../resources/2.png")
	imageUpload := ImageUpload{client.Image.ImageFileToBase64(path), "userId"}

	uploadResponse, resp, err := client.Image.UploadIdentity(imageUpload)
	assert.Nil(t, err)
	assert.Equal(t, uploadResponse.Code, 200)
	assert.NotNil(t, resp)
}

func TestUploadProfile(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		okResponse := `{
			"code": 200, "message": "Image file received.", "type": "success"
		}`
		w.Write([]byte(okResponse))
	})
	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	client := NewImageService(httpClient)

	path := path.Dir("../resources/2.png")
	imageUpload := ImageUpload{client.Image.ImageFileToBase64(path), "userId"}

	uploadResponse, resp, err := client.Image.UploadProfile(imageUpload)
	assert.Nil(t, err)
	assert.Equal(t, uploadResponse.Code, 200)
	assert.NotNil(t, resp)
}
