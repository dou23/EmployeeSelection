package selection

import (
	apiResponse "EmployeeSelection/internal/api/response"
	"EmployeeSelection/internal/database"
	"EmployeeSelection/internal/database/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type selectionDelForm struct {
	Id int64 `json:"id"`
}

func SelectionDel(c *gin.Context) {
	delForm := &selectionDelForm{}
	if delForm.Id == 0 {
		c.JSON(http.StatusOK, apiResponse.ResponseFail("请输入评选活动id", apiResponse.StatusCodeParamsError))
		return
	}
	selection := &model.Selection{}
	selection.ID = delForm.Id
	tx := database.GetDatabase().Find(selection)
	if tx.Error != nil {
		c.JSON(http.StatusOK, apiResponse.ResponseFail("评选活动不存在", apiResponse.StatusCodeParamsError))
		return
	}
}
