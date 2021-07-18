package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

func (api *API) makeURL(path string, params qs) *url.URL {
	base, _ := url.Parse("https://home.sensibo.com/api/v2/")
	base.Path += path

	q := base.Query()
	q.Add("apiKey", api.apiKey)
	for k, v := range params {
		q.Add(k, v)
	}
	base.RawQuery = q.Encode()

	return base
}

func makeGetRequest(url *url.URL) *http.Request {
	return &http.Request{
		Method: "GET",
		Header: http.Header{"accept": {"application/json"}},
		URL:    url,
	}
}

func makePostRequest(url *url.URL, body interface{}) (*http.Request, error) {
	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return &http.Request{
		Method: "POST",
		URL:    url,
		Header: http.Header{"Accept": {"application/json"}, "Content-Type": {"application/json"}},
		Body:   ioutil.NopCloser(bytes.NewReader(b)),
	}, nil
}
