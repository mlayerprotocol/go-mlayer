package rest

import (
	// "errors"
	"context"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mlayerprotocol/go-mlayer/common/constants"
	"github.com/mlayerprotocol/go-mlayer/common/utils"
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
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
func (p *RestService) Initialize() *gin.Engine {
	router := gin.Default()
	router.Use(CORSMiddleware())
	router.GET("/api/ping", func(c *gin.Context) {

		// Send a response back
		c.JSON(http.StatusOK, entities.NewClientResponse(entities.ClientResponse{}))
	})
	router.GET("/api/authorizations", func(c *gin.Context) {

		b, parseError := utils.ParseQueryString(c)
		if parseError != nil {
			logger.Error(parseError)
			c.JSON(http.StatusBadRequest, entities.NewClientResponse(entities.ClientResponse{Error: parseError.Error()}))
			return
		}

		var authEntity entities.Authorization
		var clientPayload entities.ClientPayload

		json.Unmarshal(*b, &authEntity)
		auths, err := client.GetAccountAuthorizations(&authEntity, &clientPayload)


		if err != nil {
			logger.Error(err)
			c.JSON(http.StatusBadRequest, entities.NewClientResponse(entities.ClientResponse{Error: err.Error()}))
			return
		}
		c.JSON(http.StatusOK, entities.NewClientResponse(entities.ClientResponse{Data: auths}))
	})

	router.PUT("/api/authorize", func(c *gin.Context) {
		var payload entities.ClientPayload
		if err := c.BindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, entities.NewClientResponse(entities.ClientResponse{Error: err.Error()}))
			return
		}
		// logger.Infof("PUT %s %v", "/api/authorize", payload.ToJSON())
		// copier.Copy(&payload.ClientPayload, &payload)

		logger.WithFields(logrus.Fields{"payload": string(payload.ToJSON())}).Debug("New auth payload from REST api")
		authEvent, err := client.AuthorizeAgent(payload, p.Ctx)
		if err != nil {
			c.JSON(http.StatusBadRequest, entities.NewClientResponse(entities.ClientResponse{Error: err.Error()}))
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
			c.JSON(http.StatusBadRequest, entities.NewClientResponse(entities.ClientResponse{Error: err.Error()}))
			return
		}
		logger.Infof("Payload %v", payload)
		topic := entities.Topic{}
		d, _ := json.Marshal(payload.Data)
		e := json.Unmarshal(d, &topic)
		if e != nil {
			c.JSON(http.StatusBadRequest, entities.NewClientResponse(entities.ClientResponse{Error: e.Error()}))
		}
		// topic.ID = id
		payload.Data = topic
		event, err := client.CreateEvent(payload, p.Ctx)

		if err != nil {
			logger.Error(err)
			c.JSON(http.StatusBadRequest, entities.NewClientResponse(entities.ClientResponse{Error: err.Error()}))
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
			c.JSON(http.StatusBadRequest, entities.NewClientResponse(entities.ClientResponse{Error: err.Error()}))
			return
		}
		c.JSON(http.StatusOK, entities.NewClientResponse(entities.ClientResponse{Data: topics}))
	})

	router.GET("/api/topics/subscribers/:id", func(c *gin.Context) {
		id := c.Param("id")
		topic, err := client.GetSubscription(id)

		if err != nil {
			logger.Error(err)
			c.JSON(http.StatusBadRequest, entities.NewClientResponse(entities.ClientResponse{Error: err.Error()}))
			return
		}
		c.JSON(http.StatusOK, entities.NewClientResponse(entities.ClientResponse{Data: topic}))
	})

	router.GET("/api/topics/subscribers", func(c *gin.Context) {

		b, parseError := utils.ParseQueryString(c)
		if parseError != nil {
			logger.Error(parseError)
			c.JSON(http.StatusBadRequest, entities.NewClientResponse(entities.ClientResponse{Error: parseError.Error()}))
			return
		}

		//
		var subPayload entities.Subscription
		json.Unmarshal(*b, &subPayload)

		logger.Infof("Payload %v", subPayload.Topic)

		subs, err := client.GetSubscriptions(subPayload)

		if err != nil {
			logger.Error(err)
			c.JSON(http.StatusBadRequest, entities.NewClientResponse(entities.ClientResponse{Error: err.Error()}))
			return
		}
		var payload entities.ClientPayload
		if err := c.BindJSON(&payload); err != nil {
			logger.Error(err)
			c.JSON(http.StatusBadRequest, entities.NewClientResponse(entities.ClientResponse{Error: err.Error()}))
			return
		}
		c.JSON(http.StatusOK, entities.NewClientResponse(entities.ClientResponse{Data: subs}))
	})

	router.GET("/api/topics/:id/messages", func(c *gin.Context) {
		id := c.Param("id")
		messages, err := client.GetMessages(id)

		if err != nil {
			logger.Error(err)
			c.JSON(http.StatusBadRequest, entities.NewClientResponse(entities.ClientResponse{Error: err.Error()}))
			return
		}
		c.JSON(http.StatusOK, entities.NewClientResponse(entities.ClientResponse{Data: messages}))
	})

	router.GET("/api/topics/:id", func(c *gin.Context) {
		id := c.Param("id")
		topic, err := client.GetTopicById(id)

		if err != nil {
			logger.Error(err)
			c.JSON(http.StatusBadRequest, entities.NewClientResponse(entities.ClientResponse{Error: err.Error()}))
			return
		}
		c.JSON(http.StatusOK, entities.NewClientResponse(entities.ClientResponse{Data: topic}))
	})

	router.POST("/api/topics/subscribe", func(c *gin.Context) {
		// id := c.Param("id")

		var payload entities.ClientPayload
		if err := c.BindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, entities.NewClientResponse(entities.ClientResponse{Error: err.Error()}))
			return
		}
		logger.Infof("Payload %v", payload.Data)
		subscription := entities.Subscription{}
		d, _ := json.Marshal(payload.Data)
		e := json.Unmarshal(d, &subscription)
		if e != nil {
			logger.Errorf("UnmarshalError %v", e)
			c.JSON(http.StatusBadRequest, entities.NewClientResponse(entities.ClientResponse{Error: e.Error()}))
			return
		}
		// subscription.ID = id
		payload.Data = subscription
		event, err := client.CreateEvent(payload, p.Ctx)

		if err != nil {
			c.JSON(http.StatusBadRequest, entities.NewClientResponse(entities.ClientResponse{Error: err.Error()}))
			return
		}

		c.JSON(http.StatusOK, entities.NewClientResponse(entities.ClientResponse{Data: event}))

	})

	router.PATCH("/api/topics/subscribers/approve", func(c *gin.Context) {
		id := c.Param("id")

		var payload entities.ClientPayload
		if err := c.BindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, entities.NewClientResponse(entities.ClientResponse{Error: err.Error()}))
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
		event, err := client.CreateEvent(payload, p.Ctx)

		if err != nil {
			logger.Error(err)
			c.JSON(http.StatusBadRequest, entities.NewClientResponse(entities.ClientResponse{Error: err.Error()}))
			return
		}

		c.JSON(http.StatusOK, entities.NewClientResponse(entities.ClientResponse{Data: event}))

	})

	router.PATCH("/api/topics/unsubscribe", func(c *gin.Context) {
		id := c.Param("id")

		var payload entities.ClientPayload
		if err := c.BindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, entities.NewClientResponse(entities.ClientResponse{Error: err.Error()}))
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
		event, err := client.CreateEvent(payload, p.Ctx)

		if err != nil {
			logger.Error(err)
			c.JSON(http.StatusBadRequest, entities.NewClientResponse(entities.ClientResponse{Error: err.Error()}))
			return
		}

		c.JSON(http.StatusOK, entities.NewClientResponse(entities.ClientResponse{Data: event}))

	})

	router.PATCH("/api/topics/ban", func(c *gin.Context) {
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
		event, err := client.CreateEvent(payload, p.Ctx)

		if err != nil {
			logger.Error(err)
			c.JSON(http.StatusBadRequest, entities.NewClientResponse(entities.ClientResponse{Error: err.Error()}))
			return
		}

		c.JSON(http.StatusOK, entities.NewClientResponse(entities.ClientResponse{Data: event}))

	})

	router.PATCH("/api/topics/:id", func(c *gin.Context) {
		id := c.Param("id")

		var payload entities.ClientPayload
		if err := c.BindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, entities.NewClientResponse(entities.ClientResponse{Error: err.Error()}))
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
		event, err := client.CreateEvent(payload, p.Ctx)

		if err != nil {
			logger.Error(err)
			c.JSON(http.StatusBadRequest, entities.NewClientResponse(entities.ClientResponse{Error: err.Error()}))
			return
		}

		c.JSON(http.StatusOK, entities.NewClientResponse(entities.ClientResponse{Data: map[string]any{
			"event": event,
		}}))
	})

	router.POST("/api/topics/messages", func(c *gin.Context) {
		var payload entities.ClientPayload
		if err := c.BindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, entities.NewClientResponse(entities.ClientResponse{Error: err.Error()}))
			return
		}
		logger.Infof("Payload %v", payload.Data)
		message := entities.Message{}
		d, _ := json.Marshal(payload.Data)
		e := json.Unmarshal(d, &message)
		if e != nil {
			logger.Errorf("Unmarshal Error %v", e)
			c.JSON(http.StatusBadRequest, entities.NewClientResponse(entities.ClientResponse{Error: e.Error()}))
			return
		}
		// subscription.ID = id
		payload.Data = message
		event, err := client.CreateEvent(payload, p.Ctx)

		if err != nil {
			c.JSON(http.StatusBadRequest, entities.NewClientResponse(entities.ClientResponse{Error: err.Error()}))
			return
		}

		c.JSON(http.StatusOK, entities.NewClientResponse(entities.ClientResponse{Data: event}))

	})

	router.POST("/api/subscription/account", func(c *gin.Context) {

		b, parseError := utils.ParseQueryString(c)
		if parseError != nil {
			logger.Error(parseError)
			c.JSON(http.StatusBadRequest, entities.NewClientResponse(entities.ClientResponse{Error: parseError.Error()}))
			return
		}

		//
		var payload entities.ClientPayload
		json.Unmarshal(*b, &payload)

		logger.Infof("Payload %v", payload.Account)
		subscriptions, err := client.GetAccountSubscriptions(payload)

		if err != nil {
			logger.Error(err)
			c.JSON(http.StatusBadRequest, entities.NewClientResponse(entities.ClientResponse{Error: err.Error()}))
			return
		}
		c.JSON(http.StatusOK, entities.NewClientResponse(entities.ClientResponse{Data: subscriptions}))
	})

	router.GET("/api/sync", func(c *gin.Context) {
		b, parseError := utils.ParseQueryString(c)
		if parseError != nil {
			logger.Error(parseError)
			c.JSON(http.StatusBadRequest, entities.NewClientResponse(entities.ClientResponse{Error: parseError.Error()}))
			return
		}
		var authEntity entities.Authorization
		var payload entities.ClientPayload
		json.Unmarshal(*b, &authEntity)
		json.Unmarshal(*b, &payload)

		syncResponse := entities.SyncResponse{}
		client.SyncAgent(&entities.SyncRequest{}, &entities.ClientPayload{})

		// if err != nil {
		// 	logger.Error(err)
		// 	c.JSON(http.StatusBadRequest, entities.NewClientResponse(entities.ClientResponse{Error: err.Error()}))
		// 	return
		// }
		c.JSON(http.StatusOK, entities.NewClientResponse(entities.ClientResponse{Data: syncResponse}))
	})

	return router
}
