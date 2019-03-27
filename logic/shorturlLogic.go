package logic

import (
	"github.com/astaxie/beego/context"
	"github.com/speps/go-hashids"
	"shorturl/modules/util"
	"github.com/astaxie/beego"
	"net/url"
	"regexp"
	"shorturl/models/shorturl"
)

type ShortUrlLogic struct {
}

type RetData struct {
	ShortUrl string `json:"url"`
	Host     string `json:"host"`
}

func (this *ShortUrlLogic) Short(c *context.Context, urlString string) (retData RetData) {
	// 正则验证URL
	regex := `(http[s]?|ftp):\/\/([^\/\.]+?)\..+\w$`
	if m, _ := regexp.MatchString(regex, urlString); !m {
		util.ThrowApi(c, -1, "不是合法的URL：")
	}

	// 解析URL
	urlParse, _ := url.Parse(urlString)
	host := urlParse.Host

	// 通过发号器获取ID
	redisKey := beego.AppConfig.String("redis_key")
	shortUrlId := util.GetRedisNum(redisKey)

	// 计算发号器ID HASH值
	hashId := getHashId(shortUrlId)

	// 存入数据库
	shorturlModel.AddNew(shortUrlId, hashId, urlString, host)

	retData.ShortUrl = hashId
	return
}

func getHashId(id int64) string {

	hashIdClass := hashids.NewData()
	hashIdClass.Salt = "Salt"
	hashIdClass.MinLength = 8

	hashNew, _ := hashids.NewWithData(hashIdClass)
	// 转码
	hashId, _ := hashNew.EncodeInt64([]int64{id})

	////解码
	//d, _ := hashNew.DecodeWithError(e)
	//fmt.Println(d[0])

	return hashId
}
