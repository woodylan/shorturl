package controllers

import "shorturl/logic"

type ShortUrlController struct {
	Controller
}

func (c *ShortUrlController) Create() {
	url := c.GetString("url")

	// 校验URL

	logic := logic.ShortUrlLogic{}
	ret := logic.Short(c.Ctx, url)

	c.Data["json"] = RetData{0, "success", ret}
	c.ServeJSON()
}
