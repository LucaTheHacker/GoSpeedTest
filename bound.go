package GoSpeedTest

import "math"

// GetMegaBit returns Measure speed in MegaBits, rounded at the second decimal place
func (m *Measure) GetMegaBit() float64 {
	return math.Floor(float64(m.Bandwidth/125000*100)) / 100
}

// GetMegaByte returns Measure speed in MegaBytes, rounded at the second decimal place
func (m *Measure) GetMegaByte() float64 {
	return math.Floor(float64(m.Bandwidth/1000000*100)) / 100
}

// GetGigaBit returns Measure speed in GigaBits, rounded at the second decimal place
func (m *Measure) GetGigaBit() float64 {
	return math.Floor(float64(m.Bandwidth/125000000*100)) / 100
}
