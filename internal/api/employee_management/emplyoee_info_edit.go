package employeeManagement

import (
	apiResponse "EmployeeSelection/internal/api/response"
	"EmployeeSelection/internal/database"
	"EmployeeSelection/internal/database/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 普通用户
type UserEditForm struct {
	Id       int64  `json:"id"`
	Account  string `json:"account"`
	Password string `json:"password"`
	Username string `json:"username"`
	State    int    `json:"state"`
}

// 编辑员工信息
func UserInfoEdit(c *gin.Context) {
	userEditForm := &UserEditForm{}
	c.ShouldBind(userEditForm)
	if userEditForm.Id == 0 {
		c.JSON(http.StatusOK, apiResponse.ResponseFail(apiResponse.StatusMsgNotFoundUser, apiResponse.StatusCodeUserInfoError))
		return
	}
	user := &model.User{}
	user.ID = userEditForm.Id
	db := database.GetDatabase().Find(user)
	if db.Error != nil {
		c.JSON(http.StatusOK, apiResponse.ResponseFail(apiResponse.StatusMsgNotFoundUser, apiResponse.StatusCodeUserInfoError))
		return
	}
	updateUser := &model.User{}
	if userEditForm.Account != "" {
		updateUser.Account = userEditForm.Account
	}
	if userEditForm.Password != "" || len(userEditForm.Password) >= 6 {
		updateUser.Password = userEditForm.Password
	} else {
		c.JSON(http.StatusOK, apiResponse.ResponseFail(apiResponse.StatusMsgEditError, apiResponse.StatusCodeUserInfoError))
		return
	}
	if userEditForm.Username != "" {
		updateUser.Username = userEditForm.Username
	}
	updateUser.State = userEditForm.State
	db = database.GetDatabase().Model(user).Select("*").Updates(updateUser)
	if db.Error != nil {
		c.JSON(http.StatusOK, apiResponse.ResponseFail(apiResponse.StatusMsgEditError, apiResponse.StatusCodeUserInfoError))
		return
	}
	c.JSON(http.StatusOK, apiResponse.ResponseOk(apiResponse.StatusMsgSuccess))
}
