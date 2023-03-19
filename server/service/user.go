package service

//
//import (
//	"encoding/json"
//	"fmt"
//	"net/http"
//	"server/common"
//	"server/global"
//	"server/model"
//)
//
//type AdminUserService struct {
//}
//
//type UserService struct {
//}
//
//// 商家用户登录
//func (u *AdminUserService) Login(param model.LoginParam) uint64 {
//	var user model.AdminUser
//	global.DB.Where("username = ? and password = ?", param.Username, param.Password).First(&user)
//	return user.Id
//}
//
//// 买家用户登录
//func (u *UserService) Login(code string) *model.UserInfo {
//	var acsJson model.Code2SessionResult
//	acs := model.Code2Session{
//		Code:      code,
//		AppId:     global.Config.Code2Session.AppId,
//		AppSecret: global.Config.Code2Session.AppSecret,
//	}
//	api := "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
//	res, err := http.DefaultClient.Get(fmt.Sprintf(api, acs.AppId, acs.AppSecret, acs.Code))
//	if err != nil {
//		fmt.Println("微信登录凭证校验接口请求错误")
//		return nil
//	}
//	if err := json.NewDecoder(res.Body).Decode(&acsJson); err != nil {
//		fmt.Println("decoder error...")
//		return nil
//	}
//
//	// 查看用户是否已经存在
//	rows := global.DB.Where("open_id = ?", acsJson.OpenId).First(&model.User{}).RowsAffected
//	if rows == 0 {
//		// 不存在，添加用户
//		fmt.Println(acsJson.OpenId)
//		user := model.User{
//			OpenId:  acsJson.OpenId,
//			Status:  1,
//			Created: common.NowTime(),
//		}
//		row := global.DB.Create(&user).RowsAffected
//		if row == 0 {
//			fmt.Println("add app user error...")
//			return nil
//		}
//	}
//	return &model.UserInfo{OpenId: acsJson.OpenId}
//}
