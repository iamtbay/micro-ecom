package handlersPackage

import (
	"bytes"
	"io"
	"net/http"
	"time"
)

// FORWARD REQUEST
func forwardRequest(targetServiceURL string, clientReq *http.Request) (*http.Response, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	var reqBody []byte = nil
	var err error
	//req body
	if clientReq.Body != nil {
		reqBody, err = io.ReadAll(clientReq.Body)
		if err != nil {
			return nil, err
		}
	}

	//create a new request
	req, err := http.NewRequest(clientReq.Method, targetServiceURL, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	//copy headers
	for key, values := range clientReq.Header {
		for _, val := range values {
			req.Header.Set(key, val)
		}
	}

	//send req to target
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil

}
