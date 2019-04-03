package shorturlModel

import (
	"time"
	db "shorturl/db/mysql"
)

// Model Struct
type Token struct {
	Id        int64     `gorm:"column(id);pk" json:"id"`
	Name      string    `json:"name"`
	AccessKey string    `json:"accessKey"`
	SecretKey string    `json:"secretKey"`
	CreatedAt time.Time `gorm:"column:created_at" json:"-"`
}

func (a *Token) TableName() string {
	return "tb_account"
}

func AddNew(id int64, HashId string, Url string, Host string) (ShortUrl, bool) {
	data := ShortUrl{Id: id, HashId: HashId, Url: Url, Host: Host}

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

func GetByToken(hashId string) (ShortUrl, error) {
	urlInfo := ShortUrl{}

	res := db.Conn.Where("hash_id = ?", hashId).First(&urlInfo)
	err := res.Error

	return urlInfo, err
}
