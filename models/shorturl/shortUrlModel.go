package shorturlModel

import (
	"time"
	db "shorturl/db/mysql"
	"shorturl/global"
)

// Model Struct
type ShortUrl struct {
	Id           int64     `gorm:"column(id);pk" json:"id"`
	HashId       string    `json:"hashId"`
	Url          string    `json:"url"`
	Host         string    `json:"host"`
	CreateUserId string    `json:"createUserId"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"-"`
}

func (a *ShortUrl) TableName() string {
	return "tb_shorturl"
}

func AddNew(id int64, HashId string, Url string, Host string) (ShortUrl, bool) {
	data := ShortUrl{Id: id, HashId: HashId, Url: Url, Host: Host, CreateUserId: global.TokenInfo.Id}

	db.Conn.Create(&data)
	res := db.Conn.NewRecord(&data)

	return data, !res
}

func GetByUrl(url string) (ShortUrl, error) {
	urlInfo := ShortUrl{}

	res := db.Conn.Where("url = ?", url).First(&urlInfo)
	err := res.Error

	return urlInfo, err
}

func GetByHashId(hashId string) (ShortUrl, error) {
	urlInfo := ShortUrl{}

	res := db.Conn.Where("hash_id = ?", hashId).First(&urlInfo)
	err := res.Error

	return urlInfo, err
}
