package employeeManagement

import (
	apiResponse "EmployeeSelection/internal/api/response"
	"EmployeeSelection/internal/database"
	"EmployeeSelection/internal/database/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取员工信息列表
func GetUserList(c *gin.Context) {
	var users []model.User
	database.GetDatabase().Find(&users)
	c.JSON(http.StatusOK, apiResponse.Response(users, apiResponse.StatusMsgSuccess, apiResponse.StatusCodeOk))
}
