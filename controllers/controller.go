package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"github.com/astaxie/beego/validation"
	"log"
	"shorturl/modules/util"
	"shorturl/define/retcode"
)

type Controller struct {
	beego.Controller
}

type RetData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (c *Controller) Prepare() {
}

// 获取request body数据并转换成指定struct结构
func (c *Controller) GetData(RequestData interface{}) {
	data := c.Ctx.Input.RequestBody

	//string 转 struct
	if err := json.Unmarshal([]byte(data), &RequestData); err == nil {
	}
}

// 参数校验
func (c *Controller) Valid(inputData interface{}) {
	valid := validation.Validation{}
	b, err := valid.Valid(inputData)
	if err != nil {
		// handle error
	}
	if !b {
		// 处理抛出验证不通过
		for _, err := range valid.Errors {
			util.ThrowApi(c.Ctx, retcode.ErrParam, err.Key+" "+err.Message)
			log.Println(err.Key, err.Message)
		}
	}
}
