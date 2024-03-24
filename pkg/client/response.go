package client

import (
	"github.com/mlayerprotocol/go-mlayer/entities"
	"github.com/mlayerprotocol/go-mlayer/internal/sql/models"
)


type ResponseTimeRange struct {
	Before uint64 `json:"before,omitempty"`
	After uint64 `json:"after,omitempty"`
}
type TopicResponse struct {
	Updates []models.TopicEvent `json:"updates,omitempty"`
	Joins []models.TopicEvent `json:"joins,omitempty"`
	Leaves []models.TopicEvent `json:"leaves,omitempty"`
}

 type SyncResponse struct {
	Time ResponseTimeRange `json:"time,omitempty"`
	Authorization entities.Authorization
	Presence []entities.Event `json:"presence,omitempty"`
	Topics TopicResponse `json:"topics,omitempty"`
 }