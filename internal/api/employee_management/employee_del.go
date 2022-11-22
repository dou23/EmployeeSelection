package employeeManagement

import (
	apiResponse "EmployeeSelection/internal/api/response"
	"EmployeeSelection/internal/database"
	"EmployeeSelection/internal/database/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type EmployeeDelForm struct {
	Account string `json:"account"`
}

func UserDel(c *gin.Context) {
	employeeDelForm := &EmployeeDelForm{}
	c.ShouldBind(employeeDelForm)
	user := &model.User{
		Account: employeeDelForm.Account,
	}
	db := database.GetDatabase().Find(user)
	if user.ID == 0 || db.Error != nil {
		c.JSON(http.StatusOK, apiResponse.ResponseFail(apiResponse.StatusMsgNotFoundUser, apiResponse.StatusCodeUserInfoError))
		return
	}
	database.GetDatabase().Delete(user)
	c.JSON(http.StatusOK, apiResponse.ResponseOk(apiResponse.StatusMsgSuccess))
}
