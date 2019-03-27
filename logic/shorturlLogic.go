package logic

import (
	"github.com/astaxie/beego/context"
	"github.com/speps/go-hashids"
	"shorturl/modules/util"
	"github.com/astaxie/beego"
)

type ShortUrlLogic struct {
}

type RetData struct {
	ShortUrl string `json:"url"`
	Host     string `json:"host"`
}

func (this *ShortUrlLogic) Short(c *context.Context, url string) (retData RetData) {
	redisKey := beego.AppConfig.String("redis_key")
	shortUrlId := util.GetRedisNum(redisKey)

	hashIdClass := hashids.NewData()
	hashIdClass.Salt = "Salt"
	hashIdClass.MinLength = 8

	hashNew, _ := hashids.NewWithData(hashIdClass)
	// 转码
	hashId, _ := hashNew.EncodeInt64([]int64{shortUrlId})

	////解码
	//d, _ := h.DecodeWithError(e)
	//fmt.Println(d[0])

	retData.ShortUrl = hashId

	return
}
