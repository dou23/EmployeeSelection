package selection

import (
	apiResponse "EmployeeSelection/internal/api/response"
	"EmployeeSelection/internal/database"
	"EmployeeSelection/internal/database/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type selectionEditForm struct {
	Id                 int64   `json:"id"`                  // 活动id
	ActivityTitle      string  `json:"activity_title"`      //活动标题
	ActivityStartTime  int64   `json:"activity_start_time"` //活动开始时间
	ActivityEndTime    int64   `json:"activity_end_time" `  //活动结束时间
	Evaluator          []int64 `json:"evaluators" `         //活动参与情况
	SelectedEmployees  []int64 `json:"selected_employees" ` //活动参与情况
	SelectionContentId int64   `json:"selection_content_id" gorm:"comment:活动内容模板id"`
}

// / 评选活动编辑
func SelectionEdit(c *gin.Context) {
	editForm := &selectionEditForm{}
	c.ShouldBind(editForm)
	if editForm.Id == 0 {
		c.JSON(http.StatusOK, apiResponse.ResponseFail("参数异常，当前活动不存在", apiResponse.StatusCodeParamsError))
		return
	}
	//活动标题
	if editForm.ActivityTitle == "" {
		c.JSON(http.StatusOK, apiResponse.ResponseFail("请传入标题", apiResponse.StatusCodeParamsError))
		return
	}
	//开始时间
	if editForm.ActivityStartTime == 0 {
		c.JSON(http.StatusOK, apiResponse.ResponseFail("请传入活动开始时间", apiResponse.StatusCodeParamsError))
		return
	}
	//结束时间
	if editForm.ActivityEndTime == 0 {
		c.JSON(http.StatusOK, apiResponse.ResponseFail("请传入活动结束时间", apiResponse.StatusCodeParamsError))
		return
	}
	//被评选人
	if editForm.Evaluator == nil || len(editForm.Evaluator) == 0 {
		c.JSON(http.StatusOK, apiResponse.ResponseFail("请选择评选人", apiResponse.StatusCodeParamsError))
		return
	}
	//评选人呢
	if editForm.SelectedEmployees == nil || len(editForm.SelectedEmployees) == 0 {
		c.JSON(http.StatusOK, apiResponse.ResponseFail("请选择被评选人", apiResponse.StatusCodeParamsError))
		return
	}
	//评选内容模板
	if editForm.SelectionContentId == 0 {
		c.JSON(http.StatusOK, apiResponse.ResponseFail("请选择评选内容模板", apiResponse.StatusCodeParamsError))
		return
	}
	selection := &model.Selection{
		ActivityTitle:      editForm.ActivityTitle,
		ActivityStartTime:  editForm.ActivityStartTime,
		ActivityEndTime:    editForm.ActivityEndTime,
		SelectionContentId: editForm.SelectionContentId,
	}
	selection.ID = editForm.Id
	selection.State = 0
	selection.SelectedEmployees = make([]model.SelectedEmployees, len(editForm.SelectedEmployees))
	for i := range editForm.SelectedEmployees {
		selection.SelectedEmployees[i].UserId = editForm.SelectedEmployees[i]
	}
	selection.Evaluator = make([]model.Evaluator, len(editForm.Evaluator))
	if editForm.Evaluator != nil && len(editForm.Evaluator) != 0 {
		for i := range editForm.Evaluator {
			selection.Evaluator[i].UserId = editForm.Evaluator[i]
		}
	}
	tx := database.GetDatabase().Model(selection).Updates(selection)
	if tx.Error == nil {
		c.JSON(http.StatusOK, apiResponse.ResponseFail(tx.Error.Error(), apiResponse.StatusCodeParamsError))
		return
	}
	c.JSON(http.StatusOK, apiResponse.ResponseOk("编辑成功"))
}
