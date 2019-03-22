package vlgo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

type ImageUploadResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Type    string `json:"type"`
}

func FileUploadRequest(url string, params map[string]string, paramName, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, err
}

func (client *Client) UploadIdentity(userId string, imagePath string) (*ImageUploadResponse, *http.Response, error) {
	extraParams := map[string]string{
		"userId": userId,
	}
	req, err := FileUploadRequest(fmt.Sprintf("%s/image/uploadIdentity", baseURL), extraParams, "file", imagePath)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to build request")
	}
	resp, err := client.httpClient.Do(req)
	if err != nil {
		return nil, resp, errors.Wrap(err, "request failed")
	}
	uploadResponse := new(ImageUploadResponse)
	if err := json.NewDecoder(resp.Body).Decode(&uploadResponse); err != nil {
		return nil, resp, errors.Wrap(err, "unmarshaling failed")
	}
	return uploadResponse, resp, nil
}

func (client *Client) UploadProfile(userId string, imagePath string) (*ImageUploadResponse, *http.Response, error) {
	extraParams := map[string]string{
		"userId": userId,
	}
	req, err := FileUploadRequest(fmt.Sprintf("%s/image/uploadProfile", baseURL), extraParams, "file", imagePath)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to build request")
	}
	resp, err := client.httpClient.Do(req)
	if err != nil {
		return nil, resp, errors.Wrap(err, "request failed")
	}
	uploadResponse := new(ImageUploadResponse)
	if err := json.NewDecoder(resp.Body).Decode(&uploadResponse); err != nil {
		return nil, resp, errors.Wrap(err, "unmarshaling failed")
	}
	return uploadResponse, resp, nil
}
