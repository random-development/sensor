package internal_test

import (
	"encoding/json"
	"testing"

	"github.com/random-development/sensor/internal"
	"github.com/stretchr/testify/assert"
)

func TestMeasurementJsonEncoding(t *testing.T) {
	m := internal.Measurement{Resource: "cpu", Time: 1, Value: 2.3}
	encoded, err := json.Marshal(m)
	if assert.NoError(t, err) {
		assert.Equal(t, string(encoded), "{\"time\":1,\"value\":2.3}")
	}
}
