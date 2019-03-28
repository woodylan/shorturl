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
	Url      string `json:"url"`
	ShortUrl string `json:"shortUrl"`
}

func (this *ShortUrlLogic) Short(c *context.Context, urlString string) (retData RetData) {
	localhost := beego.AppConfig.String("host")
	retData.Url = urlString

	// 正则验证URL
	regex := `(http[s]?|ftp):\/\/([^\/\.]+?)\..+\w$`
	if m, _ := regexp.MatchString(regex, urlString); !m {
		util.ThrowApi(c, -1, "不是合法的URL：")
	}

	// 查询是否存在
	model, err := shorturlModel.GetByUrl(urlString)
	if err == nil {
		retData.ShortUrl = localhost + model.HashId
		return
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

	retData.ShortUrl = localhost + hashId
	return
}

func (this *ShortUrlLogic) Jump(c *context.Context, hashId string) {
	// 查询是否存在
	model, err := shorturlModel.GetByHashId(hashId)
	if err != nil {
		util.ThrowApi(c, -1, "不存在该HashId")
	}

	// 如果用了301，搜索时会直接展示真实地址，无法统计到短地址被点击的次数了，也无法收集用户的Cookie, User Agent等信息
	c.Redirect(302, model.Url)
}

func getHashId(id int64) string {
	minLength, err := beego.AppConfig.Int("hash_minLength")
	if err != nil {
		panic("配置错误:" + err.Error())
	}

	hashIdClass := hashids.NewData()
	hashIdClass.Salt = beego.AppConfig.String("hash_salt")
	hashIdClass.MinLength = minLength

	hashNew, _ := hashids.NewWithData(hashIdClass)
	// 转码
	hashId, _ := hashNew.EncodeInt64([]int64{id})

	////解码
	//d, _ := hashNew.DecodeWithError(e)
	//fmt.Println(d[0])

	return hashId
}
