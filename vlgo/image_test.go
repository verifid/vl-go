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

	client := new(Client)
	client.httpClient = httpClient

	path := path.Dir("../resources/2.png")
	uploadResponse, resp, err := client.UploadIdentity("userId", path)
	assert.Nil(t, err)
	assert.NotNil(t, uploadResponse)
	assert.NotNil(t, resp)
}
