package utils

//* Matrix Network Events *//
// m.room.message: The most common event for sending messages in a conversation. It can include text messages, images, videos, etc.
// m.room.encrypted: Used for sending encrypted messages.
// m.room.redaction: Used for redacting (or "deleting") previously sent entities.
// m.call.invite, m.call.candidates, m.call.answer, m.call.hangup: Events used for VoIP calls.
// m.reaction: Used for adding reactions to messages.
// m.room.member:

// m.room.create: The first event in a room, which creates the room.
// m.room.name: Sets the room's name.
// m.room.topic: Sets the room's topic or description.
// m.room.avatar: Sets the room's avatar (image).
// m.room.power_levels: Defines the power levels (permissions) for users in the room.
// m.room.history_visibility: Controls the visibility of the room's history.
// m.room.guest_access: Controls whether guests can join the room.
// m.room.join_rules: Defines how users can join the room (e.g., public, invite-only).
// m.room.third_party_invite: Represents an invitation to the room via a third-party service.
// m.room.pinned_events: Indicates messages that are pinned in the room.


type EventType uint16

// Administrative Topic Actions
const (
    DeleteTopic            EventType = 1000 
	CreateTopic         EventType = 1001     // m.room.create
	PrivacySet          EventType = 1002
    BanMember           EventType = 1003
    UnbanMember         EventType = 1004
    ContractSet         EventType = 1005
    UpdateName          EventType = 1006     //  m.room.name
    UpdateDescription   EventType = 1007     //  m.room.topic
    UpdateAvatar        EventType = 1008     //  m.room.avatar
    PinMessage          EventType = 1008     //  m.room.avatar
)


// Member Topic Actions
const (
    Leave     EventType = 1100
	Join      EventType = 1101
    Requested      EventType = 1102
    Approved        EventType = 1103
    Upgraded        EventType = 1104
    Invited        EventType = 1105
)


// Message Actions
const (
    DeleteMessage              EventType = 1200      //m.room.encrypted
	CreateMessage              EventType = 1201      // m.room.message
	CreateReaction          EventType = 1202 // m.reaction
    IsTyping                EventType = 1203
)

