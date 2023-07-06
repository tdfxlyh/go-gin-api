package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tdfxlyh/go-gin-api/dal/models"
	"github.com/tdfxlyh/go-gin-api/internal/caller"
	"github.com/tdfxlyh/go-gin-api/internal/model/res"
	"github.com/tdfxlyh/go-gin-api/internal/utils"
	"github.com/tdfxlyh/go-gin-api/internal/utils/output"
	"golang.org/x/crypto/bcrypt"
)

type RegisterHandler struct {
	Ctx *gin.Context

	Name     string
	Phone    string
	Password string

	Err error
}

func NewRegisterHandler(ctx *gin.Context) *RegisterHandler {
	return &RegisterHandler{
		Ctx: ctx,
	}
}

func Register(ctx *gin.Context) *res.RespStu {
	h := NewRegisterHandler(ctx)
	// 校验参数
	if h.CheckReq(); h.Err != nil {
		fmt.Println(fmt.Sprintf("[Register-CheckReq] params fail, err=%s", h.Err))
		return res.Fail(h.Ctx, output.StatusCodeParamsError)
	}
	// 判断用户是否已经存在
	if h.CheckUser(); h.Err != nil {
		fmt.Println(fmt.Sprintf("[Register-CheckUser] err=%s", h.Err))
		return res.FailWithMsg(h.Ctx, output.StatusCodeDBError, "用户已存在")
	}

	// 创建用户
	if h.CreateUser(); h.Err != nil {
		fmt.Println(fmt.Sprintf("[Register-CreateUser] fail, err=%s", h.Err))
		return res.Fail(h.Ctx, output.StatusCodeSeverException)
	}
	return res.Success(h.Ctx, "注册成功")
}

func (h *RegisterHandler) CheckReq() {
	var reqUser = models.User{}
	h.Ctx.Bind(&reqUser)

	fmt.Printf("%s", utils.GetStuStr(reqUser))

	//获取参数
	h.Name = reqUser.Name
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
	if len(h.Name) == 0 {
		h.Name = utils.RandString(5)
	}
}

func (h *RegisterHandler) CheckUser() {
	// 判断手机号是否已经存在
	var user models.User
	caller.LyhTestDB.Where("phone = ?", h.Phone).First(&user)
	if user.ID != 0 {
		h.Err = fmt.Errorf("user exist")
		return
	}
}

func (h *RegisterHandler) CreateUser() {
	hasePassword, err := bcrypt.GenerateFromPassword([]byte(h.Password), bcrypt.DefaultCost)
	if err != nil {
		h.Err = fmt.Errorf("加密失败")
		return
	}
	newUser := &models.User{
		Name:     h.Name,
		Phone:    h.Phone,
		Password: string(hasePassword),
	}
	caller.LyhTestDB.Create(newUser)
}
