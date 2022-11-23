package model

import "EmployeeSelection/internal/database"

// 评选活动
type Selection struct {
	BaseModel          `gorm:"embedded"`
	ActivityTitle      string                    `json:"activity_title" gorm:"index;comment:活动标题"`
	ActivityStartTime  int64                     `json:"activity_start_time" gorm:"comment:活动开始时间"`
	ActivityEndTime    int64                     `json:"activity_end_time" gorm:"comment:活动结束时间"`
	Evaluator          []Evaluator               `json:"evaluators" gorm:"many2many:selection_evaluator;comment:评选人活动参与情况"`
	SelectedEmployees  []SelectedEmployees       `json:"selected_employees" gorm:"many2many:selection_selected_employees;comment:被评选员工参与情况"`
	DataStatistics     []SelectionDataStatistics `json:"data_statistics" gorm:"many2many:selection_data_statistics;comment:数据统计"`
	State              int                       `json:"state" gorm:"comment:活动状态"`
	SelectionContentId int64                     `json:"selection_content_id" gorm:"comment:活动内容模板id"`
}

// 评选人状态
type Evaluator struct {
	BaseModel      `gorm:"embedded"`
	UserId         int64 `json:"user_id" gorm:"comment:员工id"`
	SelectionState int   `json:"selection_state" gorm:"comment:评选人状态"`
}

// 被评选员工状态
type SelectedEmployees struct {
	BaseModel          `gorm:"embedded"`
	UserId             int64 `json:"user_id" gorm:"comment:员工id"`
	UserSelectionState int   `json:"user_selection_state" gorm:"comment:用户参与情况"`
}

type SelectionDataStatistics struct {
	BaseModel  `gorm:"embedded"`
	UserId     int64            `json:"user_id" gorm:"comment:员工id"`
	Score      []SelectionScore `json:"score" gorm:"many2many:selection_selected_employees_score;comment:成绩"`
	TotalScore int              `json:"total_score" gorm:"comment:总成绩"`
}

type SelectionScore struct {
	BaseModel                `gorm:"embedded"`
	SelectionContentDetailId int64 `json:"selection_content_detail_id" gorm:"comment:评选内容id"`
	Score                    int   `json:"score" gorm:"comment:成绩得分"`
}

func init() {
	database.GetDatabase().AutoMigrate(&Selection{})
}
