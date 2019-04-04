package controllers

import (
	"shorturl/logic"
)

type JumpController struct {
	Controller
}

func (c *JumpController) Run() {
	hashId := c.Ctx.Input.Param(":url")

	logic := logic.ShortUrlLogic{}
	logic.Jump(c.Ctx, hashId)
}
