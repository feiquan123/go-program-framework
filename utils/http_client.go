package utils

import (
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// HTTPClient : http client
type HTTPClient struct {
	Client http.Client
}

// NewHTTPClient : creaet a http client
func NewHTTPClient(connTimeout time.Duration, readTimeout time.Duration) *HTTPClient {
	client := http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				c, err := net.DialTimeout(netw, addr, connTimeout)
				if err != nil {
					return nil, err
				}
				c.SetDeadline(time.Now().Add(readTimeout))
				return c, nil
			},
		},
	}
	return &HTTPClient{Client: client}
}

// SendGetRequest : get requset
func (h *HTTPClient) SendGetRequest(httpURL string, postParams map[string]string) (string, error) {
	u, err := url.Parse(httpURL)
	if err != nil {
		log.Println(err)
	}
	q := u.Query()
	for key, value := range postParams {
		q.Set(key, value)
	}

	u.RawQuery = q.Encode()

	resp, reqErr := h.Client.Get(u.String())

	if reqErr != nil {
		return "", reqErr
	}

	defer resp.Body.Close()

	data, respErr := ioutil.ReadAll(resp.Body)

	if respErr != nil {
		return "", respErr
	}
	return string(data), nil
}

// SendPostRequest : post request
func (h *HTTPClient) SendPostRequest(httpURL string, headers map[string]string, body string) (string, error) {

	req, _ := http.NewRequest("POST", httpURL, strings.NewReader(body))

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, reqErr := h.Client.Do(req)

	if reqErr != nil {
		return "", reqErr
	}

	defer resp.Body.Close()

	data, respErr := ioutil.ReadAll(resp.Body)

	if respErr != nil {
		return "", respErr
	}

	return string(data), nil
}

// Close : http client close
func (h *HTTPClient) Close() {
	h.Client.CloseIdleConnections()
}
