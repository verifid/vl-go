package vlgo

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"

	"github.com/dghubble/sling"
)

// ImageService provides an interface for image endpoints.
type ImageService struct {
	sling *sling.Sling
}

// newImageService returns a new ImageService.
func newImageService(sling *sling.Sling) *ImageService {
	return &ImageService{
		sling: sling,
	}
}

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
func (imageService *ImageService) ImageFileToBase64(filePath string) string {
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

func (imageService *ImageService) uploadImage(imageUpload ImageUpload, imageType Enum) (*ImageUploadResponse, *http.Response, error) {
	imageUploadResponse := new(ImageUploadResponse)
	resp, err := imageService.sling.New().Post(fmt.Sprintf("/image/%s", Enum.ValueOfImageType(imageType))).QueryStruct(imageUpload).ReceiveSuccess(imageUploadResponse)
	return imageUploadResponse, resp, err
}

// UploadIdentity uploads identity image of user.
// Takes user id and image path as parameters.
// Returns image upload response, http response and error.
func (imageService *ImageService) UploadIdentity(imageUpload ImageUpload) (*ImageUploadResponse, *http.Response, error) {
	return imageService.uploadImage(imageUpload, ImageType.Identity)
}

// UploadProfile uploads profile image of user.
// Takes user id and image path as parameters.
// Returns image upload response, http response and error.
func (imageService *ImageService) UploadProfile(imageUpload ImageUpload) (*ImageUploadResponse, *http.Response, error) {
	return imageService.uploadImage(imageUpload, ImageType.Profile)
}
