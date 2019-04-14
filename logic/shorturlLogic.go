package logic

import (
	"github.com/astaxie/beego/context"
	"github.com/speps/go-hashids"
	"shorturl/modules/util"
	"github.com/astaxie/beego"
	"net/url"
	"shorturl/models/shorturl"
	"shorturl/define/retcode"
	redisDB "shorturl/db/redis"
	"strings"
	"shorturl/models/jumpLogModel"
	"strconv"
)

type ShortUrlLogic struct {
}

type RetData struct {
	LongUrl  string `json:"longUrl"`
	ShortUrl string `json:"shortUrl"`
}

type QueryRetData struct {
	LongUrl  string `json:"longUrl"`
	ShortUrl string `json:"shortUrl"`
}

// 创建短链接
func (this *ShortUrlLogic) Create(c *context.Context, urlString string) (retData RetData) {
	localhost := beego.AppConfig.String("host")
	retData.LongUrl = urlString

	// 解析URL
	urlParse, _ := url.Parse(urlString)
	host := urlParse.Host

	if host == "" {
		util.ThrowApi(c, retcode.ErrValidateFailUrl, "不是合法的URL")
		return
	}

	// 查询是否存在
	model, err := shorturlModel.GetByUrl(urlString)
	if err == nil {
		retData.ShortUrl = localhost + model.HashId
		return
	}

	// 通过发号器获取ID
	redisKey := beego.AppConfig.String("redis_key")
	shortUrlId := util.GetRedisNum(redisKey)

	// 计算发号器ID HASH值
	hashId := getHashId(shortUrlId)

	// 存入数据库
	shorturlModel.AddNew(shortUrlId, hashId, urlString, host)

	hashMap := map[string]interface{}{
		"url": urlString,
		"id":  shortUrlId,
	}

	// 写入到Redis
	redisDB.RedisConnect.HMSet(redisKey+":Code:"+hashId, hashMap)

	retData.ShortUrl = localhost + hashId
	return
}

func (this *ShortUrlLogic) Jump(c *context.Context, hashId string) {
	redisKey := beego.AppConfig.String("redis_key")
	hashValue := redisDB.RedisConnect.HGetAll(redisKey + ":Code:" + hashId).Val()
	if len(hashValue) == 0 {
		util.ThrowApi(c, retcode.ErrHashIdNotFound, "不存在该HashId")
		return
	}

	userAgent := c.Request.Header.Get("User-Agent")
	remoteAddr := c.Request.RemoteAddr
	referer := c.Request.Referer()

	pos := strings.Index(remoteAddr, ":")
	ip := string([]rune(remoteAddr)[:pos])

	// 添加访问日志
	urlId, _ := strconv.ParseInt(hashValue["id"], 10, 64)
	jumpLogModel.AddNew(urlId, userAgent, ip, referer)

	// 如果用了301，搜索时会直接展示真实地址，无法统计到短地址被点击的次数了，也无法收集用户的Cookie, User Agent等信息
	c.Redirect(302, hashValue["url"])
}

// 查询短链接
func (this *ShortUrlLogic) Query(c *context.Context, urlString string) (retData QueryRetData) {
	// 解析URL
	urlParse, _ := url.Parse(urlString)
	uri := urlParse.Path

	if uri == "" {
		util.ThrowApi(c, retcode.ErrUrlNotFound, "URI不存在")
		return
	}

	hashId := string([]rune(uri)[1:])

	// 查询是否存在
	model, err := shorturlModel.GetByHashId(hashId)
	if err != nil {
		util.ThrowApi(c, retcode.ErrHashIdNotFound, "不存在该HashId")
		return
	}

	retData.ShortUrl = urlString
	retData.LongUrl = model.Url

	return
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

	return hashId
}
