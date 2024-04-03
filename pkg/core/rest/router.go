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
		json.Unmarshal(*b, &authEntity)
		auths, err := client.GetAuthorizations(authEntity)

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
		var payload entities.Subscription
		json.Unmarshal(*b, &payload)

		logger.Infof("Payload %v", payload.Topic)

		subs, err := client.GetSubscriptions(payload)

		if err != nil {
			logger.Error(err)
			c.JSON(http.StatusBadRequest, entities.NewClientResponse(entities.ClientResponse{Error: err.Error()}))
			return
		}
		c.JSON(http.StatusOK, entities.NewClientResponse(entities.ClientResponse{Data: subs}))
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
			logger.Errorf("UnmarshalError %v", e)
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

	router.GET("/api/block-stats", func(c *gin.Context) {
		blockStats, err := client.GetBlockStats()
		if err != nil {
			logger.Error(err)
			c.JSON(http.StatusBadRequest, entities.NewClientResponse(entities.ClientResponse{Error: err.Error()}))
			return
		}
		c.JSON(http.StatusOK, entities.NewClientResponse(entities.ClientResponse{Data: blockStats}))
	})

	// router.GET("/api/block-stats", func(c *gin.Context) {
	// 	b, parseError := utils.ParseQueryString(c)
	// 	if parseError != nil {
	// 		logger.Error(parseError)
	// 		c.JSON(http.StatusBadRequest, entities.NewClientResponse(entities.ClientResponse{Error: parseError.Error()}))
	// 		return
	// 	}

	// 	//
	// 	var params BlockParams
	// 	json.Unmarshal(*b, &params)
	// 	fromBlock, fromBlockErr := strconv.Atoi(params.FromBlock)
	// 	toBlock, toBlockErr := strconv.Atoi(params.ToBlock)

	// 	if fromBlockErr != nil || toBlockErr != nil {
	// 		logger.Error(parseError)
	// 		c.JSON(http.StatusBadRequest, entities.NewClientResponse(entities.ClientResponse{Error: utils.IfThenElse(fromBlockErr != nil, fromBlockErr.Error(), toBlockErr.Error())}))
	// 		return
	// 	}
	// 	stats := []BlockStat{}
	// 	for i := fromBlock; i <= toBlock; i++ {

	// 		topicEvents, err := client.GetTopicEvents()
	// 		if err != nil {
	// 			logger.Error(err)
	// 			c.JSON(http.StatusBadRequest, entities.NewClientResponse(entities.ClientResponse{Error: err.Error()}))
	// 			return
	// 		}
	// 		stats = append(stats, BlockStat{
	// 			Events:   i,
	// 			Topics:   i,
	// 			Messages: i,
	// 		})
	// 	}

	// 	logger.Infof("Payload %v", params)

	// 	c.JSON(http.StatusOK, entities.NewClientResponse(entities.ClientResponse{Data: stats}))
	// })
	return router
}

type BlockParams struct {
	FromBlock string `json:"from_block"`
	ToBlock   string `json:"to_block"`
}
type BlockStat struct {
	Events   int `json:"events"`
	Topics   int `json:"topics"`
	Messages int `json:"messages"`
}
