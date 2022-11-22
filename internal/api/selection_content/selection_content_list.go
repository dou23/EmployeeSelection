package selectionContent

import (
	apiResponse "EmployeeSelection/internal/api/response"
	"EmployeeSelection/internal/database"
	"EmployeeSelection/internal/database/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SelectionContentList(c *gin.Context) {
	var selectionContents []model.SelectionContent
	tx := database.GetDatabase().Find(&selectionContents)
	if tx.Error != nil {
		c.JSON(http.StatusOK, apiResponse.ResponseFail(tx.Error.Error(), apiResponse.StatusCodeUserInfoError))
		return
	}
	for i := range selectionContents {
		var selectionContentDetails []model.SelectionContentDetail
		tx = database.GetDatabase().Where("selection_content_id = ?", selectionContents[i].ID).Find(&selectionContentDetails)
		if tx.Error != nil {
			c.JSON(http.StatusOK, apiResponse.ResponseFail(tx.Error.Error(), apiResponse.StatusCodeUserInfoError))
			return
		}
		selectionContents[i].ContentDetail = selectionContentDetails
	}
	c.JSON(http.StatusOK, apiResponse.Response(selectionContents, apiResponse.StatusMsgSuccess, apiResponse.StatusCodeOk))
}
