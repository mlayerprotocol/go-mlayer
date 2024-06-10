package constants

type SubscriptionStatus int16

var (
	UnsubscribedSubscriptionStatus SubscriptionStatus = 0
	InvitedSubscriptionStatus      SubscriptionStatus = 10
	PendingSubscriptionStatus      SubscriptionStatus = 20
	SubscribedSubscriptionStatus   SubscriptionStatus = 30
	// ApprovedSubscriptionStatus      SubscriptionStatus = "approved"
	BannedSubscriptionStatus SubscriptionStatus = 40
	// UNBANNED     SubscriptionStatus = "unbanned"
)