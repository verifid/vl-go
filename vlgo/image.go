package vlgo

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/pkg/errors"
)

// ImageUpload is request model for image upload.
type ImageUpload struct {
	Image  string `json:"image"`
	UserID string `json:"userId"`
}

// ImageUploadResponse is response model for image upload.
type ImageUploadResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Type    string `json:"type"`
}

// ImageFileToBase64 read image file and creates base64 encoded string.
func (client *Client) ImageFileToBase64(filePath string) string {
	imgFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	defer imgFile.Close()

	// create a new buffer base on file size
	fInfo, _ := imgFile.Stat()
	var size int64 = fInfo.Size()
	buf := make([]byte, size)

	// read file content into buffer
	fReader := bufio.NewReader(imgFile)
	fReader.Read(buf)

	// convert the buffer bytes to base64 string - use buf.Bytes() for new image
	imgBase64Str := base64.StdEncoding.EncodeToString(buf)
	return imgBase64Str
}

func imageModelToJSON(imageUpload ImageUpload) []byte {
	b, err := json.Marshal(imageUpload)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return b
}

func (client *Client) uploadImage(imageUpload ImageUpload, imageType Enum) (*ImageUploadResponse, *http.Response, error) {
	b := imageModelToJSON(imageUpload)
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/image/%s", baseURL, Enum.ValueOfImageType(imageType)), bytes.NewBufferString(string(b)))
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to build request")
	}
	resp, err := client.HTTPClient.Do(req)
	if err != nil {
		return nil, resp, errors.Wrap(err, "request failed")
	}
	imageUploadResponse := new(ImageUploadResponse)
	if err := json.NewDecoder(resp.Body).Decode(&imageUploadResponse); err != nil {
		return nil, resp, errors.Wrap(err, "unmarshaling failed")
	}
	return imageUploadResponse, resp, nil
}

// UploadIdentity uploads identity image of user.
// Takes user id and image path as parameters.
// Returns image upload response, http response and error.
func (client *Client) UploadIdentity(imageUpload ImageUpload, imageType Enum) (*ImageUploadResponse, *http.Response, error) {
	return client.uploadImage(imageUpload, imageType)
}

// UploadProfile uploads profile image of user.
// Takes user id and image path as parameters.
// Returns image upload response, http response and error.
func (client *Client) UploadProfile(imageUpload ImageUpload, imageType Enum) (*ImageUploadResponse, *http.Response, error) {
	return client.uploadImage(imageUpload, imageType)
}
