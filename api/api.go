package api

import (
	"github.com/ikimia/hksensibo/sensibo"
)

type qs map[string]string

type setACStateResponse struct {
	Status            string          `json:"status"`
	Reason            string          `json:"reason"`
	ACState           sensibo.ACState `json:"acState"`
	ChangedProperties []string        `json:"changedProperties"`
	ID                string          `json:"id"`
	FailureReason     string          `json:"failureReason"`
}

// API handle for Sensibo
type API struct {
	apiKey string
}

// NewAPI returns a new Sensibo API configured with the given api key
func NewAPI(apiKey string) *API {
	return &API{apiKey}
}

// GetPod fetches the configuration for the pod of the given pod id
func (api *API) GetPod(podID string) (*sensibo.Pod, error) {
	url := api.makeURL("pods/"+podID, qs{"fields": "id,acState,measurements,room"})
	req := makeGetRequest(url)
	var res sensibo.Pod
	err := doRequest(req, &res)
	if err != nil {
		return nil, err
	}
	return &res, err
}

// GetAllPodIDs fetches the ids for all the pods related to the API key
func (api *API) GetAllPodIDs() ([]string, error) {
	url := api.makeURL("users/me/pods", qs{"fields": "id"})
	req := makeGetRequest(url)
	var res []struct {
		ID string `json:"id"`
	}
	err := doRequest(req, &res)
	ids := make([]string, 0, len(res))
	for _, o := range res {
		ids = append(ids, o.ID)
	}
	return ids, err
}

// SetACState sets the remote AC state for the given device id
func (api *API) SetACState(deviceID string, newState *sensibo.ACState) (*sensibo.ACState, error) {
	url := api.makeURL("pods/"+deviceID+"/acStates", nil)
	req, err := makePostRequest(url, map[string]interface{}{"acState": newState})
	if err != nil {
		return nil, err
	}
	res := &setACStateResponse{}
	err = doRequest(req, res)
	return &res.ACState, err
}
