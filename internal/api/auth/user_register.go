package auth

import (
	apiResponse "EmployeeSelection/internal/api/response"
	"EmployeeSelection/internal/database/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRegister(c *gin.Context) {
	registerForm := &Form{}
	c.ShouldBind(registerForm)
	if registerForm.Account == "" { //空账号
		c.JSON(http.StatusOK, apiResponse.ResponseFail(apiResponse.StatusMsgParamsError, apiResponse.StatusCodeUserInfoError))
		return
	}
	if registerForm.Password == "" { //空密码
		c.JSON(http.StatusOK, apiResponse.ResponseFail(apiResponse.StatusMsgParamsError, apiResponse.StatusCodeUserInfoError))
		return
	}
	if len(registerForm.Password) < 6 { //密码过短
		c.JSON(http.StatusOK, apiResponse.ResponseFail(apiResponse.StatusMsgPasswordShortError, apiResponse.StatusCodeUserInfoError))
		return
	}
	user := model.User{
		Account: registerForm.Account,
		Username: func() string {
			if registerForm.Username == "" {
				return registerForm.Account
			} else {
				return registerForm.Username
			}
		}(),
	}
	var err error
	_, err = user.Register(registerForm.Password)
	if err != nil {
		c.JSON(http.StatusOK, apiResponse.ResponseFail("用户注册失败: "+err.Error(), 500))
		return
	}
	c.JSON(http.StatusOK, apiResponse.ResponseOk(apiResponse.StatusMsgUserRegisterSuccess))
}
