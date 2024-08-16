// Get response headers from an outgoing http request

package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	call("https://localhost:8085/example", "GET")
}

func call(url, method string) error {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return fmt.Errorf("Got error %s", err.Error())
	}
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Got error %s", err.Error())
	}
	defer res.Body.Close()

	contentTypeValues := res.Header.Values("content-type")
	fmt.Printf("res.Header.Values(\"content-type\"):: %s\n\n", contentTypeValues)

	contentType := res.Header.Get("content-type")
	fmt.Printf("res.Header.Get(\"content-type\"):: %s\n\n", contentType)

	headers := res.Header
	fmt.Printf("res.Headers:: %s\n\n", headers)
	fmt.Printf("res.Headers[\"Content-Type\"]:: %s\n\n", headers["Content-Type"])
	return nil

}
