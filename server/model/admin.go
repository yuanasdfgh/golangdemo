package model

// 用户数据映射模型
type AdminUser struct {
	Id       uint64 `gorm:"primaryKey"` // 用户编号
	OpenId   string `gorm:"open_id"`    // 微信用户唯一标识
	Avatar   string `gorm:"avatar"`     // 用户头像
	NickName string `gorm:"nick_Name"`  // 用户昵称
	Username string `gorm:"username"`   // 用户名称
	Password string `gorm:"password"`   // 用户密码
	Status   uint   `gorm:"status"`     // 用户状态
	Created  string `gorm:"created"`    // 创建时间
	Updated  string `gorm:"updated"`    // 更新时间
}
