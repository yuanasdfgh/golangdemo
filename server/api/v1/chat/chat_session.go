package chat

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"server/global"
	"server/response"
	"server/service"
)

/**
* 获取列表
 */
func GetChatSession(c *gin.Context) {

	data, err := service.ChatSessionService.GetList()
	if err != nil {
		global.Log.Error("获取失败!", zap.Error(err))
		response.Failed("失败", c)
		return
	}
	response.Success("成功", data, c)
}
