package constants

import "html/template"

const (
	VALID_HANDSHAKE_SECONDS = 15 // time interval within which to accept a handshake
)

const (
	DefaultRPCPort            string = "9521"
	DefaultWebSocketAddress   string = "localhost:8088"
	DefaultRestAddress        string = ":9531"
	DefaultDataDir            string = "./data/store"
	DefaultMLBlockchainAPIUrl string = ":9520"
)
const (
	RelayNodeType     uint = 0
	ValidatorNodeType      = 1
)

const MaxBlockSize = 1000

const (
	ErrorUnauthorized = "4001"
	ErrorBadRequest   = "4000"
	ErrorForbidden    = "4003"
)

const (
	UnprocessedClientPayloadStore string = "unprocessed-client-payload-store"
	ValidMessageStore             string = "valid-messages"
	UnsentMessageStore                   = "unsent-messages"
	SentMessageStore                     = "sent-messages"
	NewTopicSubscriptionStore            = "new-topic-subscription"
	TopicSubscriptionStore               = "top-subscriptions"
	TopicSubscriptionCountStore          = "topic-subscription-count"
	DeliveryProofStore                   = "delivery-proof-store"
	UnconfirmedDeliveryProofStore        = "unconfirmed-delivery-proof-store"
	DeliveryProofBlockStateStore         = "delivery-proof-block-state-store"
	SubscriptionBlockStateStore          = "subscription-block-state-store"
	DeliveryProofBlockStore              = "dp-block-store"
	SubscriptionBlockStore               = "sub-block-store"
	ConnectedSubscriber                  = "connected-subscriber"
)

// Channel Ids within main context
const (
	ConfigKey                       string = "Config"
	BroadcastAuthorizationEventChId        = "BroadcastAuthorizationEventChannel"
	BroadcastTopicEventChId                = "BroadcastTopicEventChannel"
	// IncomingAuthorizationEventChId         = "IncomingAuthorizationEventChannel"
	IncomingTopicEventChId = "IncomingTopicEventChannel"

	OutgoingMessageChId     = "OutgoingMessageChannel"
	OutgoingMessageDP2PChId = "OutgoingMessageDP2PChannel"
	IncomingMessageChId     = "IncomingMessageChannel"

	PublishMessageChId              = "PublishMessageChannel"
	SubscribeChId                   = "SubscribeChannel"
	SubscriptionDP2PChId            = "SubscriptionDP2PChannel"
	ClientHandShackChId             = "ClientHandshakeChannel"
	OutgoingDeliveryProof_BlockChId = "OutgoingDeliveryProofBlockChannel"
	OutgoingDeliveryProofChId       = "OutgoinDeliveryProofChannel"
	PubSubBlockChId                 = "PubSubBlockChannel"
	PubsubDeliverProofChId          = "PubsubProofChannel"
	PublishedSubChId                = "PublishedSubChannel"
	SQLDB                           = "sqldb"
)

// State store key
const (
	CurrentDeliveryProofBlockStateKey string = "/df-block/current-state"
	CurrentSubscriptionBlockStateKey  string = "/sub-block/current-state"
)

type Protocol string

const (
	WS   Protocol = "ws"
	MQTT Protocol = "mqtt"
	RPC  Protocol = "rpc"
)

type SubscriptionStatuses int16

const (
	UnsubscribedSubscriptionStatus SubscriptionStatuses = 0
	PendingSubscriptionStatus      SubscriptionStatuses = 1
	SubscribedSubscriptionStatus   SubscriptionStatuses = 2
	// ApprovedSubscriptionStatus      SubscriptionStatuses = "approved"
	BannedSubscriptionStatus SubscriptionStatuses = 3
	// UNBANNED     SubscriptionStatuses = "unbanned"
)

/* KEY MAPS
Always enter map in sorted order

Abi             = abi
Account         = acct
Action          = a
Actions         = as
Address         = addr
Agent           = agt
Amount          = amt
Approval        = ap
ApprovalExpiry  = apExp
Associations    = assoc
Attachments		= atts
Authority       = auth
Avatar          = ava
Body            = b
Block           = blk
BlockId         = blkId
Broadcast       = br
Chain           = c
ChainId         = cId
ChannelExpiry   = chEx
ChannelName     = chN
Channel         = ch
Channels        = chs
CID             = cid
Closed          = cl
Contract        = co
Data            = d
DataHash        = dH
Description     = desc
Duration        = du
EventHash       = eH
File			= f
Grantor         = gr
Handle          = hand
Hash            = h
Header          = head
Identifier		= id
Index           = i
InviteOnly 		= invO
IsValid         = isVal
Length          = len
Message         = m
MessageHash     = mH
MessageSender   = mS
Name            = nm
Node            = n
NodeAddress     = nA
NodeHeight      = nH
NodeSignature   = nS
NodeType        = nT
Origin          = o
Parameters      = pa
ParentTopicHash = pTH
Paylaod         = pld
Platform        = p
Privilege       = privi
Proof           = pr
Proofs          = prs
ProtocolId      = proId
PublicKey		= pubK
ReadOnly		= rO
Receiver        = r
Ref             = ref
Roles 			= rol
Secret          = sec
Sender          = s
SenderAddress   = sA
SenderSignature = sSig
SignatureExpiry = sigExp
Signature       = sig
SignatureData   = sigD
Signer          = sigr
Size            = si
Socket          = sock
Status          = st
Subject         = su
SubjectHash     = suH
Subscriber      = sub
Synced          = sync
Timestamp       = ts
TopicHash       =topH
TopicId         = topId
TopicIds        = topIds
Type            = ty
Url				= url
Validator       = v







*/

var HomeTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<script>  
window.addEventListener("load", function(evt) {
    var output = document.getElementById("output");
    var input = document.getElementById("input");
    var ws;
    var print = function(message) {
        var d = document.createElement("div");
        d.textContent = message;
        output.appendChild(d);
        output.scroll(0, output.scrollHeight);
    };
    document.getElementById("open").onclick = function(evt) {
        if (ws) {
            return false;
        }
        ws = new WebSocket("{{.}}");
        ws.onopen = function(evt) {
            print("OPEN");
        }
        ws.onclose = function(evt) {
            print("CLOSE");
            ws = null;
        }
        ws.onmessage = function(evt) {
            print("RESPONSE: " + evt.data);
        }
        ws.onerror = function(evt) {
            print("ERROR: " + evt.data);
        }
        return false;
    };
    document.getElementById("send").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        print("SEND: " + input.value);
        ws.send(input.value);
        return false;
    };
    document.getElementById("close").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        ws.close();
        return false;
    };
});
</script>
</head>
<body>
<table>
<tr><td valign="top" width="50%">
<p>Click "Open" to create a connection to the server, 
"Send" to send a message to the server and "Close" to close the connection. 
You can change the message and send multiple times.
<p>
<form>
<button id="open">Open</button>
<button id="close">Close</button>
<p><input id="input" type="text" value="Hello world!">
<button id="send">Send</button>
</form>
</td><td valign="top" width="50%">
<div id="output" style="max-height: 70vh;overflow-y: scroll;"></div>
</td></tr></table>
</body>
</html>
`))

/*
{\"type\": \"cosmos-sdk/StdTx\",
\"value\": {\"msg\": [{  \"type\": \"sign/MsgSignData\",  \"value\": {\t\"signer\": \"cosmos14y0pyqjay3p8dsqp2jd5rkft7vf9cdkqnrc43l\",\t\"data\": \"aGVsbG93b3JsZA==\"  }}],\"fee\": {\"amount\": [],\"gas\": \"0\"},\"memo\": \"\"}}
*/
