package handler

import (
	"github.com/LianJianTech/lj-go-common/errno"
	"github.com/LianJianTech/lj-go-common/log"
	"github.com/gin-gonic/gin"

	"lj-go-practice/model"
	"lj-go-practice/service"
)

func LoginAccount(c *gin.Context) {
	form := &model.LoginForm{}
	if err := c.ShouldBindJSON(form); err != nil {
		log.Errorf(err, "c.ShouldBindJSON error")
		SendResponse(c, errno.AuthError, nil)
		return
	}

	service := service.NewAccountService()

	resp, err := service.LoginAccount(form)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}

	SendResponse(c, nil, resp)
}

func UpdateAccountPwd(c *gin.Context) {
	model := &model.UpdatePwdForm{}
	if err := c.ShouldBindJSON(model); err != nil {
		log.Errorf(err, "c.ShouldBindJSON error")
		SendResponse(c, errno.AuthError, nil)
		return
	}
	if len(model.NewPwd) < 8 {
		SendResponse(c, errno.AuthError, nil)
		return
	}

	service := service.NewAccountService()

	if err := service.UpdateAccountPwd(model); err != nil {
		SendResponse(c, err, nil)
		return
	}

	SendResponse(c, nil, nil)
}
