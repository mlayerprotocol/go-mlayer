package entities

import "math/big"

type MainStat struct {
	Accounts     int64    `json:"accounts"`
	TopicBalance int64    `json:"topic_balance"`
	Messages     int64    `json:"messages"`
	MessageCount *big.Int `json:"message_cost"`
}
