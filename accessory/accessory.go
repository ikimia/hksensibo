package accessory

import (
	"hash/fnv"

	hca "github.com/brutella/hc/accessory"
	"github.com/ikimia/hksensibo/accessory/fan"
	"github.com/ikimia/hksensibo/accessory/thermostat"
	"github.com/ikimia/hksensibo/api"
	"github.com/ikimia/hksensibo/pod"
)

func generateID(s string) uint64 {
	h := fnv.New64()
	h.Write([]byte(s))
	return h.Sum64()
}

// NewAccessory creates an accessory for the given pod id
func NewAccessory(podID string, api *api.API) *hca.Accessory {
	pa := pod.NewPodAccessor(podID, api)

	info := hca.Info{Name: "Pod", ID: generateID(podID)}
	acc := hca.New(info, hca.TypeAirConditioner)

	acc.Info.Name.OnValueRemoteGet(func() string {
		return pa.GetPod().Room.Name
	})

	acc.AddService(thermostat.NewService(pa))
	acc.AddService(fan.NewService(pa))

	return acc
}
