# raw_http_utils

Golang doesn't currently provide a way to retrieve the raw HTTP request and response in a byte slice from the net/http library. This is a simple library to provide those objects in a byte slice format.

To use this library simple add `"github.com/wdahlenburg/raw_http_utils/utils"` to your imports. You can then pass the net/http `Response` object to the `RawHttpResponse` struct like so:

```go
resp, _ := http.Get("https://www.google.com")
raw := &utils.RawHttpResponse{
   Response: resp,
}

result, err := raw.GetRawResponse()
if err != nil {
   panic(err)
}
```

* Note that net/http uses a map for the HTTP headers. This results in the header order changing despite the order that the server sent. The headers have been sorted in this library to provide the most consistent output.
