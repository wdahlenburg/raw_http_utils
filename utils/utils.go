package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
)

type RawHttpResponse struct {
	Response *http.Response
}

func (r *RawHttpResponse) GetRawResponse() ([]byte, error) {
	var raw []byte

	// Add the first line - HTTP/1.1 200 OK

	raw = []byte(fmt.Sprintf("%s %s\r\n", r.Response.Proto, r.Response.Status))

	// Sort the headers alphabetically because net/http uses a map, which has no order

	keys := make([]string, 0, len(r.Response.Header))
	for k := range r.Response.Header {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		for _, value := range r.Response.Header[key] {
			raw = append(raw, []byte(fmt.Sprintf("%s: %s\r\n", key, value))...)
		}
	}

	// Add the body delimiter

	raw = append(raw, []byte("\r\n")...)

	// Add the body

	body, err := ioutil.ReadAll(r.Response.Body)
	if err != nil {
		return nil, err
	}

	raw = append(raw, body...)

	return raw, nil
}
