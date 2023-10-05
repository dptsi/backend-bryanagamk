package common

import "time"

// The minimal interface DomainEvent, implemented by all Events, ensures support of an occurredOn() accessor. It enforces a basic contract for all Events:
// Implementing Domain-Driven Design, Vaughn Vernon

type Event interface {
	OccuredOn() time.Time
	JSON() ([]byte, error)
}
