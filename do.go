package request

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func Do(method, urli, postdata string, headers map[string]string, followlocation bool, proxy string) (*Response, error) {

	req, err := http.NewRequest(strings.ToUpper(method), urli, strings.NewReader(postdata))

	if err != nil {
		return nil, err
	}

	for v, k := range headers {
		req.Header.Add(v, k)
	}

	var a *url.URL

	if proxy != "" {
		a, _ = url.Parse(proxy)
	}

	client := &http.Client{
		Jar:     jar,
		Timeout: 60 * time.Second,
		Transport: &http.Transport{
			IdleConnTimeout:    60 * time.Second,
			DisableCompression: true,
			Proxy:              http.ProxyURL(a),
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
