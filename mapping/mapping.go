package mapping

// Float64Mapping holds a bidirectional float64-string mapping
type Float64Mapping struct {
	Names  map[float64]string
	Values map[string]float64
}

// Float64 creates a bidirectional float64-string mapping
func Float64(valueToName map[float64]string) *Float64Mapping {
	return &Float64Mapping{valueToName, reverseFloat64Map(valueToName)}
}

func reverseFloat64Map(float64ToString map[float64]string) map[string]float64 {
	reversed := make(map[string]float64)
	for k, v := range float64ToString {
		reversed[v] = k
	}
	return reversed
}

// IntMapping holds a bidirectional int-string mapping
type IntMapping struct {
	Names  map[int]string
	Values map[string]int
}

// Int creates a bidirectional float64-string mapping
func Int(valueToName map[int]string) *IntMapping {
	return &IntMapping{valueToName, reverseIntMap(valueToName)}
}

func reverseIntMap(intToString map[int]string) map[string]int {
	reversed := make(map[string]int)
	for k, v := range intToString {
		reversed[v] = k
	}
	return reversed
}
