package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tdfxlyh/go-gin-api/dal/models"
	"github.com/tdfxlyh/go-gin-api/internal/caller"
	"github.com/tdfxlyh/go-gin-api/internal/utils"
	"github.com/tdfxlyh/go-gin-api/internal/utils/output"
	"golang.org/x/crypto/bcrypt"
)

type RegisterHandler struct {
	Ctx *gin.Context

	UserName string
	Phone    string
	Password string

	Err error
}

func NewRegisterHandler(ctx *gin.Context) *RegisterHandler {
	return &RegisterHandler{
		Ctx: ctx,
	}
}

func Register(ctx *gin.Context) *output.RespStu {
	h := NewRegisterHandler(ctx)
	// 校验参数
	if msg := h.CheckReq(); msg != "" {
		caller.Logger.Warn(fmt.Sprintf("[Register-CheckReq] params fail, err=%s", msg))
		return output.FailWithMsg(h.Ctx, output.StatusCodeParamsError, msg)
	}
	// 判断用户是否已经存在
	if h.CheckUserExists() {
		return output.FailWithMsg(h.Ctx, output.StatusCodeParamsError, "用户已存在")
	}
	// 创建用户
	if h.CreateUser(); h.Err != nil {
		caller.Logger.Warn(fmt.Sprintf("[Register-CreateUser] fail, err=%s", h.Err))
		return output.Fail(h.Ctx, output.StatusCodeSeverException)
	}
	return output.Success(h.Ctx, "注册成功")
}

func (h *RegisterHandler) CheckReq() string {
	var reqUser = models.User{}
	h.Ctx.Bind(&reqUser)

	caller.Logger.Info(utils.GetStuStr(reqUser))

	//获取参数
	h.UserName = reqUser.UserName
	h.Phone = reqUser.Phone
	h.Password = reqUser.Password

	// 长度校验
	if len(h.Phone) != 11 {
		return "手机号位数错误"
	}
	if len(h.Password) < 6 {
		return "密码长度小于6"
	}
	if len(h.UserName) == 0 {
		h.UserName = utils.RandString(5)
	}
	return ""
}

func (h *RegisterHandler) CheckUserExists() bool {
	// 判断手机号是否已经存在
	var user models.User
	caller.LyhTestDB.Where("phone=? and status=0", h.Phone).First(&user)
	return user.UID != 0
}

func (h *RegisterHandler) CreateUser() {
	hasePassword, err := bcrypt.GenerateFromPassword([]byte(h.Password), bcrypt.DefaultCost)
	if err != nil {
		h.Err = fmt.Errorf("加密失败")
		return
	}
	newUser := &models.User{
		UserName: h.UserName,
		Phone:    h.Phone,
		Password: string(hasePassword),
		Avatar:   "https://img0.baidu.com/it/u=1821253856,3774998416&fm=253&fmt=auto&app=138&f=JPEG?w=500&h=500",
	}
	h.Err = caller.LyhTestDB.Create(newUser).Error
}
