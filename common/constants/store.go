package constants

type DataStore string

const (
	// UnprocessedClientPayloadStore DataStore = "unprocessed-client-payload-store"
	// ValidMessageStore             DataStore = "valid-messages"
	// UnsentMessageStore                      = "unsent-messages"
	// SentMessageStore                        = "sent-messages"
	NewTopicSubscriptionStore          DataStore     = "new-topic-subscription"
	TopicSubscriptionStore              DataStore    = "top-subscriptions"
	TopicSubscriptionCountStore             = "topic-subscription-count"
	DeliveryProofStore                      = "delivery-proof-store"
	UnconfirmedDeliveryProofStore           = "unconfirmed-delivery-proof-store"
	DeliveryProofBlockStateStore            = "delivery-proof-block-state-store"
	SubscriptionBlockStateStore             = "subscription-block-state-store"
	DeliveryProofBlockStore                 = "dp-block-store"
	SubscriptionBlockStore                  = "sub-block-store"
	ConnectedSubscribersMap          DataStore       = "connected-subscribers-map"
	EventCountStore                 DataStore		= "event-count-store"
	ClaimedRewardStore               DataStore  		= "claimed-reward-store"
	P2PDataStore                 DataStore		= "p2p-data-store"
)
