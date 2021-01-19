package templater

import (
	"net/http"

	"moul.io/http2curl"
)

func CURLFromRequest(req *http.Request) (s string, err error) {
	if req == nil {
		return
	}

	curlstr, err := http2curl.GetCurlCommand(req)
	if err != nil {
		return
	}

	s = curlstr.String()

	return
}
