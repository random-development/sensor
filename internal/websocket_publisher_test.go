package internal_test

import (
	"testing"

	"github.com/golang/mock/gomock"
    "github.com/random-development/sensor/mocks"
    "github.com/random-development/sensor/internal"
)

func TestMakeWebSocketPublisher(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dialer := mocks.NewMockDialer(ctrl)
    conn := mocks.NewMockConn(ctrl)

    topic := "topic"
    url := "ws://endpoint"

    dialer.EXPECT().
      Dial(url, nil).
      Times(1).
      Return(conn, nil, nil)

    internal.MakeWebSocketPublisher(topic, url, dialer)
}
