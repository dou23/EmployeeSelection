package auth

import (
	apiResponse "EmployeeSelection/internal/api/response"
	"EmployeeSelection/internal/config"
	"EmployeeSelection/internal/database/model"
	"EmployeeSelection/internal/database/model/admin"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Form 登录注册
type Form struct {
	Account  string `form:"account" json:"account"`
	Password string `form:"password" json:"password"`
	Username string `form:"username" json:"username"`
}

type Info struct {
	Username string `json:"username"`
	Account  string `json:"account"`
	Token    string `json:"token"`
	Expires  int    `json:"expires"`
	Avatar   string `json:"avatar"`
}

func AdminConvert2LoginInfo(u *admin.Admin) *Info {
	jwtInfo := u.GetJwtInfo()
	info := &Info{
		Account:  u.Account,
		Username: u.Username,
		Avatar:   u.Avatar,
		Token:    jwtInfo.Token,
		Expires:  jwtInfo.Expiresin,
	}
	return info
}

func UserConvert2LoginInfo(u *model.User) *Info {
	jwtInfo := u.GetJwtInfo()
	info := &Info{
		Account:  u.Account,
		Avatar:   u.Avatar,
		Username: u.Username,
		Token:    jwtInfo.Token,
		Expires:  jwtInfo.Expiresin,
	}

	return info
}

func CheckUserByJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get(config.AuthToken)
		if token == "" {
			c.JSON(http.StatusOK, apiResponse.ResponseFail("缺少token参数，无权访问", 401))
			c.Abort()
			return
		}
		u := model.User{}
		user, err := u.GetUserByJwt(token)
		if err != nil {
			c.JSON(http.StatusOK, apiResponse.ResponseFail("鉴权错误:"+err.Error(), 401))
			return
		}
		c.Set("user", user)
	}
}

func CheckAdminByJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get(config.AuthToken)
		if token == "" {
			c.JSON(http.StatusOK, apiResponse.ResponseFail("缺少token参数，无权访问", 401))
			c.Abort()
			return
		}
		u := admin.Admin{}
		user, err := u.GetUserByJwt(token)
		if err != nil {
			c.JSON(http.StatusOK, apiResponse.ResponseFail("鉴权错误:"+err.Error(), 401))
			return
		}
		c.Set("user", user)
	}
}
