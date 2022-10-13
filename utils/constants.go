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

const (
	ValidMessageStore            string = "valid-messages"
	UnsentMessageStore                  = "unsent-messages"
	SentMessageStore                    = "sent-messages"
	ChannelSubscriberStore              = "channels-subscribers"
	ChannelSubscribersCountStore        = "channels-subscribers-count"
)

// Values withing the main context
const (
	ConfigKey          string = "Config"
	OutgoingMessageCh         = "OutgoingMessageChannel"
	IncomingMessageCh         = "IncomingMessageChannel"
	PublishMessageCh          = "PublishMessageChannel"
	SubscribeCh               = "SubscribeChannel"
	SubscriptionDP2PCh        = "SubscriptionDP2PChannel"
)

type SubAction string

const (
	Join  SubAction = "join"
	Leave SubAction = "leave"
)
