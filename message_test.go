package simplysocket

import (
	"testing"
)

func TestMessage_Encode(t *testing.T) {
	message := &Message{
		Action: "test_action",
		MessageBody: map[string]interface{}{
			"key": "value",
		},
		IsTargetClient: true,
		Target:         "target_client",
		Sender:         "sender_client",
	}

	encoded := message.Encode()
	expected := `{"action":"test_action","message_body":{"key":"value"},"IsTargetClient":true,"target":"target_client","sender":"sender_client"}`

	if string(encoded) != expected {
		t.Errorf("Expected %s, got %s", expected, string(encoded))
	}
}
