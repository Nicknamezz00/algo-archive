package api

import (
	"algo-archive/internal/service"
	"algo-archive/pkg/app"
	"algo-archive/pkg/errcode"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Register(c *gin.Context) {
	param := service.RegisterReq{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValidate(c, &param)
	if !valid {
		logrus.Errorf("app.BindAndValidate errs: %v", errs)
		response.ReplyError(errcode.InvalidParameters.WithDetails(errs.Errors()...))
		return
	}

	// validate username
	err := service.ValidateUsername(param.Username)
	if err != nil {
		logrus.Errorf("service.Register err: %v", err)
		response.ReplyError(err.(*errcode.Error))
		return
	}

	// validate password
	err = service.ValidatePassword(param.Password)
	if err != nil {
		logrus.Errorf("service.Register err: %v", err)
		response.ReplyError(err.(*errcode.Error))
		return
	}

	user, err := service.Register(
		param.Username,
		param.Password,
	)
	if err != nil {
		logrus.Errorf("service.Register err: %v", err)
		response.ReplyError(errcode.UserRegisterFailed)
		return
	}
	response.Reply(gin.H{
		"id":       user.ID,
		"username": user.Username,
	})
}

func Login(c *gin.Context) {
	param := service.AuthRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValidate(c, &param)
	if !valid {
		logrus.Errorf("app.BindAndValidate errs: %v", errs)
		response.ReplyError(errcode.InvalidParameters.WithDetails(errs.Errors()...))
		return
	}

	user, err := service.DoLogin(c, &param)
	if err != nil {
		logrus.Errorf("service.DoLogin err: %v", err)
		response.ReplyError(err.(*errcode.Error))
		return
	}

	token, err := app.GenerateToken(user)
	if err != nil {
		logrus.Errorf("app.GenerateToken err: %v", err)
		response.ReplyError(errcode.UnauthorizedTokenGenFailed)
		return
	}

	response.Reply(gin.H{
		"token": token,
	})
}
