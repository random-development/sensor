package internal

import (
	"fmt"
	"time"
)

// Measurement is a struct sent between collectors and publishers
type Measurement struct {
	Resource string
	Time     int64
	Value    float64
}

// NewMeasurement creates measurement with current time
func NewMeasurement(resource string, value float64) *Measurement {
	return &Measurement{
		Resource: resource,
		Time:     time.Now().UTC().Unix(),
		Value:    value}
}

func (m *Measurement) String() string {
	return fmt.Sprintf("[%d]: %s=%f", m.Time, m.Resource, m.Value)
}
