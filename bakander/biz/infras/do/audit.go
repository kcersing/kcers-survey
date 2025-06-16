package do

import (
	"time"
)

// Audit 审计日志
type Audit interface {
	LogEvent(event *AuditEvent) error
}
type AuditEvent struct {
	OrderId     int64
	Action      string
	Description string
	Timestamp   time.Time
}
