package http

import "errors"

// IncomingPort is an incoming port that receives data from the outside world
// using HTTP. It implements IncomingPort[T any].
type IncomingPort[T any] struct {
	c chan T
}

// HasChannel returns true if the port has a channel.
func (p *IncomingPort[T]) HasChannel() bool {
	return p.c != nil
}

// Listen will start listening for data on the channel.
func (p *IncomingPort[T]) Listen() error {
	if !p.HasChannel() {
		return errors.New("port does not have a channel")
	}
	// TODO(sa): to receive data from the outside world
	//   create an http server and listen for requests
	return nil
}

// NewIncomingPort creates a new incoming port.
func NewIncomingPort[T any](c chan T) (*OutgoingPort[T], error) {
	if c == nil {
		return nil, errors.New("channel is nil")
	}
	// TODO(sa): read configuration from environment variables
	return &OutgoingPort[T]{
		c: c,
	}, nil
}
