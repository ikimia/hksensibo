package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type apiResponse struct {
	Status  string      `json:"status"`
	Reason  string      `json:"reason"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

func doRequest(req *http.Request, result interface{}) error {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	res := &apiResponse{Result: result}
	err = json.NewDecoder(resp.Body).Decode(res)

	if res.Status == "failed" {
		return fmt.Errorf("%+v", res)
	}

	return err
}
