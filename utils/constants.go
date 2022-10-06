package utils

const (
	VALID_HANDSHAKE_SECONDS = 15 // time interval within which to accept a handshake
)

const (
	DefaultRPCPort string = "9521" // time interval within which to accept a handshake
)

const (
	RelayNodeType     uint = 0
	ValidatorNodeType      = 1
)

// Values withing the main context
const (
	ConfigKey         string = "Config"
	OutgoingMessageCh        = "OutgoingMessageChannel"
	IncomingMessageCh        = "IncomingMessageChannel"
	PublishMessageCh         = "PublishMessageChannel"
)
