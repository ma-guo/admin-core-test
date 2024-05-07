package models

import "time"

type (
	User struct {
		Id       int64     `xorm:"NOT NULL PK AUTOINCR INT(11)"` // 排序 ID
		OpenId   string    `xorm:"VARCHAR(64) NOT NULL index"`   // openid
		Nickname string    `xorm:"VARCHAR(255)"`                 // 用户昵称
		Avatar   string    `xorm:"TEXT"`                         // 用户头像
		UnionId  string    `xorm:"VARCHAR(64)"`                  // 用户 union id
		Ip       string    `xorm:"VARCHAR(64)"`                  // 用户登录 ip
		CreateAt time.Time `xorm:"created"`                      // 创建时间
		UpdateAt time.Time `xorm:"updated"`                      // 更新时间
		DeleteAt time.Time `xorm:"deleted"`                      // 删除时间
	}
)

func GetSyncModels() []interface{} {
	return []interface{}{
		new(User),
	}
}
