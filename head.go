package request

import (
	"io/ioutil"
	"net/http"
	"time"
)

func Head(url string, headers map[string]string, followlocation bool) (*Response, error) {

	req, err := http.NewRequest("HEAD", url, nil)

	if err != nil {
		return nil, err
	}

	for v, k := range headers {
		req.Header.Add(v, k)
	}

	client := &http.Client{
		Jar:     jar,
		Timeout: 60 * time.Second,
		Transport: &http.Transport{
			IdleConnTimeout:    60 * time.Second,
		},
	}

	if !followlocation {
		client.CheckRedirect = func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse }
	}

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	return &Response{Status: res.Status, StatusCode: res.StatusCode, ContentLength: res.ContentLength, Header: res.Header, Body: string(body)}, nil

}
