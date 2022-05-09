package service

import (
	"accelerator/entity/errcode"
	"accelerator/entity/response"
	"accelerator/entity/table"
	"accelerator/mysql"
	"accelerator/util"
	"time"

	"github.com/gin-gonic/gin"
)

// UserService 管理用户登录的服务
type UserService struct {
	Email        string `form:"email" json:"email" binding:"required"`
	ChannelId    int64  `form:"channel_id" json:"channel_id" binding:"required"`
	Source       string `form:"source" json:"source" binding:"required"`
	orderService *OrderService
	token        string
}

func NewUserService() *UserService {
	return &UserService{
		orderService: NewOrderService(),
	}
}

// Login 用户登录函数
func (u *UserService) Login(c *gin.Context) response.Response {

	// 设置session
	// service.setSession(c, user)

	user, err := mysql.GetUserByEmail(u.Email)
	if err != nil {
		util.Log().Error("get user by email err: %s", err)
		return errcode.NewErr(errcode.CodeDBError, err)
	}
	// 判断用户是否存在
	if user.ID == 0 {
		user := u.createNewUser()
		id, err := mysql.InsertUser(user)
		if err != nil {
			util.Log().Error("insert user err: %v", err)
			return errcode.NewErr(errcode.CodeDBError, err)
		}
		if err := u.createToken(id); err != nil {
			util.Log().Error("create token err: %v", err)
			return errcode.NewErr(errcode.CodeDBError, err)
		}
	}
	// 如果存在，则查询剩余时间
	remainingTime, err := u.orderService.GetRemainingTimeByUserId(user.ID)
	if err != nil {
		util.Log().Error("get remaining time by user id err: %v", err)
		return errcode.NewErr(errcode.CodeDBError, err)
	}
	// 查询token
	u.getTokenByUserID(user.ID)
	return u.setRsponse(user, remainingTime)
}

// getTokenByUserID 根据用户id获取token
func (u *UserService) getTokenByUserID(id int64) {
	token, err := mysql.GetTokenByUserID(id)
	if err != nil {
		util.Log().Error("get token by user id err: %v", err)
		return
	}
	u.token = token.Token

}

// setRsponse 设置返回值
func (u *UserService) setRsponse(user *table.User, remainingTime int64) response.Response {
	return response.Response{
		Code: errcode.CodeSuccess,
		Msg:  errcode.Text(errcode.CodeSuccess),
		Data: response.UserServiceRsp{
			ID:            user.ID,
			Email:         user.Email,
			Token:         u.token,
			RemainingTime: remainingTime,
		},
	}
}

// createNewUser 创建新用户
func (u *UserService) createNewUser() *table.User {
	user := new(table.User)
	user.Email = u.Email
	user.ChannelId = u.ChannelId
	user.Source = u.Source
	return user
}

// createToken 创建token
func (u *UserService) createToken(id int64) error {
	token := new(table.Token)
	token.UserId = id
	token.Token = util.RandStringRunes(int(id))
	u.token = token.Token
	token.ExpireDate = time.Now().AddDate(1, 0, 0)
	return mysql.InsertToken(token)

}
