package preauth

import (
	"encoding/json"
	"fmt"

	"github.com/ElioenaiFerrari/tds/client"
)

func Call(request *Request) (*Response, error) {
	c := client.New[Response]()
	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	response, exception, err := c.Post("/preauth", body)
	if err != nil {
		return nil, err
	}

	if exception != nil {
		return nil, fmt.Errorf("%s: %s", exception.ErrorDescription, exception.ErrorDetail)
	}

	return &response, nil
}
