package userapi

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
)

const (
	ValidateTokenEndpoint = "/api/v1/user/auth/token/validate"
)

type httpClient struct {
	baseUrl string
	client  *resty.Client
}

func NewHttpClient(baseUrl string, client *resty.Client) *httpClient {
	return &httpClient{
		baseUrl: baseUrl,
		client:  client,
	}
}

func (c *httpClient) ValidateToken(ctx context.Context, token string) error {
	targetUrl := fmt.Sprintf("%s%s", c.baseUrl, ValidateTokenEndpoint)
	resp, err := c.client.R().SetHeader("Authorization", token).EnableTrace().Post(targetUrl)
	if err != nil {
		return err
	}
	if resp.IsError() {
		return fmt.Errorf("%v", resp.Error())
	}
	return nil
}
