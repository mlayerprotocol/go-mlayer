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

var logger = &log.Logger;
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
        c.JSON(http.StatusOK, gin.H{
            "name": "mLayer node",
			"apiVersion": "1.0.0",
			"nodeVersion": "1.0.0",
		})
    })

	router.PUT("/api/authorize", func(c *gin.Context) {
		var payload entities.ClientPayload
        if err := c.BindJSON(&payload); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
		// logger.Infof("PUT %s %v", "/api/authorize", payload.ToJSON())
		// copier.Copy(&payload.ClientPayload, &payload)

		logger.WithFields(logrus.Fields{"payload": string(payload.ToJSON())}).Debug("New auth payload from REST api")
		authEvent, err := client.AuthorizeAgent(payload, p.Ctx)
		if err != nil {
			logger.Error(err)
			c.JSON(http.StatusOK, gin.H{
				"error": err.Error(),
				"apiVersion": "1.0.0",
				"nodeVersion": "1.0.0",
			})
			return
		}
		
        // Send a response back
        c.JSON(http.StatusOK, gin.H{
            //"state": authState,
			"event": authEvent,
		})
    })

	router.POST("/api/topics", func(c *gin.Context) {
		var payload entities.ClientPayload
        if err := c.BindJSON(&payload); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
		logger.Infof("Payload %v", payload.Data)
		topic := entities.Topic{} 
		d, _ := json.Marshal(payload.Data)
		e := json.Unmarshal(d, &topic)
		if e!=nil {
			logger.Errorf("UnmarshalError %v", e)
		}
		payload.Data = topic
		event, err := client.CreateTopic(payload, p.Ctx)

		if err != nil {
			logger.Error(err)
			c.JSON(http.StatusOK, gin.H{
				"error": err.Error(),
				"apiVersion": "1.0.0",
				"nodeVersion": "1.0.0",
			})
			return
		}
        c.JSON(http.StatusOK, gin.H{
            "status": "mLayer node",
			"data": event,
        })
    })

	router.POST("/api/topics/:id/update-avatar", func(c *gin.Context) {
		var payload entities.ClientPayload
        if err := c.BindJSON(&payload); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
		logger.Infof("Payload %v", payload.Data)
		topic := entities.Topic{} 
		d, _ := json.Marshal(payload.Data)
		e := json.Unmarshal(d, &topic)
		if e!=nil {
			logger.Errorf("UnmarshalError %v", e)
		}
		payload.Data = topic
		event, err := client.CreateTopic(payload, p.Ctx)

		if err != nil {
			logger.Error(err)
			c.JSON(http.StatusOK, gin.H{
				"error": err.Error(),
				"apiVersion": "1.0.0",
				"nodeVersion": "1.0.0",
			})
			return
		}
        c.JSON(http.StatusOK, gin.H{
            "status": "mLayer node",
			"data": event,
        })
    })
	return router
}