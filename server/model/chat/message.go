package chat

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Role            string `json:"role" form:"role" gorm:"column:role;comment:角色"`
	Content         string `json:"content" form:"content" gorm:"column:content;comment:内容"`
	ConverssationID uint   `json:"converssation_id" form:"converssation_id" gorm:"column:converssation_id;comment:对话ID"`
}
