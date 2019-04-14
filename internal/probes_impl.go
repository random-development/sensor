package internal

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

// MemProbe allows to measure memory usage
type MemProbe struct{}

// MetricName describes probe
func (p MemProbe) MetricName() string { return "memory" }

// Measure is used to collect one measurement
func (p MemProbe) Measure() (Measurement, error) {
	v, err := mem.VirtualMemory()
	if err != nil {
		return Measurement{}, err
	}
	return NewMeasurement(p.MetricName(), v.UsedPercent), nil
}

// CPUProbe allows to measure CPU usage
type CPUProbe struct{}

// MetricName describes probe
func (p CPUProbe) MetricName() string { return "cpu" }

// Measure is used to collect one measurement
func (p CPUProbe) Measure() (Measurement, error) {
	v, err := cpu.Percent(0, false)
	if err != nil {
		return Measurement{}, err
	}
	return NewMeasurement(p.MetricName(), v[0]), nil
}
