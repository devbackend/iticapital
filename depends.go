package iticapital

import "context"

type QueueMessage []byte

type QueueReader interface {
	Read(context.Context) (QueueMessage, error)
}

type QueueReaderFactory interface {
	NewQueueReader(queue string) QueueReader
}

type Logger interface {
	Error(description string, err error)
}
