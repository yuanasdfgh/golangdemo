package chat

import "gorm.io/gorm"

type ChatSessions struct {
	gorm.Model
	UserId    int    `json:"user_id" form:"user_id" gorm:"column:user_id;comment:用户id"`
	ChatTitle string `json:"chat_title" form:"chat_title" gorm:"column:chat_title;comment:对话标题"`
}
