package selection

import (
	apiResponse "EmployeeSelection/internal/api/response"
	"EmployeeSelection/internal/database"
	"EmployeeSelection/internal/database/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// / 评选活动列表
func SelectionList(c *gin.Context) {
	var selection []model.Selection
	tx := database.GetDatabase().Find(&selection)
	if tx.Error != nil {
		c.JSON(http.StatusOK, apiResponse.ResponseFail(tx.Error.Error(), apiResponse.StatusCodeUserInfoError))
		return
	}
	c.JSON(http.StatusOK, apiResponse.Response(selection, apiResponse.StatusMsgSuccess, apiResponse.StatusCodeOk))
}
