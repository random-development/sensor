package internal_test

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/random-development/sensor/internal"
	"github.com/random-development/sensor/mocks"
)

func TestCollectorRun(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	probe := mocks.NewMockProbe(mockCtrl)
	broker := mocks.NewMockBroker(mockCtrl)

	sut := internal.MakeCollector(probe, broker)

	m := internal.Measurement{}
	r := "resource"
	probe.EXPECT().MetricName().AnyTimes().Return(r)
	probe.EXPECT().Measure().Times(5).Return(m, nil)
	broker.EXPECT().Pub(m, r).Times(5)

	timer := time.NewTimer(time.Microsecond)
	sut.Run(5 * time.Microsecond)
	<-timer.C
}
