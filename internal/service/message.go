package service

import (
	"github.com/mlayerprotocol/go-mlayer/common/apperror"
	"github.com/mlayerprotocol/go-mlayer/entities"
)


func IsValidTopic(ch entities.Message, signature string, channelOwner string) bool {
	
	
	return true
}



// func IsValidMessage(msg entities.Message, signature string) bool {
// 	chatMessage := msg.ToJSON()
// 	msgByte := []byte(msg.ToString())
// 	signer, _ := crypto.GetSignerECC(&msgByte, &signature)
// 	channel := strings.Split(string(msg.Receiver), ":")
// 	chaByte := []byte(strings.ToLower(channel[0]))
// 	channelOwner, _ := crypto.GetSignerECC(&chaByte, &(channel[1]))
// 	if !strings.EqualFold(channelOwner, signer) {
// 		return false
// 	}
// 	if !IsValidTopic(msg, channel[1], channelOwner) {
// 		return false
// 	}
// 	if math.Abs(float64(int(msg.Timestamp)-int(time.Now().Unix()))) > constants.VALID_HANDSHAKE_SECONDS {
// 		logger.WithFields(logrus.Fields{"data": chatMessage}).Warnf("Message Expired: %s", chatMessage)
// 		return false
// 	}
// 	message := []byte(msg.ToString())
// 	isValid := crypto.VerifySignatureECC(signer, &message, signature)
// 	if !isValid {
// 		logger.WithFields(logrus.Fields{"message": string(message), "signature": signature}).Warnf("Invalid signer %s", signer)
// 		return false
// 	} else {

// 	}
// 	return true
// }

/*
Validate an agent authorization
*/
func ValidateMessageData(message *entities.Message, payload *entities.ClientPayload) ( err error) {
	// check fields of subscription
	

	if len(message.Data) > 1024 {
		return  apperror.BadRequest("Message data too long")
	}
	
	// err = query.GetOne(models.MessageState{
	// 	Message: entities.Message{Hash: message.Hash},
	// }, &currentState)
	// if err != nil {
	// 	if err != gorm.ErrRecordNotFound {
	// 		return nil, err
	// 	}
	// }


	return  nil
}