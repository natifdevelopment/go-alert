package alert

import "time"

// Payload is the alert message published to Kafka.
type Payload struct {
	Timestamp    time.Time `json:"timestamp"`
	Level        string    `json:"level"`
	Service      string    `json:"service"`
	Version      string    `json:"version"`
	Environment  string    `json:"environment"`
	Hostname     string    `json:"hostname"`
	TraceID      string    `json:"trace_id,omitempty"`
	RequestID    string    `json:"request_id,omitempty"`
	Method       string    `json:"method,omitempty"`
	Path         string    `json:"path,omitempty"`
	StatusCode   int       `json:"status_code,omitempty"`
	ErrorMessage string    `json:"error_message"`
	StackTrace   string    `json:"stack_trace,omitempty"`
}
