package selectionContent

import (
	apiResponse "EmployeeSelection/internal/api/response"
	"EmployeeSelection/internal/database"
	"EmployeeSelection/internal/database/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type SelectionContentEditForm struct {
	Name    string                           `json:"name"`    //评选活动模板名称
	Id      int64                            `json:"id"`      //评选活动模板id
	Details []SelectionContentDetailEditForm `json:"details"` //评选详情
}

type SelectionContentDetailEditForm struct {
	Id           int64  `json:"id"`            //评选详情id
	Name         string `json:"name"`          //评选内容
	Description  string `json:"description"`   //评选备注
	HighestScore int    `json:"highest_score"` //最高得分
	LowestScore  int    `json:"lowest_score"`  //最低得分
}

// SelectionContentEdit 更新评选内容模板
func SelectionContentEdit(c *gin.Context) {
	editForm := &SelectionContentEditForm{}
	c.ShouldBind(editForm)
	if editForm.Id == 0 || editForm.Details == nil || len(editForm.Details) == 0 || editForm.Name == "" {
		c.JSON(http.StatusOK, apiResponse.ResponseFail(apiResponse.StatusMsgParamsError, apiResponse.StatusCodeParamsError))
		return
	}
	selectionContent := &model.SelectionContent{
		Name: editForm.Name,
	}
	selectionContent.ID = editForm.Id
	details := make([]model.SelectionContentDetail, len(editForm.Details))
	err := database.GetDatabase().Transaction(func(tx *gorm.DB) error {
		tx = database.GetDatabase().Model(selectionContent).Updates(selectionContent)
		if tx.Error != nil {
			c.JSON(http.StatusOK, apiResponse.ResponseFail(tx.Error.Error(), apiResponse.StatusCodeParamsError))
			return tx.Error
		}
		for i := range editForm.Details {
			details[i].SelectionContentId = editForm.Id
			details[i].Name = editForm.Details[i].Name
			details[i].Description = editForm.Details[i].Description
			details[i].LowestScore = editForm.Details[i].LowestScore
			details[i].HighestScore = editForm.Details[i].HighestScore
			details[i].ID = editForm.Details[i].Id
			tx = database.GetDatabase().Model(details[i]).Updates(details[i])
			if tx.Error != nil {
				c.JSON(http.StatusOK, apiResponse.ResponseFail(tx.Error.Error(), apiResponse.StatusCodeParamsError))
				return tx.Error
			}
		}
		selectionContent.ContentDetail = details
		return nil
	})
	if err != nil {
		c.JSON(http.StatusOK, apiResponse.ResponseFail(err.Error(), apiResponse.StatusCodeParamsError))
		return
	}
	c.JSON(http.StatusOK, apiResponse.ResponseOk(apiResponse.StatusMsgSuccess))
}
