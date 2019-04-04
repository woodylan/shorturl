package controllers

import (
	"shorturl/logic"
	"github.com/astaxie/beego/validation"
)

type CreateController struct {
	Controller
}

func (c *CreateController) Run() {
	inputData := struct {
		Url string `data:"url" valid:"Required"`
	}{}

	c.GetData(&inputData)
	c.Valid(&inputData)

	valid := validation.Validation{}
	valid.Valid(&inputData)

	logic := logic.ShortUrlLogic{}
	ret := logic.Create(c.Ctx, inputData.Url)

	c.Data["json"] = RetData{0, "success", ret}
	c.ServeJSON()
}
