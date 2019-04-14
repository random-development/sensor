package internal_test

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/random-development/sensor/internal"
	"github.com/random-development/sensor/mocks"
)

func TestProbeRun(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	probe := mocks.NewMockProbe(mockCtrl)
	broker := mocks.NewMockBroker(mockCtrl)

	m := internal.Measurement{}
	r := "resource"
	probe.EXPECT().MetricName().AnyTimes().Return(r)
	probe.EXPECT().Measure().MinTimes(5).Return(m, nil)
	broker.EXPECT().Pub(m, r).MinTimes(5)

	timer := time.NewTimer(time.Microsecond)
	internal.RunProbe(probe, broker, 5*time.Microsecond)
	<-timer.C
}
