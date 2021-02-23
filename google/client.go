package google

import (
	"context"

	"google.golang.org/api/serviceusage/v1"
)

// Client is a wrapper of the Google API client
type Client struct {
	ServiceUsage *serviceusage.Service
}

// NewClient returns a new Client
func NewClient() (*Client, error) {
	serviceusageClient, err := serviceusage.NewService(context.TODO())
	if err != nil {
		return nil, err
	}

	return &Client{
		ServiceUsage: serviceusageClient,
	}, nil
}
