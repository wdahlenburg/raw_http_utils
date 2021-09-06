package main

import (
	"fmt"
	"net/http"

	"github.com/wdahlenburg/raw_http_utils/utils"
)

func main() {
	resp, err := http.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	raw := &utils.RawHttpResponse{
		Response: resp,
	}

	result, err := raw.GetRawResponse()
	if err != nil {
		panic(err)
	}

	fmt.Printf(string(result))
}
