package models

import (
	"time"
)

type TwitterUser struct {
	ID              string     `gorm:"column:id;primary_key;NOT NULL;comment:'主键id 推特用户id'" json:"id"`
	Name            string     `gorm:"column:name;default:'';comment:'姓名'" json:"name"`
	Username        string     `gorm:"column:username;default:'';comment:'用户名'" json:"username"`
	Location        string     `gorm:"column:location;default:'';comment:'位置'" json:"location"`
	Description     string     `gorm:"column:description;default:'';comment:'描述'" json:"description"`
	ProfileImageUrl string     `gorm:"column:profile_image_url;default:'';comment:'头像链接'" json:"profile_image_url"`
	SpaceId         string     `gorm:"column:space_id;default:'';comment:'所属Space的ID'" json:"space_id"`
	CreatedAt       *time.Time `gorm:"column:created_at;default:NULL;comment:'创建时间'" json:"created_at"`
	UpdatedAt       *time.Time `gorm:"column:updated_at;default:NULL;comment:'更新时间'" json:"updated_at"`
	DeletedAt       *time.Time `gorm:"column:deleted_at;default:NULL;comment:'删除时间'" json:"deleted_at"`
	DataStatus      []byte     `gorm:"column:data_status;type:bit(1);default:1;NOT NULL;comment:'数据状态:1=正常,0=删除'" json:"data_status"`
	Url             string     `gorm:"column:url;default:'';comment:'用户主页链接 (推特)';" json:"url"`
}

func (tu *TwitterUser) TableName() string {
	return "im_twitter_user"
}
