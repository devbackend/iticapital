package iticapital

type eventHandler struct {
	readerFactory QueueReaderFactory
	logger        Logger
}