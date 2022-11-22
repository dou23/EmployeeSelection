package model

import "EmployeeSelection/internal/database"

type SelectionContent struct {
	BaseModel     `gorm:"embedded"`
	Name          string                   `json:"name" gorm:"comment:评选活动名称"`
	ContentDetail []SelectionContentDetail `json:"contentDetail" gorm:"-"`
}

type SelectionContentDetail struct {
	BaseModel          `gorm:"embedded"`
	SelectionContentId int64  `json:"selection_content_id" gorm:"primaryKey;index;comment:评选活动id"`
	Name               string `json:"name" gorm:"comment:评选内容"`
	Description        string `json:"description" gorm:"comment:评选说明备注"`
	HighestScore       int    `json:"highest_score" gorm:"comment:最高分"`
	LowestScore        int    `json:"lowest_score" gorm:"comment:最低分"`
}

func init() {
	err := database.GetDatabase().AutoMigrate(&SelectionContent{}, &SelectionContentDetail{})
	if err != nil {
		panic(err)
	}
}
