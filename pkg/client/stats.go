package client

import (
	// "errors"

	"math/big"

	"github.com/mlayerprotocol/go-mlayer/entities"

	"github.com/mlayerprotocol/go-mlayer/internal/chain"
	"github.com/mlayerprotocol/go-mlayer/internal/sql/models"
	query "github.com/mlayerprotocol/go-mlayer/internal/sql/query"
	"gorm.io/gorm"
)

func GetBlockStats() (*[]models.BlockStat, error) {
	var blockStat []models.BlockStat

	err := query.GetManyTx(models.BlockStat{}).Order("block_number desc").Find(&blockStat).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &blockStat, nil
}

func GetMainStats() (*entities.MainStat, error) {
	// var mainStat []entities.MainStat
	var accountCount int64
	var topicBalanceCount int64
	var messageCount int64
	

	err := query.GetTx().Model(&models.AuthorizationState{}).Group("account").Count(&accountCount).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	err = query.GetTx().Model(&models.TopicState{}).Select("sum(balance)").Row().Scan(&topicBalanceCount)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	// err = query.GetTx().Model(&models.MessageState{}).Count(&messages).Error
	// if err != nil {
	// 	if err == gorm.ErrRecordNotFound {
	// 		return nil, nil
	// 	}
	// 	return nil, err
	// }

	err = query.GetTx().Model(&models.MessageState{}).Count(&messageCount).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &entities.MainStat{
		Accounts:     accountCount,
		TopicBalance: topicBalanceCount,
		Messages:     messageCount,
		MessageCount: big.NewInt(0).Mul(chain.MLChainApi.GetCurrentMessageCost(), big.NewInt(messageCount)),
	}, nil
}
