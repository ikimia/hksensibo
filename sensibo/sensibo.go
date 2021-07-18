package sensibo

// ACState describes the remote state of the AC
type ACState struct {
	On                bool   `json:"on"`
	FanLevel          string `json:"fanLevel"`
	TemperatureUnit   string `json:"temperatureUnit"`
	TargetTemperature int    `json:"targetTemperature"`
	Mode              string `json:"mode"`
	Swing             string `json:"swing"`
}

// Measurements holds the currently measured temperature and humidity
type Measurements struct {
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
}

// Room holds the name and icon where the pod is situated
type Room struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
}

// Pod represents a Sensibo AC unit
type Pod struct {
	ID           string       `json:"id"`
	ACState      ACState      `json:"acState"`
	Measurements Measurements `json:"measurements"`
	Room         Room         `json:"room"`
}
