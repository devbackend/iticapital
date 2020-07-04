package iticapital

import "context"

// Client for calling gRPC ITICapital
type Client struct {
	grpc ITICapitalAPIClient
	eh   *eventHandler
}

// Create client
func NewAPIClient(grpc ITICapitalAPIClient, rf QueueReaderFactory, logger Logger) *Client {
	eh := &eventHandler{rf, logger}

	return &Client{grpc, eh}
}

// Getting portfolio
func (c *Client) GetPortfolioList(ctx context.Context) (<-chan *AddPortfolio, error) {
	resp, err := c.grpc.GetPortfolioList(ctx, &Nothing{})
	if err != nil {
		return nil, err
	}

	ch := make(chan *AddPortfolio)

	go c.eh.handleAddPortfolio(ctx, resp.Queue, ch)

	return ch, nil
}

// Get all available instruments
func (c *Client) GetSymbols(ctx context.Context) (<-chan *AddSymbol, error) {
	resp, err := c.grpc.GetSymbols(ctx, &Nothing{})
	if err != nil {
		return nil, err
	}

	ch := make(chan *AddSymbol)

	go c.eh.handleAddSymbol(ctx, resp.Queue, ch)

	return ch, nil
}
