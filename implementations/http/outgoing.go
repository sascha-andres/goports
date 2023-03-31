package http

import (
	"errors"
)

// OutgoingPort is an outgoing port that sends data to the outside world
// using HTTP. It implements OutgoingPort[T any].
type OutgoingPort[T any] struct {
	// c is the receiving channel
	c chan T
}

// Send sends data to the outside world using HTTP.
func (p *OutgoingPort[T]) Send(data T) error {
	return errors.New("not implemented")
}

// HasChannel returns true if the port has a channel.
func (p *OutgoingPort[T]) HasChannel() bool {
	return p.c != nil
}

// Listen will start listening for data on the channel.
func (p *OutgoingPort[T]) Listen() error {
	if !p.HasChannel() {
		return errors.New("port does not have a channel")
	}
	for data := range p.c {
		_ = data
		// TODO(sa): send data to the outside world
	}
	return nil
}

// NewOutgoingPort creates a new outgoing port.
func NewOutgoingPort[T any](c chan T) (*OutgoingPort[T], error) {
	// TODO(sa): read configuration from environment variables
	return &OutgoingPort[T]{
		c: c,
	}, nil
}
