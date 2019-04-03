package shorturlModel

import (
	"time"
	db "shorturl/db/mysql"
)

// Model Struct
type TokenModel struct {
	Id        int64     `gorm:"column(id);pk" json:"id"`
	Name      string    `json:"name"`
	Token     string    `json:"token"`
	CreatedAt time.Time `gorm:"column:created_at" json:"-"`
}

func (a *TokenModel) TableName() string {
	return "tb_token"
}

func GetByToken(token string) (TokenModel, error) {
	tokenInfo := TokenModel{}

	res := db.Conn.Where("token = ?", token).First(&tokenInfo)
	err := res.Error

	return tokenInfo, err
}
