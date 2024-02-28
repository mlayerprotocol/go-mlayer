package rest

import (
	// "errors"
	"context"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mlayerprotocol/go-mlayer/common/constants"
	"github.com/mlayerprotocol/go-mlayer/configs"
	"github.com/mlayerprotocol/go-mlayer/entities"
	"github.com/mlayerprotocol/go-mlayer/pkg/client"
	"github.com/mlayerprotocol/go-mlayer/pkg/log"
	"github.com/sirupsen/logrus"
)

var logger = &log.Logger

type Flag string

type RestService struct {
	Ctx                    *context.Context
	Cfg                    *configs.MainConfiguration
	ClientHandshakeChannel *chan *entities.ClientHandshake
}

type RestResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func NewRestService(mainCtx *context.Context) *RestService {
	cfg, _ := (*mainCtx).Value(constants.ConfigKey).(*configs.MainConfiguration)
	clientVerificationc, _ := (*mainCtx).Value(constants.ClientHandShackChId).(*chan *entities.ClientHandshake)
	return &RestService{
		Ctx:                    mainCtx,
		Cfg:                    cfg,
		ClientHandshakeChannel: clientVerificationc,
	}
}

func (p *RestService) Initialize() *gin.Engine {
	router := gin.Default()
	router.GET("/api/ping", func(c *gin.Context) {

		// Send a response back
		c.JSON(http.StatusOK, entities.NewClientResponse(entities.ClientResponse{}))
	})

	router.PUT("/api/authorize", func(c *gin.Context) {
		var payload entities.ClientPayload
		if err := c.BindJSON(&payload); err != nil {
			c.JSON(http.StatusOK, entities.NewClientResponse(entities.ClientResponse{Error: err}))
			return
		}
		// logger.Infof("PUT %s %v", "/api/authorize", payload.ToJSON())
		// copier.Copy(&payload.ClientPayload, &payload)

		logger.WithFields(logrus.Fields{"payload": string(payload.ToJSON())}).Debug("New auth payload from REST api")
		authEvent, err := client.AuthorizeAgent(payload, p.Ctx)
		if err != nil {
			
			c.JSON(http.StatusOK, entities.NewClientResponse(entities.ClientResponse{Error: err}))
			return
		}

		// Send a response back
		c.JSON(http.StatusOK, entities.NewClientResponse(entities.ClientResponse{Data: map[string]any{
			"event": authEvent,
		}}))
	})

	router.POST("/api/topics", func(c *gin.Context) {
		var payload entities.ClientPayload
		if err := c.BindJSON(&payload); err != nil {
			c.JSON(http.StatusOK, entities.NewClientResponse(entities.ClientResponse{Error: err}))
			return
		}
		logger.Infof("Payload %v", payload.Data)
		topic := entities.Topic{}
		d, _ := json.Marshal(payload.Data)
		e := json.Unmarshal(d, &topic)
		if e != nil {
			c.JSON(http.StatusOK, entities.NewClientResponse(entities.ClientResponse{Error: e}))
		}
		payload.Data = topic
		event, err := client.CreateTopic(payload, p.Ctx)

		if err != nil {
			logger.Error(err)
			c.JSON(http.StatusOK, entities.NewClientResponse(entities.ClientResponse{Error: err}))
			return
		}

		c.JSON(http.StatusOK, entities.NewClientResponse(entities.ClientResponse{Data: map[string]any{
			"event": event,
		}}))
	})

	router.GET("/api/topics", func(c *gin.Context) {
		topics, err := client.GetTopics()

		if err != nil {
			logger.Error(err)
			c.JSON(http.StatusOK, entities.NewClientResponse(entities.ClientResponse{Error: err}))
			return
		}
		c.JSON(http.StatusOK, entities.NewClientResponse(entities.ClientResponse{Data: map[string]any{
			"data": topics,
		}}))
	})

	router.GET("/api/topics/:id", func(c *gin.Context) {
		id := c.Param("id")
		topic, err := client.GetTopic(id)

		if err != nil {
			logger.Error(err)
			c.JSON(http.StatusOK, entities.NewClientResponse(entities.ClientResponse{Error: err}))
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "mLayer node",
			"data":   topic,
		})
	})

	router.PATCH("/api/topics/:id", func(c *gin.Context) {
		id := c.Param("id")

		var payload entities.ClientPayload
		if err := c.BindJSON(&payload); err != nil {
			c.JSON(http.StatusOK, entities.NewClientResponse(entities.ClientResponse{Error: err}))
			return
		}
		logger.Infof("Payload %v", payload.Data)
		topic := entities.Topic{}
		d, _ := json.Marshal(payload.Data)
		e := json.Unmarshal(d, &topic)
		if e != nil {
			logger.Errorf("UnmarshalError %v", e)
		}
		topic.Hash = id
		payload.Data = topic
		event, err := client.CreateUpdateTopicEvent(payload, p.Ctx)

		if err != nil {
			logger.Error(err)
			c.JSON(http.StatusOK, entities.NewClientResponse(entities.ClientResponse{Error: err}))
			return
		}

		c.JSON(http.StatusOK, entities.NewClientResponse(entities.ClientResponse{Data: map[string]any{
			"event": event,
		}}))
	})

	router.POST("/api/topic/subscribe", func(c *gin.Context) {
		// id := c.Param("id")

		var payload entities.ClientPayload
		if err := c.BindJSON(&payload); err != nil {
			c.JSON(http.StatusOK, entities.NewClientResponse(entities.ClientResponse{Error: err}))
			return
		}
		logger.Infof("Payload %v", payload.Data)
		subscription := entities.Subscription{}
		d, _ := json.Marshal(payload.Data)
		e := json.Unmarshal(d, &subscription)
		if e != nil {
			logger.Errorf("UnmarshalError %v", e)
		}
		// subscription.Hash = id
		payload.Data = subscription
		event, err := client.CreateSubscribeTopicEvent(payload, p.Ctx)

		if err != nil {
			c.JSON(http.StatusOK, entities.NewClientResponse(entities.ClientResponse{Error: err}))
			return
		}

		c.JSON(http.StatusOK, entities.NewClientResponse(entities.ClientResponse{Data: map[string]any{
			"event": event,
		}}))

	})

	router.PATCH("/api/topic/unsubscribe", func(c *gin.Context) {
		// id := c.Param("id")

		var payload entities.ClientPayload
		if err := c.BindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		logger.Infof("Payload %v", payload.Data)
		subscription := entities.Subscription{}
		d, _ := json.Marshal(payload.Data)
		e := json.Unmarshal(d, &subscription)
		if e != nil {
			logger.Errorf("UnmarshalError %v", e)
		}
		// subscription.ID = id
		payload.Data = subscription
		event, err := client.CreateUnSubscribeTopicEvent(payload, p.Ctx)

		if err != nil {
			logger.Error(err)
			c.JSON(http.StatusOK, gin.H{
				"error":       err.Error(),
				"apiVersion":  "1.0.0",
				"nodeVersion": "1.0.0",
			})
			return
		}

		c.JSON(http.StatusOK, entities.NewClientResponse(entities.ClientResponse{Data: map[string]any{
			"event": event,
		}}))

	})

	router.PATCH("/api/subscription/:id/approve", func(c *gin.Context) {
		id := c.Param("id")

		var payload entities.ClientPayload
		if err := c.BindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		logger.Infof("Payload %v", payload.Data)
		subscription := entities.Subscription{}
		d, _ := json.Marshal(payload.Data)
		e := json.Unmarshal(d, &subscription)
		if e != nil {
			logger.Errorf("UnmarshalError %v", e)
		}
		subscription.ID = id
		payload.Data = subscription
		event, err := client.CreateSubscriptionApprovalEvent(payload, p.Ctx)

		if err != nil {
			logger.Error(err)
			c.JSON(http.StatusOK, gin.H{
				"error":       err.Error(),
				"apiVersion":  "1.0.0",
				"nodeVersion": "1.0.0",
			})
			return
		}

		c.JSON(http.StatusOK, entities.NewClientResponse(entities.ClientResponse{Data: map[string]any{
			"event": event,
		}}))

	})

	router.GET("/api/subscription/:id", func(c *gin.Context) {
		id := c.Param("id")
		topic, err := client.GetSubscription(id)

		if err != nil {
			logger.Error(err)
			c.JSON(http.StatusOK, gin.H{
				"error":       err.Error(),
				"apiVersion":  "1.0.0",
				"nodeVersion": "1.0.0",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "mLayer node",
			"data":   topic,
		})
	})

	router.GET("/api/subscriptions", func(c *gin.Context) {
		subs, err := client.GetSubscriptions()

		if err != nil {
			logger.Error(err)
			c.JSON(http.StatusOK, gin.H{
				"error":       err.Error(),
				"apiVersion":  "1.0.0",
				"nodeVersion": "1.0.0",
			})
			return
		}
		c.JSON(http.StatusOK, entities.NewClientResponse(entities.ClientResponse{Data: map[string]any{
			"subscriptions": subs,
		}}))
	})

	return router
}
