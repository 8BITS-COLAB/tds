package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/ElioenaiFerrari/tds/exception"
)

type Client[T any] struct {
	baseURL string
}

func New[T any]() *Client[T] {
	return &Client[T]{
		baseURL: os.Getenv("TDS_URL"),
	}
}

func (c *Client[T]) urlWithPath(path string) string {
	return fmt.Sprintf("%s%s", c.baseURL, path)
}

func (c *Client[T]) withDefaultHeaders(request *http.Request) {
	request.Header.Add("APIKey", os.Getenv("TDS_API_KEY"))
	request.Header.Add("Content-Type", "application/json; charset=utf-8")
}

func (c *Client[T]) Post(path string, body []byte) (T, *exception.Response, error) {
	var t T
	request, err := http.NewRequest(http.MethodPost, c.urlWithPath(path), bytes.NewReader(body))

	if err != nil {
		return t, nil, err
	}

	c.withDefaultHeaders(request)

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return t, nil, err
	}

	b, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return t, nil, err
	}

	if strings.Contains(string(b), "errorCode") {
		var exception exception.Response

		if err := json.Unmarshal(b, &exception); err != nil {
			return t, nil, err
		}

		return t, &exception, nil
	}

	if err := json.Unmarshal(b, &t); err != nil {
		return t, nil, err
	}

	return t, nil, nil
}
