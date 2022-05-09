package controllers

import (
	"github.com/dchest/captcha"
)

type CaptchaController struct {
	BaseController
}

func (c *CaptchaController) Captcha() {
	captcha.Server(captcha.StdWidth, captcha.StdHeight)
}
