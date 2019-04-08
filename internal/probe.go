package internal

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

// Probe is
type Probe interface {
	Resource() string
	Measure() (*Measurement, error)
}

// MemProbe allows to measure memory usage
type MemProbe struct{}

// Resource describes probe
func (p MemProbe) Resource() string { return "memory" }

// Measure is used to collect one measurement
func (p MemProbe) Measure() (*Measurement, error) {
	v, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}
	return NewMeasurement(p.Resource(), v.UsedPercent), nil
}

// CPUProbe allows to measure CPU usage
type CPUProbe struct{}

// Resource describes probe
func (p CPUProbe) Resource() string { return "cpu" }

// Measure is used to collect one measurement
func (p CPUProbe) Measure() (*Measurement, error) {
	v, err := cpu.Percent(0, false)
	if err != nil {
		return nil, err
	}
	return NewMeasurement(p.Resource(), v[0]), nil
}
