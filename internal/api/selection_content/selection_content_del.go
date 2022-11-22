package selectionContent

import (
	apiResponse "EmployeeSelection/internal/api/response"
	"EmployeeSelection/internal/database"
	"EmployeeSelection/internal/database/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type SelectionContentDelForm struct {
	Id int64 `json:"id"`
}

func SelectionContentDel(c *gin.Context) {
	delForm := &SelectionContentDelForm{}
	c.ShouldBind(delForm)
	if delForm.Id == 0 {
		c.JSON(http.StatusOK, apiResponse.ResponseFail(apiResponse.StatusMsgParamsError, apiResponse.StatusCodeParamsError))
		return
	}
	err := database.GetDatabase().Transaction(func(tx *gorm.DB) error {
		db := database.GetDatabase().Where("id = ?", delForm.Id).Delete(&model.SelectionContent{})
		if db.Error != nil {
			c.JSON(http.StatusOK, apiResponse.ResponseFail(apiResponse.StatusMsgParamsError, apiResponse.StatusCodeParamsError))
			return db.Error
		}
		db = database.GetDatabase().Where("selection_content_id = ?", delForm.Id).Delete(&model.SelectionContentDetail{})
		if db.Error != nil {
			c.JSON(http.StatusOK, apiResponse.ResponseFail(apiResponse.StatusMsgParamsError, apiResponse.StatusCodeParamsError))
			return db.Error
		}
		c.JSON(http.StatusOK, apiResponse.ResponseFail(apiResponse.StatusMsgSuccess, apiResponse.StatusCodeOk))
		return nil
	})
	if err != nil { //创建失败
		c.JSON(http.StatusOK, apiResponse.ResponseFail(apiResponse.StatusMsgParamsError, apiResponse.StatusCodeParamsError))
	}
}
