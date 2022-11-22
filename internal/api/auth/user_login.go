package auth

import (
	apiResponse "EmployeeSelection/internal/api/response"
	"EmployeeSelection/internal/database/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserLogin(c *gin.Context) {
	loginForm := &Form{}
	err := c.ShouldBind(loginForm)
	if err != nil {
		c.JSON(http.StatusOK, apiResponse.ResponseFail(apiResponse.StatusMsgParamsError, apiResponse.StatusCodeParamsError))
		return
	}
	if loginForm.Account == "" { //空账号
		c.JSON(http.StatusOK, apiResponse.ResponseFail(apiResponse.StatusMsgParamsError, apiResponse.StatusCodeUserInfoError))
		return
	}
	if loginForm.Password == "" { //空密码
		c.JSON(http.StatusOK, apiResponse.ResponseFail(apiResponse.StatusMsgParamsError, apiResponse.StatusCodeUserInfoError))
		return
	}
	if len(loginForm.Password) < 6 { //密码过短
		c.JSON(http.StatusOK, apiResponse.ResponseFail(apiResponse.StatusMsgPasswordShortError, apiResponse.StatusCodeUserInfoError))
		return
	}
	user := &model.User{
		Account: loginForm.Account,
	}
	model.GetModel(user)
	if user.ID == 0 { //没找到用户
		c.JSON(http.StatusOK, apiResponse.ResponseFail(apiResponse.StatusMsgNotFoundUser, apiResponse.StatusCodeUserInfoError))
	} else {
		if !user.CheckPwd(loginForm.Password) { //密码错误
			c.JSON(http.StatusOK, apiResponse.ResponseFail(apiResponse.StatusMsgPasswordError, apiResponse.StatusCodeUserInfoError))
			return
		}
		if user.State == 1 { //1的时候表示被禁用了
			c.JSON(http.StatusOK, apiResponse.ResponseFail(apiResponse.StatusMsgUserAccountDisable, apiResponse.StatusCodeUserInfoError))
			return
		}
		info := UserConvert2LoginInfo(user)
		c.JSON(http.StatusOK, apiResponse.Response(info, apiResponse.StatusMsgSuccess, apiResponse.StatusCodeOk))
	}
}
