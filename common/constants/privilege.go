package constants

type AuthorizationPrivilege uint8

const (
	UnauthorizedPriviledge  AuthorizationPrivilege = 0
	BasicPriviledge  AuthorizationPrivilege = 10
	StandardPriviledge AuthorizationPrivilege = 20
	ManagerPriviledge AuthorizationPrivilege = 30
	AdminPriviledge AuthorizationPrivilege = 40
)
type SubscriberRole uint8

var (
	TopicReaderRole SubscriberRole = 0
	TopicWriterRole  SubscriberRole = 10
	TopicManagerRole SubscriberRole = 20
	TopicAdminRole  SubscriberRole = 30
)
