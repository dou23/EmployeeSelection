package selection

import (
	apiResponse "EmployeeSelection/internal/api/response"
	"EmployeeSelection/internal/database"
	"EmployeeSelection/internal/database/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type selectionCreateForm struct {
	ActivityTitle      string  `json:"activity_title"`      //活动标题
	ActivityStartTime  int64   `json:"activity_start_time"` //活动开始时间
	ActivityEndTime    int64   `json:"activity_end_time" `  //活动结束时间
	Evaluator          []int64 `json:"evaluators" `         //活动参与情况
	SelectedEmployees  []int64 `json:"selected_employees" ` //活动参与情况
	SelectionContentId int64   `json:"selection_content_id" gorm:"comment:活动内容模板id"`
}

func SelectionCreate(c *gin.Context) {
	selectionForm := &selectionCreateForm{}
	c.ShouldBind(selectionForm)
	//活动标题
	if selectionForm.ActivityTitle == "" {
		c.JSON(http.StatusOK, apiResponse.ResponseFail("请传入标题", apiResponse.StatusCodeParamsError))
		return
	}
	//开始时间
	if selectionForm.ActivityStartTime == 0 {
		c.JSON(http.StatusOK, apiResponse.ResponseFail("请传入活动开始时间", apiResponse.StatusCodeParamsError))
		return
	}
	//结束时间
	if selectionForm.ActivityEndTime == 0 {
		c.JSON(http.StatusOK, apiResponse.ResponseFail("请传入活动结束时间", apiResponse.StatusCodeParamsError))
		return
	}
	//被评选人
	if selectionForm.Evaluator == nil || len(selectionForm.Evaluator) == 0 {
		c.JSON(http.StatusOK, apiResponse.ResponseFail("请选择评选人", apiResponse.StatusCodeParamsError))
		return
	}
	//评选人呢
	if selectionForm.SelectedEmployees == nil || len(selectionForm.SelectedEmployees) == 0 {
		c.JSON(http.StatusOK, apiResponse.ResponseFail("请选择被评选人", apiResponse.StatusCodeParamsError))
		return
	}
	//评选内容模板
	if selectionForm.SelectionContentId == 0 {
		c.JSON(http.StatusOK, apiResponse.ResponseFail("请选择评选内容模板", apiResponse.StatusCodeParamsError))
		return
	}
	selection := &model.Selection{
		ActivityTitle:      selectionForm.ActivityTitle,
		ActivityStartTime:  selectionForm.ActivityStartTime,
		ActivityEndTime:    selectionForm.ActivityEndTime,
		SelectionContentId: selectionForm.SelectionContentId,
	}
	selection.State = 0
	selection.SelectedEmployees = make([]model.SelectedEmployees, len(selectionForm.SelectedEmployees))
	for i := range selectionForm.SelectedEmployees {
		selection.SelectedEmployees[i].UserId = selectionForm.SelectedEmployees[i]
		selection.SelectedEmployees[i].UserSelectionState = 0
	}
	selection.Evaluator = make([]model.Evaluator, len(selectionForm.Evaluator))
	for i := range selectionForm.Evaluator {
		selection.Evaluator[i].UserId = selectionForm.Evaluator[i]
		selection.Evaluator[i].SelectionState = 0
	}
	selection.GenerateID()
	tx := database.GetDatabase().Create(selection)
	if tx.Error == nil {
		c.JSON(http.StatusOK, apiResponse.ResponseFail(tx.Error.Error(), apiResponse.StatusCodeParamsError))
		return
	}
	c.JSON(http.StatusOK, apiResponse.ResponseOk("创建成功"))
}
