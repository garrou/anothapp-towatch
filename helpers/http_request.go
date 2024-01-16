package helpers

import (
	"io"
	"net/http"
)

func HttpGet(url, key string) []byte {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("X-BetaSeries-Key", key)
	resp, getErr := client.Do(req)

	if getErr != nil {
		panic(getErr.Error())
	}
	defer func(body io.ReadCloser) {
		if err := body.Close(); err != nil {
			panic(err.Error())
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err.Error())
	}
	return body
}
