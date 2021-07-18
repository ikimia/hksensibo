package thermostat

import (
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/hc/service"
	"github.com/ikimia/hksensibo/mapping"
	"github.com/ikimia/hksensibo/pod"
	"github.com/ikimia/hksensibo/sensibo"
)

var mode = mapping.Int(map[int]string{1: "heat", 2: "cool", 3: "auto"})

func NewService(pa *pod.PodAccessor) *service.Service {
	thermostat := service.NewThermostat()
	thermostat.TargetTemperature.SetStepValue(1)

	humidity := characteristic.NewCurrentRelativeHumidity()
	thermostat.AddCharacteristic(humidity.Characteristic)

	humidity.OnValueRemoteGet(func() float64 {
		return pa.GetPod().Measurements.Humidity
	})

	thermostat.TargetHeatingCoolingState.OnValueRemoteUpdate(func(v int) {
		pa.UpdateState(func(s *sensibo.ACState) {
			if v == 0 {
				s.On = false
			} else {
				s.On = true
				s.Mode = mode.Names[v]
			}
		})
	})

	thermostat.TargetHeatingCoolingState.OnValueRemoteGet(func() int {
		p := pa.GetPod()
		if p.ACState.On {
			return mode.Values[p.ACState.Mode]
		}
		return 0
	})

	thermostat.TargetTemperature.OnValueRemoteUpdate(func(v float64) {
		pa.UpdateState(func(s *sensibo.ACState) { s.TargetTemperature = int(v) })
	})

	thermostat.TargetTemperature.OnValueRemoteGet(func() float64 {
		return float64(pa.GetPod().ACState.TargetTemperature)
	})

	thermostat.CurrentTemperature.OnValueRemoteGet(func() float64 {
		return pa.GetPod().Measurements.Temperature
	})

	return thermostat.Service
}
