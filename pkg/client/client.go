package client

import (
	"encoding/hex"

	"github.com/mlayerprotocol/go-mlayer/common/apperror"
	"github.com/mlayerprotocol/go-mlayer/entities"
	"github.com/mlayerprotocol/go-mlayer/internal/crypto"
	"github.com/mlayerprotocol/go-mlayer/internal/sql/models"
	query "github.com/mlayerprotocol/go-mlayer/internal/sql/query"
	"gorm.io/gorm"
)

func ValidateClientPayload(
	payload *entities.ClientPayload,
) (*models.AuthorizationState, error) {
	logger.Info("PAYLOAD", string(payload.ToJSON()))
	d, err := payload.EncodeBytes()
	logger.Info("ENCODED", hex.EncodeToString(d))
	agent, err := crypto.GetSignerECC(&d, &payload.Signature)
	logger.Infof("device %s", agent)
	if err != nil {
		return nil, err
	}
	// check if device is authorized
	if agent != "" {
		authData := models.AuthorizationState{}
		err := query.GetOne(models.AuthorizationState{
			Authorization: entities.Authorization{Account: payload.Account,
				Agent: agent},
		}, &authData)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, nil
			}
			return nil, err
		} else {
			return &authData, nil
		}
	}
	return nil, apperror.BadRequest("Unable to resolve agent")
}
