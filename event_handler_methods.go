package iticapital

import (
	"context"
	"encoding/json"
)

func (eh *eventHandler) handleAddBar(ctx context.Context, queueName string, ch chan<- *AddBar) {
	reader := eh.readerFactory.NewQueueReader(queueName)

ReadLoop:
	for {
		select {
		case <-ctx.Done():
			break ReadLoop
		default:
			msg, err := reader.Read(ctx)

			if err != nil {
				eh.logger.Error("failed read from " + queueName, err)
				break ReadLoop
			}

			event := &AddBar{}

			err = json.Unmarshal(msg, event)
			if err != nil {
				eh.logger.Error("failed unmarshall message from " + queueName, err)
				continue
			}

			ch <- event
		}
	}
}

func (eh *eventHandler) handleAddPortfolio(ctx context.Context, queueName string, ch chan<- *AddPortfolio) {
	reader := eh.readerFactory.NewQueueReader(queueName)

ReadLoop:
	for {
		select {
		case <-ctx.Done():
			break ReadLoop
		default:
			msg, err := reader.Read(ctx)

			if err != nil {
				eh.logger.Error("failed read from " + queueName, err)
				break ReadLoop
			}

			event := &AddPortfolio{}

			err = json.Unmarshal(msg, event)
			if err != nil {
				eh.logger.Error("failed unmarshall message from " + queueName, err)
				continue
			}

			ch <- event
		}
	}
}

func (eh *eventHandler) handleAddSymbol(ctx context.Context, queueName string, ch chan<- *AddSymbol) {
	reader := eh.readerFactory.NewQueueReader(queueName)

ReadLoop:
	for {
		select {
		case <-ctx.Done():
			break ReadLoop
		default:
			msg, err := reader.Read(ctx)

			if err != nil {
				eh.logger.Error("failed read from " + queueName, err)
				break ReadLoop
			}

			event := &AddSymbol{}

			err = json.Unmarshal(msg, event)
			if err != nil {
				eh.logger.Error("failed unmarshall message from " + queueName, err)
				continue
			}

			ch <- event
		}
	}
}

func (eh *eventHandler) handleAddTick(ctx context.Context, queueName string, ch chan<- *AddTick) {
	reader := eh.readerFactory.NewQueueReader(queueName)

ReadLoop:
	for {
		select {
		case <-ctx.Done():
			break ReadLoop
		default:
			msg, err := reader.Read(ctx)

			if err != nil {
				eh.logger.Error("failed read from " + queueName, err)
				break ReadLoop
			}

			event := &AddTick{}

			err = json.Unmarshal(msg, event)
			if err != nil {
				eh.logger.Error("failed unmarshall message from " + queueName, err)
				continue
			}

			ch <- event
		}
	}
}

func (eh *eventHandler) handleAddTickHistory(ctx context.Context, queueName string, ch chan<- *AddTickHistory) {
	reader := eh.readerFactory.NewQueueReader(queueName)

ReadLoop:
	for {
		select {
		case <-ctx.Done():
			break ReadLoop
		default:
			msg, err := reader.Read(ctx)

			if err != nil {
				eh.logger.Error("failed read from " + queueName, err)
				break ReadLoop
			}

			event := &AddTickHistory{}

			err = json.Unmarshal(msg, event)
			if err != nil {
				eh.logger.Error("failed unmarshall message from " + queueName, err)
				continue
			}

			ch <- event
		}
	}
}

func (eh *eventHandler) handleAddTrade(ctx context.Context, queueName string, ch chan<- *AddTrade) {
	reader := eh.readerFactory.NewQueueReader(queueName)

ReadLoop:
	for {
		select {
		case <-ctx.Done():
			break ReadLoop
		default:
			msg, err := reader.Read(ctx)

			if err != nil {
				eh.logger.Error("failed read from " + queueName, err)
				break ReadLoop
			}

			event := &AddTrade{}

			err = json.Unmarshal(msg, event)
			if err != nil {
				eh.logger.Error("failed unmarshall message from " + queueName, err)
				continue
			}

			ch <- event
		}
	}
}

func (eh *eventHandler) handleOrderCancelFailed(ctx context.Context, queueName string, ch chan<- *OrderCancelFailed) {
	reader := eh.readerFactory.NewQueueReader(queueName)

ReadLoop:
	for {
		select {
		case <-ctx.Done():
			break ReadLoop
		default:
			msg, err := reader.Read(ctx)

			if err != nil {
				eh.logger.Error("failed read from " + queueName, err)
				break ReadLoop
			}

			event := &OrderCancelFailed{}

			err = json.Unmarshal(msg, event)
			if err != nil {
				eh.logger.Error("failed unmarshall message from " + queueName, err)
				continue
			}

			ch <- event
		}
	}
}

func (eh *eventHandler) handleOrderCancelSucceeded(ctx context.Context, queueName string, ch chan<- *OrderCancelSucceeded) {
	reader := eh.readerFactory.NewQueueReader(queueName)

ReadLoop:
	for {
		select {
		case <-ctx.Done():
			break ReadLoop
		default:
			msg, err := reader.Read(ctx)

			if err != nil {
				eh.logger.Error("failed read from " + queueName, err)
				break ReadLoop
			}

			event := &OrderCancelSucceeded{}

			err = json.Unmarshal(msg, event)
			if err != nil {
				eh.logger.Error("failed unmarshall message from " + queueName, err)
				continue
			}

			ch <- event
		}
	}
}

func (eh *eventHandler) handleOrderMoveFailed(ctx context.Context, queueName string, ch chan<- *OrderMoveFailed) {
	reader := eh.readerFactory.NewQueueReader(queueName)

ReadLoop:
	for {
		select {
		case <-ctx.Done():
			break ReadLoop
		default:
			msg, err := reader.Read(ctx)

			if err != nil {
				eh.logger.Error("failed read from " + queueName, err)
				break ReadLoop
			}

			event := &OrderMoveFailed{}

			err = json.Unmarshal(msg, event)
			if err != nil {
				eh.logger.Error("failed unmarshall message from " + queueName, err)
				continue
			}

			ch <- event
		}
	}
}

func (eh *eventHandler) handleOrderMoveSucceeded(ctx context.Context, queueName string, ch chan<- *OrderMoveSucceeded) {
	reader := eh.readerFactory.NewQueueReader(queueName)

ReadLoop:
	for {
		select {
		case <-ctx.Done():
			break ReadLoop
		default:
			msg, err := reader.Read(ctx)

			if err != nil {
				eh.logger.Error("failed read from " + queueName, err)
				break ReadLoop
			}

			event := &OrderMoveSucceeded{}

			err = json.Unmarshal(msg, event)
			if err != nil {
				eh.logger.Error("failed unmarshall message from " + queueName, err)
				continue
			}

			ch <- event
		}
	}
}

func (eh *eventHandler) handleOrderFailed(ctx context.Context, queueName string, ch chan<- *OrderFailed) {
	reader := eh.readerFactory.NewQueueReader(queueName)

ReadLoop:
	for {
		select {
		case <-ctx.Done():
			break ReadLoop
		default:
			msg, err := reader.Read(ctx)

			if err != nil {
				eh.logger.Error("failed read from " + queueName, err)
				break ReadLoop
			}

			event := &OrderFailed{}

			err = json.Unmarshal(msg, event)
			if err != nil {
				eh.logger.Error("failed unmarshall message from " + queueName, err)
				continue
			}

			ch <- event
		}
	}
}

func (eh *eventHandler) handleOrderSucceeded(ctx context.Context, queueName string, ch chan<- *OrderSucceeded) {
	reader := eh.readerFactory.NewQueueReader(queueName)

ReadLoop:
	for {
		select {
		case <-ctx.Done():
			break ReadLoop
		default:
			msg, err := reader.Read(ctx)

			if err != nil {
				eh.logger.Error("failed read from " + queueName, err)
				break ReadLoop
			}

			event := &OrderSucceeded{}

			err = json.Unmarshal(msg, event)
			if err != nil {
				eh.logger.Error("failed unmarshall message from " + queueName, err)
				continue
			}

			ch <- event
		}
	}
}

func (eh *eventHandler) handleSetMyClosePosition(ctx context.Context, queueName string, ch chan<- *SetMyClosePosition) {
	reader := eh.readerFactory.NewQueueReader(queueName)

ReadLoop:
	for {
		select {
		case <-ctx.Done():
			break ReadLoop
		default:
			msg, err := reader.Read(ctx)

			if err != nil {
				eh.logger.Error("failed read from " + queueName, err)
				break ReadLoop
			}

			event := &SetMyClosePosition{}

			err = json.Unmarshal(msg, event)
			if err != nil {
				eh.logger.Error("failed unmarshall message from " + queueName, err)
				continue
			}

			ch <- event
		}
	}
}

func (eh *eventHandler) handleSetMyOrder(ctx context.Context, queueName string, ch chan<- *SetMyOrder) {
	reader := eh.readerFactory.NewQueueReader(queueName)

ReadLoop:
	for {
		select {
		case <-ctx.Done():
			break ReadLoop
		default:
			msg, err := reader.Read(ctx)

			if err != nil {
				eh.logger.Error("failed read from " + queueName, err)
				break ReadLoop
			}

			event := &SetMyOrder{}

			err = json.Unmarshal(msg, event)
			if err != nil {
				eh.logger.Error("failed unmarshall message from " + queueName, err)
				continue
			}

			ch <- event
		}
	}
}

func (eh *eventHandler) handleSetMyTrade(ctx context.Context, queueName string, ch chan<- *SetMyTrade) {
	reader := eh.readerFactory.NewQueueReader(queueName)

ReadLoop:
	for {
		select {
		case <-ctx.Done():
			break ReadLoop
		default:
			msg, err := reader.Read(ctx)

			if err != nil {
				eh.logger.Error("failed read from " + queueName, err)
				break ReadLoop
			}

			event := &SetMyTrade{}

			err = json.Unmarshal(msg, event)
			if err != nil {
				eh.logger.Error("failed unmarshall message from " + queueName, err)
				continue
			}

			ch <- event
		}
	}
}

func (eh *eventHandler) handleSetPortfolio(ctx context.Context, queueName string, ch chan<- *SetPortfolio) {
	reader := eh.readerFactory.NewQueueReader(queueName)

ReadLoop:
	for {
		select {
		case <-ctx.Done():
			break ReadLoop
		default:
			msg, err := reader.Read(ctx)

			if err != nil {
				eh.logger.Error("failed read from " + queueName, err)
				break ReadLoop
			}

			event := &SetPortfolio{}

			err = json.Unmarshal(msg, event)
			if err != nil {
				eh.logger.Error("failed unmarshall message from " + queueName, err)
				continue
			}

			ch <- event
		}
	}
}

func (eh *eventHandler) handleUpdateOrder(ctx context.Context, queueName string, ch chan<- *UpdateOrder) {
	reader := eh.readerFactory.NewQueueReader(queueName)

ReadLoop:
	for {
		select {
		case <-ctx.Done():
			break ReadLoop
		default:
			msg, err := reader.Read(ctx)

			if err != nil {
				eh.logger.Error("failed read from " + queueName, err)
				break ReadLoop
			}

			event := &UpdateOrder{}

			err = json.Unmarshal(msg, event)
			if err != nil {
				eh.logger.Error("failed unmarshall message from " + queueName, err)
				continue
			}

			ch <- event
		}
	}
}

func (eh *eventHandler) handleUpdatePosition(ctx context.Context, queueName string, ch chan<- *UpdatePosition) {
	reader := eh.readerFactory.NewQueueReader(queueName)

ReadLoop:
	for {
		select {
		case <-ctx.Done():
			break ReadLoop
		default:
			msg, err := reader.Read(ctx)

			if err != nil {
				eh.logger.Error("failed read from " + queueName, err)
				break ReadLoop
			}

			event := &UpdatePosition{}

			err = json.Unmarshal(msg, event)
			if err != nil {
				eh.logger.Error("failed unmarshall message from " + queueName, err)
				continue
			}

			ch <- event
		}
	}
}

func (eh *eventHandler) handleUpdateBidAsks(ctx context.Context, queueName string, ch chan<- *UpdateBidAsks) {
	reader := eh.readerFactory.NewQueueReader(queueName)

ReadLoop:
	for {
		select {
		case <-ctx.Done():
			break ReadLoop
		default:
			msg, err := reader.Read(ctx)

			if err != nil {
				eh.logger.Error("failed read from " + queueName, err)
				break ReadLoop
			}

			event := &UpdateBidAsks{}

			err = json.Unmarshal(msg, event)
			if err != nil {
				eh.logger.Error("failed unmarshall message from " + queueName, err)
				continue
			}

			ch <- event
		}
	}
}

func (eh *eventHandler) handleUpdateQuote(ctx context.Context, queueName string, ch chan<- *UpdateQuote) {
	reader := eh.readerFactory.NewQueueReader(queueName)

ReadLoop:
	for {
		select {
		case <-ctx.Done():
			break ReadLoop
		default:
			msg, err := reader.Read(ctx)

			if err != nil {
				eh.logger.Error("failed read from " + queueName, err)
				break ReadLoop
			}

			event := &UpdateQuote{}

			err = json.Unmarshal(msg, event)
			if err != nil {
				eh.logger.Error("failed unmarshall message from " + queueName, err)
				continue
			}

			ch <- event
		}
 	}
}
