package internal_test

import (
	"testing"
	"time"

	"github.com/random-development/sensor/internal"

	"github.com/golang/mock/gomock"
	"github.com/random-development/sensor/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type PublisherTestSuite struct {
	suite.Suite

	mockCtrl *gomock.Controller

	publisher *mocks.MockPublisher
	broker    *mocks.MockBroker
	done      chan bool
	measCh    chan interface{}
	topic     string
}

func (suite *PublisherTestSuite) SetupTest() {
	suite.mockCtrl = gomock.NewController(suite.T())

	suite.publisher = mocks.NewMockPublisher(suite.mockCtrl)
	suite.broker = mocks.NewMockBroker(suite.mockCtrl)
	suite.done = make(chan bool)
	suite.measCh = make(chan interface{})
	suite.topic = "topic"
}

func (suite *PublisherTestSuite) TearDownTest() {
	suite.mockCtrl.Finish()
}

func (suite *PublisherTestSuite) TestPublisherRun_ShouldPublishReceivedMeasurements() {
	m := internal.Measurement{}
	suite.broker.EXPECT().Sub(suite.topic).Times(1).Return(suite.measCh)
	suite.publisher.EXPECT().Publish(m).Times(1)

	internal.RunPublisher(suite.topic, suite.publisher, suite.broker, suite.done)
	suite.measCh <- m
}

func (suite *PublisherTestSuite) TestPublisherRun_ShouldExitOnDone() {
	suite.broker.EXPECT().Sub(suite.topic).Times(1).Return(suite.measCh)

	internal.RunPublisher(suite.topic, suite.publisher, suite.broker, suite.done)
	suite.done <- true

	timer := time.NewTimer(time.Millisecond)
	select {
	case suite.measCh <- internal.Measurement{}:
		assert.Fail(suite.T(), "Goroutine handling publisher shouldn't be running after done")
	case <-timer.C:
	}
}

func TestPublisherTestSuite(t *testing.T) {
	suite.Run(t, new(PublisherTestSuite))
}
