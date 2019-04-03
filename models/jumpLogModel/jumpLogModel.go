package jumpLogModel

import (
	"time"
	db "shorturl/db/mysql"
)

// Model Struct
type JumpLogModel struct {
	Id        int64     `gorm:"column(id);pk" json:"id"`
	UrlId     int64     `json:"urlId"`
	UserAgent string    `json:"userAgent"`
	Ip        string    `json:"ip"`
	Referer   string    `json:"Referer"`
	CreatedAt time.Time `gorm:"column:created_at" json:"-"`
}

func (a *JumpLogModel) TableName() string {
	return "tb_jump_log"
}

func AddNew(UrlId int64, UserAgent, Ip, Referer string) (JumpLogModel, bool) {
	data := JumpLogModel{UrlId: UrlId, UserAgent: UserAgent, Ip: Ip, Referer: Referer}

	db.Conn.Create(&data)
	res := db.Conn.NewRecord(&data)

	return data, !res
}
