package controllers

import (
	"shorturl/logic"
)

type ShortUrlController struct {
	Controller
}

func (c *ShortUrlController) Create() {
	url := c.GetString("url")

	logic := logic.ShortUrlLogic{}
	ret := logic.Short(c.Ctx, url)

	c.Data["json"] = RetData{0, "success", ret}
	c.ServeJSON()
}

func (c *ShortUrlController) Jump() {
	hashId := c.Ctx.Input.Param(":url")

	logic := logic.ShortUrlLogic{}
	logic.Jump(c.Ctx, hashId)
}
