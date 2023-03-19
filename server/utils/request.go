package utils

import (
	"context"
	"fmt"
	"net/http"
	"server/global"
)

func StreamRequest() (res interface{}, err error) {
	client := http.Client{}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "http://13.229.248.26/model", nil)
	if err != nil {
		fmt.Println("err-", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "text/event-stream")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", global.Gpt.Key))
	res, err = client.Do(req)
	fmt.Println("res", res)
	return res, err
}
