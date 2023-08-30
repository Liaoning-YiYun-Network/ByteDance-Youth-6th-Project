package type_conv

import "SkyLine/entity"

func ToMessageList(dbMessageList []entity.DBMessage) (messageList []entity.Message) {
	for _, dbMessage := range dbMessageList {
		messageList = append(messageList, entity.Message{
			Id:         dbMessage.MessageID,
			Content:    dbMessage.Content,
			CreateTime: dbMessage.CreateTime,
		})
	}
	return
}
