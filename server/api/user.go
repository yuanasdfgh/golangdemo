package api

//import (
//	"github.com/gin-gonic/gin"
//	"server/common"
//	"server/model"
//	"server/response"
//	"server/service"
//)
//
//type WebUser struct {
//	service.AdminUserService
//}
//
//type AppUser struct {
//	service.UserService
//}
//
//func GetWebUser() *WebUser {
//	return &WebUser{}
//}
//
//func GetAppUser() *AppUser {
//	return &AppUser{}
//}
//
//// 获取验证码
//func (u *WebUser) GetCaptcha(c *gin.Context) {
//	id, base64s, _ := common.GenerateCaptcha()
//	data := map[string]interface{}{"captchaId": id, "captchaImg": base64s}
//	response.Success("操作成功", data, c)
//}
//
//// 用户登录
//func (u *WebUser) UserLogin(c *gin.Context) {
//	var param model.LoginParam
//	if err := c.ShouldBind(&param); err != nil {
//		response.Failed("error", c)
//		return
//	}
//
//	// 检查验证码
//	if !common.VerifyCaptcha(param.CaptchaId, param.CaptchaValue) {
//		response.Failed("验证码错误", c)
//		return
//	}
//	// 生成token
//	uid := u.Login(param)
//	if uid > 0 {
//		token, _ := common.GenerateToke(param.Username)
//		userInfo := model.AdminUserInfo{
//			Sid:   uid,
//			Token: token,
//		}
//		response.Success("登录成功", userInfo, c)
//		return
//	}
//	response.Failed("用户名或密码错误", c)
//}
//
//func (u *AppUser) UserLogin(context *gin.Context) {
//	code := context.PostForm("code")
//	if code == "" {
//		response.Failed("error", context)
//		return
//	}
//	if userInfo := u.Login(code); userInfo != nil {
//		response.Success("ok", userInfo, context)
//		return
//	}
//	response.Failed("error", context)
//}
