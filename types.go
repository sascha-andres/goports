package goports

// OutgoingPort is a port that is used to send data to the outside world.
type OutgoingPort[T any] interface {
	// Send sends data to the outside world.
	Send(data T) error
	// HasChannel returns true if the port has a channel.
	HasChannel() bool
	// Listen will start listening for data on the channel.
	// If the port does not have a channel, this method will
	// return an error.
	Listen() error
}

// IncomingPort is a port that is used to receive data from the outside world.
// Currently, only a method that indicates whether the port has a channel is
// is provided. This is because the channel is not exposed to the outside world.
type IncomingPort[T any] interface {
	// HasChannel returns true if the port has a channel.
	HasChannel() bool
	// Listen will start listening for data on the channel.
	// If the port does not have a channel, this method will
	// return an error.
	Listen() error
}
