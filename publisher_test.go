package alert

import (
	"encoding/json"
	"testing"
	"time"
)

func TestPayload_Marshal(t *testing.T) {
	p := Payload{
		Timestamp:    time.Date(2026, 8, 8, 0, 0, 0, 0, time.UTC),
		Level:        "ERROR",
		Service:      "test",
		ErrorMessage: "boom",
	}

	b, err := json.Marshal(p)
	if err != nil {
		t.Fatalf("marshal payload: %v", err)
	}

	if !json.Valid(b) {
		t.Fatalf("invalid json: %s", string(b))
	}

	var decoded Payload
	if err := json.Unmarshal(b, &decoded); err != nil {
		t.Fatalf("unmarshal payload: %v", err)
	}

	if decoded.Level != p.Level || decoded.Service != p.Service {
		t.Fatalf("decoded mismatch: got %+v", decoded)
	}
}
