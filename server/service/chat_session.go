package service

import (
	"server/global"
	"server/model/chat"
)

type chatSessionService struct{}

var ChatSessionService = new(chatSessionService)

func (chatSessionService *chatSessionService) add(chatSession chat.ChatSessions) (err error) {
	err = global.DB.Create(&chatSession).Error
	return err
}

func (chatSessionService *chatSessionService) GetList() (chatSession []chat.ChatSessions, err error) {
	err = global.DB.Find(&chatSession).Error
	return chatSession, err
}
