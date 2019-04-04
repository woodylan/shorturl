package controllers

import (
	"shorturl/logic"
	"github.com/astaxie/beego/validation"
	)

type QueryController struct {
	Controller
}

func (c *QueryController) Run() {
	inputData := struct {
		ShortUrl string `data:"shortUrl" valid:"Required"`
	}{}

	c.GetData(&inputData)
	c.Valid(&inputData)

	valid := validation.Validation{}
	valid.Valid(&inputData)

	logic := logic.ShortUrlLogic{}
	ret := logic.Query(c.Ctx, inputData.ShortUrl)

	c.Data["json"] = RetData{0, "success", ret}
	c.ServeJSON()
}