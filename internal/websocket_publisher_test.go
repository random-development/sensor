package internal_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/random-development/sensor/internal"
	"github.com/random-development/sensor/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type WebSocketPublisherTestSuite struct {
	suite.Suite

	sut      internal.WebSocketPublisher
	mockCtrl *gomock.Controller
	dialer   *mocks.MockDialer
	conn     *mocks.MockConn
	ch       chan interface{}
	done     chan bool

	topic string
	url   string
}

func (suite *WebSocketPublisherTestSuite) SetupTest() {
	suite.mockCtrl = gomock.NewController(suite.T())
	suite.done = make(chan bool)
	suite.topic = "topic"
	suite.url = "ws://websocket/endpoint"
	suite.sut = buildPublisher(suite)
	suite.sut.Run(suite.done)
}

func buildPublisher(suite *WebSocketPublisherTestSuite) internal.WebSocketPublisher {
	suite.dialer = mocks.NewMockDialer(suite.mockCtrl)
	suite.conn = mocks.NewMockConn(suite.mockCtrl)
	suite.ch = make(chan interface{})

	suite.dialer.EXPECT().
		Dial(suite.url, nil).
		Times(1).
		Return(suite.conn, nil, nil)

	pub, err := internal.MakeWebSocketPublisher(suite.topic, suite.url, suite.dialer, suite.ch)
	assert.NoError(suite.T(), err)
	return pub
}

func (suite *WebSocketPublisherTestSuite) TearDownTest() {
	suite.done <- true
	suite.mockCtrl.Finish()
}

func (suite *WebSocketPublisherTestSuite) Test_WhenReceiveMeasurement_ShouldPublish() {
	meas := internal.Measurement{Resource: suite.topic, Time: 1, Value: 1.0}
	suite.conn.EXPECT().WriteJSON(meas).Return(nil)
	suite.ch <- meas
}

func TestMakeWebSocketPublisher(t *testing.T) {
	suite.Run(t, new(WebSocketPublisherTestSuite))
}
