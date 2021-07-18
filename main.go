package main

import (
	"log"

	"github.com/brutella/hc"
	hca "github.com/brutella/hc/accessory"

	"github.com/ikimia/hksensibo/accessory"
	"github.com/ikimia/hksensibo/api"
	"github.com/ikimia/hksensibo/config"
)

func initAccessories(api *api.API) []*hca.Accessory {
	ids, err := api.GetAllPodIDs()
	if err != nil {
		log.Fatalf("Error initializing accessories: %s\n", err)
	}
	as := make([]*hca.Accessory, 0, len(ids))
	for _, id := range ids {
		as = append(as, accessory.NewAccessory(id, api))
	}
	return as
}

func main() {
	c := config.FromEnv()
	api := api.NewAPI(c.APIKey)
	as := initAccessories(api)

	config := hc.Config{Pin: c.Pin, StoragePath: c.StoragePath}
	bridge := hca.NewBridge(hca.Info{Name: "Sensibo", ID: 0})
	t, err := hc.NewIPTransport(config, bridge.Accessory, as...)
	if err != nil {
		log.Fatalf("Error initializing HomeKit bridge: %s\n", err)
	}

	hc.OnTermination(func() {
		<-t.Stop()
	})

	t.Start()
}
