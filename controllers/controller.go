package controllers

import (
	"github.com/astaxie/beego"
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
