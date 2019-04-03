package controllers

import (
	"shorturl/logic"
	"github.com/astaxie/beego/validation"
	)

type ShortUrlController struct {
	Controller
}

func (c *ShortUrlController) Create() {
	inputData := struct {
		Url string `data:"url" valid:"Required"`
	}{}

	c.GetData(&inputData)
	c.Valid(&inputData)

	valid := validation.Validation{}
	valid.Valid(&inputData)

	logic := logic.ShortUrlLogic{}
	ret := logic.Short(c.Ctx, inputData.Url)

	c.Data["json"] = RetData{0, "success", ret}
	c.ServeJSON()
}

func (c *ShortUrlController) Jump() {
	hashId := c.Ctx.Input.Param(":url")

	logic := logic.ShortUrlLogic{}
	logic.Jump(c.Ctx, hashId)
}
