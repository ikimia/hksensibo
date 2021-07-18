package fan

import (
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/hc/service"
	"github.com/ikimia/hksensibo/mapping"
	"github.com/ikimia/hksensibo/pod"
	"github.com/ikimia/hksensibo/sensibo"
)

var fanLevel = mapping.Float64(map[float64]string{20: "quiet", 40: "low", 60: "medium", 80: "high", 100: "auto"})
var swing = mapping.Int(map[int]string{0: "stopped", 1: "rangeFull"})

func NewService(pa *pod.PodAccessor) *service.Service {
	fan := service.NewFanV2()

	fan.Active.SetValue(1)

	rotationSpeed := characteristic.NewRotationSpeed()
	rotationSpeed.SetStepValue(20)
	fan.AddCharacteristic(rotationSpeed.Characteristic)

	rotationSpeed.OnValueRemoteGet(func() float64 {
		return fanLevel.Values[pa.GetPod().ACState.FanLevel]
	})

	swingMode := characteristic.NewSwingMode()
	fan.AddCharacteristic(swingMode.Characteristic)

	rotationSpeed.OnValueRemoteUpdate(func(v float64) {
		pa.UpdateState(func(s *sensibo.ACState) { s.FanLevel = fanLevel.Names[v] })
	})

	swingMode.OnValueRemoteUpdate(func(v int) {
		pa.UpdateState(func(s *sensibo.ACState) { s.Swing = swing.Names[v] })
	})

	swingMode.OnValueRemoteGet((func() int {
		return swing.Values[pa.GetPod().ACState.Swing]
	}))

	return fan.Service
}
