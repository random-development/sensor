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

	topic string
	url   string
}

func (suite *WebSocketPublisherTestSuite) SetupTest() {
	suite.mockCtrl = gomock.NewController(suite.T())
	suite.url = "ws://websocket/endpoint"
	suite.sut = buildPublisher(suite)
}

func buildPublisher(suite *WebSocketPublisherTestSuite) internal.WebSocketPublisher {
	suite.dialer = mocks.NewMockDialer(suite.mockCtrl)
	suite.conn = mocks.NewMockConn(suite.mockCtrl)

	suite.dialer.EXPECT().
		Dial(suite.url, nil).
		Times(1).
		Return(suite.conn, nil, nil)

	pub, err := internal.MakeWebSocketPublisher(suite.url, suite.dialer)
	assert.NoError(suite.T(), err)
	return pub
}

func (suite *WebSocketPublisherTestSuite) TearDownTest() {
	suite.mockCtrl.Finish()
}

func (suite *WebSocketPublisherTestSuite) Test_PublishCalled_ShouldWriteJsonToConnection() {
	meas := internal.Measurement{Resource: suite.topic, Time: 1, Value: 1.0}
	suite.conn.EXPECT().WriteJSON(meas).Return(nil)
	suite.sut.Publish(meas)
}

func TestMakeWebSocketPublisher(t *testing.T) {
	suite.Run(t, new(WebSocketPublisherTestSuite))
}
