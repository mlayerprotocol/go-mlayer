package db

import (
	// "errors"
	"context"
	"errors"

	"github.com/mlayerprotocol/go-mlayer/common/constants"
	"github.com/mlayerprotocol/go-mlayer/entities"
	"github.com/mlayerprotocol/go-mlayer/pkg/core/db"
)

func SaveUnprocessedAuthorizationPayload(key string, payload *entities.AuthorizationPayload, ctx *context.Context) error {
		unprocessedStore, ok := (*ctx).Value(constants.UnprocessedClientPayloadStore).(*db.Datastore)
		if !ok {
			return errors.New("Could not connect to subscription datastore")
		}
		val, _ := unprocessedStore.Get(*ctx,  db.Key(key))
	
		if len(val) > 0 {
			// return errors.New("Duplicate entry found for key")
		}
		error := unprocessedStore.Set(*ctx, db.Key(key), payload.MsgPack(), false)
		if error != nil {
			return error
		}
	
	return nil
}


