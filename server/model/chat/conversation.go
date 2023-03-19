package chat

import "gorm.io/gorm"

type Converssation struct {
	gorm.Model
	Messages []Message `json:"messages" form:"messages"`
}
