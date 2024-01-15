package utils

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
)

const (
	MethodPost        = "POST"
	MethodGet         = "GET"
	SuccessStatusCode = 200
)

type Resp struct {
	StatusCode int
	Body       []byte
}

func PostJSON(ctx context.Context, url string, reqBody []byte, headers map[string]string) (*Resp, error) {
	payload := bytes.NewReader(reqBody)

	client := &http.Client{}

	req, err := http.NewRequest(MethodPost, url, payload)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	for headerK, headerV := range headers {
		req.Header.Add(headerK, headerV)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("have no response")
	}

	defer res.Body.Close()

	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return &Resp{
		StatusCode: res.StatusCode,
		Body:       respBody,
	}, nil
}
