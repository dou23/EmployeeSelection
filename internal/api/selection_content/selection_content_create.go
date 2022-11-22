package selectionContent

import (
	apiResponse "EmployeeSelection/internal/api/response"
	"EmployeeSelection/internal/database"
	"EmployeeSelection/internal/database/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type SelectionContentForm struct {
	Name    string                       `json:"name"`    //评选活动模板名称
	Details []SelectionContentDetailForm `json:"details"` //评选内容数据列表
}

type SelectionContentDetailForm struct {
	Name         string `json:"name"`          //评选内容
	Description  string `json:"description"`   //评选备注
	HighestScore int    `json:"highest_score"` //最高得分
	LowestScore  int    `json:"lowest_score"`  //最低得分
}

func SelectionContentCreate(c *gin.Context) {
	selectionContentForm := &SelectionContentForm{}
	c.ShouldBind(selectionContentForm)
	if selectionContentForm.Name == "" {
		c.JSON(http.StatusOK, apiResponse.ResponseFail(apiResponse.StatusMsgParamsError, apiResponse.StatusCodeParamsError))
		return
	}
	if selectionContentForm.Details == nil || len(selectionContentForm.Details) == 0 {
		c.JSON(http.StatusOK, apiResponse.ResponseFail(apiResponse.StatusMsgParamsError, apiResponse.StatusCodeParamsError))
		return
	}
	selectionContent := &model.SelectionContent{
		Name: selectionContentForm.Name,
	}
	err := database.GetDatabase().Transaction(func(tx *gorm.DB) error {
		_, err := model.CreateModel(selectionContent)
		if err != nil { //创建失败
			c.JSON(http.StatusOK, apiResponse.ResponseFail(apiResponse.StatusMsgParamsError, apiResponse.StatusCodeParamsError))
			return err
		}
		selectionContentDetails := make([]*model.SelectionContentDetail, len(selectionContentForm.Details))
		for i := range selectionContentForm.Details {
			detailForm := selectionContentForm.Details[i]
			selectionContentDetails[i] = &model.SelectionContentDetail{
				SelectionContentId: selectionContent.ID,
				Name:               detailForm.Name,
				Description:        detailForm.Description,
				HighestScore:       detailForm.HighestScore,
				LowestScore:        detailForm.LowestScore,
			}
			_, err := model.CreateModel(selectionContentDetails[i])
			if err != nil {
				return err
			}
		}
		c.JSON(http.StatusOK, apiResponse.ResponseOk(apiResponse.StatusMsgSuccess))
		return nil
	})
	if err != nil { //创建失败
		c.JSON(http.StatusOK, apiResponse.ResponseFail(apiResponse.StatusMsgParamsError, apiResponse.StatusCodeParamsError))
	}
}
