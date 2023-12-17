package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tdfxlyh/go-gin-api/dal/models"
	"github.com/tdfxlyh/go-gin-api/internal/caller"
	"github.com/tdfxlyh/go-gin-api/internal/utils/output"
	"golang.org/x/crypto/bcrypt"
)

type LoginHandler struct {
	Ctx *gin.Context

	Phone    string
	Password string
	DBUser   models.User

	Resp map[string]interface{}

	Err error
}

func NewLoginHandler(ctx *gin.Context) *LoginHandler {
	return &LoginHandler{
		Ctx:  ctx,
		Resp: make(map[string]interface{}),
	}
}

func Login(ctx *gin.Context) *output.RespStu {
	h := NewLoginHandler(ctx)

	if h.CheckReq(); h.Err != nil {
		caller.Logger.Warn(fmt.Sprintf("[Login-CheckReq] fail, err=%s", h.Err))
		return output.Fail(h.Ctx, output.StatusCodeParamsError)
	}
	if h.CheckDB(); h.Err != nil {
		return output.FailWithMsg(h.Ctx, output.StatusCodeParamsError, h.Err.Error())
	}
	// 获取token
	if h.PackToken(); h.Err != nil {
		caller.Logger.Warn(fmt.Sprintf("[Login-GetToken] err=%s", h.Err))
		return output.Fail(h.Ctx, output.StatusCodeSeverException)
	}
	return output.Success(h.Ctx, h.Resp)
}

func (h *LoginHandler) CheckReq() {
	var reqUser = models.User{}
	h.Ctx.Bind(&reqUser)
	//获取参数
	h.Phone = reqUser.Phone
	h.Password = reqUser.Password
	// 长度校验
	if len(h.Phone) != 11 {
		h.Err = fmt.Errorf("phone long not 11")
		return
	}
	if len(h.Password) < 6 {
		h.Err = fmt.Errorf("password long < 6")
		return
	}
}

func (h *LoginHandler) CheckDB() {
	// 判断手机号是否存在
	caller.LyhTestDB.Where("phone = ? and status=0", h.Phone).First(&h.DBUser)
	if h.DBUser.UID == 0 {
		h.Err = fmt.Errorf("phone not exist")
		return
	}
	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(h.DBUser.Password), []byte(h.Password)); err != nil {
		caller.Logger.Warn(fmt.Sprintf("err=%s", err))
		h.Err = fmt.Errorf("password is error")
		return
	}
}

func (h *LoginHandler) PackToken() {
	token := ""
	token, h.Err = caller.ReleaseToken(h.DBUser)
	if h.Err != nil {
		caller.Logger.Warn(fmt.Sprintf("token generate error: %v", h.Err))
		return
	}
	h.Resp = map[string]interface{}{
		"token": token,
	}
}
