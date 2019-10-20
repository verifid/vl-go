package vlgo

import (
	"net/http"

	"github.com/dghubble/sling"
)

// Client is a Verification Layer client for making requests.
type Client struct {
	sling *sling.Sling
	User  *UserService
	Image *ImageService
}

// NewUserService returns a new UserService.
func NewUserService(httpClient *http.Client) *Client {
	base := sling.New().Client(httpClient).Base("https://verifid.app")
	return &Client{
		sling: base,
		User:  newUserService(base.New()),
	}
}

// NewImageService returns a new UserService.
func NewImageService(httpClient *http.Client) *Client {
	base := sling.New().Client(httpClient).Base("https://verifid.app")
	return &Client{
		sling: base,
		Image: newImageService(base.New()),
	}
}
