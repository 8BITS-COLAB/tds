package auth

import (
	"encoding/json"
	"fmt"

	"github.com/8BITS-COLAB/tds/client"
)

func Call(request *Request) (*Response, error) {
	c := client.New[Response]()
	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	response, exception, err := c.Post("/auth", body)
	if err != nil {
		return nil, err
	}

	if exception != nil {
		return nil, fmt.Errorf("%s: %s", exception.ErrorDescription, exception.ErrorDetail)
	}

	return &response, nil
}
