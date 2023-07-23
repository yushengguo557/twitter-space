package models

import "time"

type TwitterSpace struct {
	ID               string    `gorm:"column:id;primary_key;NOT NULL;comment:'主键 Space ID'" json:"id"`
	CreatorId        string    `gorm:"column:creator_id;default:'';comment:'创建者id'" json:"creator_id"`
	ParticipantCount int       `gorm:"column:participant_count;default:0;comment:'参与人数'" json:"participant_count"`
	Title            string    `gorm:"column:title;default:;comment:'Space标题'" json:"title"`
	Description      string    `gorm:"column:description;default:;comment:'描述'" json:"description"`
	Lang             string    `gorm:"column:lang;default:en;comment:'Space语言'" json:"lang"`
	Url              string    `gorm:"column:url;default:;comment:'链接 跳转推特'" json:"url"`
	Status           string    `gorm:"column:status;default:1;comment:'状态: 0=end 1=live 2=Scheduled'" json:"status"`
	StartedAt        time.Time `gorm:"column:started_at;default:NULL;comment:'创建时间'" json:"started_at"`
	ScheduledStart   time.Time `gorm:"column:scheduled_start;default:NULL;comment:'预定开始时间'" json:"scheduled_start"`
	EndedAt          time.Time `gorm:"column:ended_at;default:NULL;comment:'结束时间'" json:"ended_at"`
	UpdatedAt        time.Time `gorm:"column:updated_at;default:NULL;comment:'更新时间'" json:"updated_at"`
	DeletedAt        time.Time `gorm:"column:deleted_at;default:NULL;comment:'删除时间'" json:"deleted_at"`
	DataStatus       []byte    `gorm:"column:data_status;type:bit(1);default:1;NOT NULL;comment:'数据状态:1=正常,0=删除'" json:"data_status"`
	Tag              string    `gorm:"column:tag;type:char(16);default:'NFT';NOT NULL;comment:'标签 NFT,WEB3,Game,MetaVerse,DeFi';" json:"tag"`
}

func (ts *TwitterSpace) TableName() string {
	return "im_twitter_space"
}
