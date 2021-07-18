package pod

import (
	"log"
	"sync"
	"time"

	"github.com/ikimia/hksensibo/api"
	"github.com/ikimia/hksensibo/sensibo"
)

// PodAccessor allows getting a pod and updating a pod's state
type PodAccessor struct {
	podID  string
	api    *api.API
	expiry time.Time
	pod    *sensibo.Pod
	gmu    sync.Mutex
	umu    sync.Mutex
}

// NewPodAccessor creates a PodAccessor for the given pod id
func NewPodAccessor(podID string, api *api.API) *PodAccessor {
	return &PodAccessor{podID: podID, api: api}
}

func (pa *PodAccessor) fetchPod() error {
	pa.gmu.Lock()
	defer pa.gmu.Unlock()

	if pa.pod == nil || time.Now().After(pa.expiry) {
		var err error
		pa.pod, err = pa.api.GetPod(pa.podID)
		if err != nil {
			return err
		}
		pa.expiry = time.Now().Add(5 * time.Second)
	}
	return nil
}

func (pa *PodAccessor) GetPod() *sensibo.Pod {
	err := pa.fetchPod()
	if err != nil {
		log.Printf("Error getting pod: %s\n", err)
	}
	return pa.pod
}

func (pa *PodAccessor) UpdateState(f func(*sensibo.ACState)) {
	pa.umu.Lock()
	defer pa.umu.Unlock()

	updatedState := pa.pod.ACState
	f(&updatedState)
	if pa.pod.ACState == updatedState {
		return
	}
	newState, err := pa.api.SetACState(pa.podID, &updatedState)
	if err != nil {
		log.Printf("Error updating state: %s\n", err)
		return
	}
	pa.pod.ACState = *newState
}
