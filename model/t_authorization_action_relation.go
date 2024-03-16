package model

// TAuthorizationActionRelation 凭证授权表
type TAuthorizationActionRelation struct {
	ID       uint `gorm:"column:id; type:int unsigned; primaryKey; autoIncrement; comment:主码" json:"id,omitempty"`
	AuthID   uint `gorm:"column:auth_id; type:int unsigned not null; comment:凭证ID，外码"`
	ActionID uint `gorm:"column:action_id; type:int unsigned not null; comment:授权项ID，外码"`
}

func (*TAuthorizationActionRelation) TableName() string {
	return "t_authorization_action_relation"
}
