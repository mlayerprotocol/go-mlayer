package client

import (
	// "errors"

	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/mlayerprotocol/go-mlayer/common/apperror"
	"github.com/mlayerprotocol/go-mlayer/common/constants"
	"github.com/mlayerprotocol/go-mlayer/configs"
	"github.com/mlayerprotocol/go-mlayer/entities"
	"github.com/mlayerprotocol/go-mlayer/internal/crypto"
	dsquery "github.com/mlayerprotocol/go-mlayer/internal/ds/query"
	"github.com/mlayerprotocol/go-mlayer/internal/service"
	"github.com/mlayerprotocol/go-mlayer/internal/sql/models"
	query "github.com/mlayerprotocol/go-mlayer/internal/sql/query"
	"github.com/mlayerprotocol/go-mlayer/pkg/core/p2p"
)

/// When a client connects, they are interested in 2 things
// 1. To create messages
// 2. To subscribe
// Therefore
// In case of one
// - check to see if node has interacted with appect
// - if not, sync from bootstrap
// - if client sends a message to a topic not found locally, check bootstrap for topic
// - load all nodes interested in the topic from bootstrap
// - after processing message locally, broadcast to all nodes interested

// In case of 2
// publish your interest in the topic to all locally connected nodes. They will rebroadcast
// to their own locally connected nodes
// that node will receive any message sent to the topic


func GetSubscribedApplications(item models.ApplicationState) (state *[]models.ApplicationState, err error) {

	// _, cacheError := dsquery.GetCacheKey(dsquery.AccountConnecteaKey, string(item.Account))
	// if dsquery.IsErrorNotFound(cacheError) {
	// 	// never interacted with
	// 	defer func () {
	// 		if err == nil {
	// 			dsquery.SetCacheKey(dsquery.AccountConnecteaKey,string(item.Account), []byte{} )
	// 		}
	// 	}()

	// 	// go the p2p route
	// 	payload := p2p.NewP2pPayload(config, p2p.P2pActionGetAccountSubscriptions, []byte(item.Account))
		
	// 	return
	// }

	 ApplicationStates := []models.ApplicationState{}
	_apps, err := dsquery.GetAccountApplications(item.Account, *dsquery.DefaultQueryLimit)
	if err != nil {
		return &ApplicationStates, err
	}
	logger.Infof("AccountApplications: %v", _apps)
	for _, _sub := range _apps {
		ApplicationStates = append(ApplicationStates, models.ApplicationState{Application: *_sub})
	}
	subscriptionStates, err := dsquery.GetSubscriptions(entities.Subscription{Subscriber: entities.AddressString(item.Account)}, nil, nil)

	if err != nil {

		return &ApplicationStates, err
	}
	var appIds = []string{}
	


	for _, sub := range subscriptionStates {
		appIds = append(appIds, sub.Application)
	}
	var subApplicationStates []models.ApplicationState
	if len(appIds) > 0 {
		subApplicationErr := query.GetWithIN(models.ApplicationState{}, &subApplicationStates, appIds)
		if subApplicationErr != nil {
			return &ApplicationStates, err
		}
	}

	ApplicationStates = append(ApplicationStates, subApplicationStates...)

	return &ApplicationStates, nil
}

// func GetApplicationEvents() (*[]models.ApplicationEvent, error) {
// 	var ApplicationEvents []models.ApplicationEvent

// 	err := query.GetMany(models.ApplicationEvent{
// 		Event: entities.Event{
// 			BlockNumber: 1,
// 		},
// 	}, &ApplicationEvents, nil)
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return nil, nil
// 		}
// 		return nil, err
// 	}
// 	return &ApplicationEvents, nil
// }

// func ListenForNewApplicationEventFromPubSub (mainCtx *context.Context) {
// 	ctx, cancel := context.WithCancel(*mainCtx)
// 	defer cancel()

//		incomingApplicationC, ok := (*mainCtx).Value(constants.IncomingApplicationEventChId).(*chan *entities.Event)
//		if !ok {
//			logger.Errorf("incomingApplicationC closed")
//			return
//		}
//		for {
//			event, ok :=  <-*incomingApplicationC
//			if !ok {
//				logger.Fatal("incomingApplicationC closed for read")
//				return
//			}
//			go service.HandleNewPubSubApplicationEvent(event, ctx)
//		}
//	}
func ValidateApplicationPayload(payload entities.ClientPayload, authState *models.AuthorizationState, ctx *context.Context) (assocPrevEvent *entities.EventPath, assocAuthEvent *entities.EventPath, err error) {

	payloadData := entities.Application{}
	d, _ := json.Marshal(payload.Data)
	e := json.Unmarshal(d, &payloadData)
	if e != nil {
		logger.Errorf("UnmarshalError %v", e)
		return nil, nil, apperror.BadRequest(e.Error())
	}

	payload.Data = payloadData


	if uint64(payloadData.Timestamp) == 0 || uint64(payloadData.Timestamp) > uint64(time.Now().UnixMilli())+15000 || uint64(payloadData.Timestamp) < uint64(time.Now().UnixMilli())-15000 {
		return nil, nil, apperror.BadRequest("Invalid event timestamp")
	}
	cfg, _ := (*ctx).Value(constants.ConfigKey).(*configs.MainConfiguration)

	currentState, err2 := service.ValidateApplicationData(&payload, cfg.ChainId)
	logger.Infof("IVLAIDERR %v, %v", err2, currentState)
	if err2 != nil {
		return nil, nil, err2
	}
	if payload.EventType == constants.CreateApplicationEvent {
		// dont worry validating the AuthHash for Authorization requests
		// if entities.AddressFromString(payloadData.Owner.ToString()).Addr == "" {
		// 	return nil, nil, apperror.BadRequest("You must specify the owner of the app")
		// }
		if payloadData.ID != "" {
			return nil, nil, apperror.BadRequest("You cannot set an id when creating a app")
		}
		// var found []models.ApplicationState
		// query.GetMany(&models.ApplicationState{Application: entities.Application{Ref: payloadData.Ref}}, &found, nil)
		
		refExists, err := dsquery.RefExists(entities.ApplicationModel, payloadData.Ref, "")
		
		if err != nil {
			return nil, nil, err
		}
		// if len(found) > 0 {
		// 	return nil, nil, apperror.BadRequest(fmt.Sprintf("Application with reference %s already exists", payloadData.Ref))
		// }
		if refExists {
			return nil, nil, apperror.BadRequest(fmt.Sprintf("Application with reference %s already exists", payloadData.Ref))
		}
		keySecP := "/ml/snetref/" + hex.EncodeToString(crypto.Keccak256Hash([]byte(payloadData.Ref)))
		v, err := p2p.GetDhtValue(keySecP)
		if err != nil {
			logger.Debugf("DHTERROR: %v", err)
			// return nil, nil, err
		}
		if  len(v) > 0 {
			return nil, nil, apperror.BadRequest(fmt.Sprintf("Application with reference %s already exists", payloadData.Ref))
		}
		// logger.Debug("FOUNDDDDD", found, payloadData.Ref)

	}
	if payload.EventType == constants.UpdateApplicationEvent {
		if payloadData.ID == "" {
			
			return nil, nil, apperror.BadRequest("Application ID must be provided")
		}
	}
	
	
	// generate associations
	if currentState != nil {
		//logger.Debugf("APPINFO %v, %s, %s", strings.EqualFold(currentState.Account.ToString(), payloadData.Account.ToString()), currentState.Account.ToString(), payloadData.Account.ToString())
		if !strings.EqualFold(string(currentState.Account), string(payloadData.Account)) {
			return nil, nil, apperror.BadRequest("app account do not match")
		}
		assocPrevEvent = &currentState.Event
	}
	if authState != nil {
		assocAuthEvent = &authState.Event
	}
	
	return assocPrevEvent, assocAuthEvent, nil
}
