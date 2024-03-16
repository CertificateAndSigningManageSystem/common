package model

const (
	// 获取应用信息
	TAuthorizationAction_ID_GetAppInfo uint = 1 + iota
	// 获取签名信息信息
	TAuthorizationAction_ID_GetSignJobInfo
	// 使用签名服务
	TAuthorizationAction_ID_SignJob
)

// TAuthorizationAction 凭证授权项表
type TAuthorizationAction struct {
	ID       uint   `gorm:"column:id; type:int unsigned; primaryKey; autoIncrement; comment:主码" json:"id,omitempty"`
	Name     string `gorm:"column:name; type:varchar(64) not null; comment:名称" json:"name,omitempty"`
	Platform uint8  `gorm:"column:platform; type:tinyint unsigned not null; comment:适用应用平台，1=Windows,2=Android,3=Apple,4=所有" json:"platform,omitempty"`
}

func (*TAuthorizationAction) TableName() string {
	return "t_authorization_action"
}