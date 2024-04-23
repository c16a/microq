package broker

import (
	"testing"
)

type TestingWebSocketConnection struct {
}

func (t *TestingWebSocketConnection) WriteMessage(data []byte) error {
	return nil
}

func TestConnectedClient_Subscribe_Unsubscribe(t *testing.T) {
	client := NewConnectedClient(&TestingWebSocketConnection{}, "client-1")

	client.SubscribeToPattern("t1.*", "g1")
	client.UnsubscribeFromPattern("t1.accounts")

	s1 := client.GetEligibility("t1.payments")
	if s1 == nil {
		t.Fatal("GetEligibility returned nil")
	}

	s2 := client.GetEligibility("t1.accounts")
	if s2 != nil {
		t.Fatal("GetEligibility returned non-nil eligibility")
	}

	client.UnsubscribeFromPattern("t1.payments")

	s3 := client.GetEligibility("t1.payments")
	if s3 != nil {
		t.Fatal("GetEligibility returned non-nil eligibility")
	}
}
