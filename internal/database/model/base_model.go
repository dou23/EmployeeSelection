package model

import (
	"EmployeeSelection/internal/config"
	"EmployeeSelection/internal/database"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"log"
	"time"
)

var (
	snode *snowflake.Node
)

type IModel interface {
	GenerateID() int64
	ParseID() snowflake.ID
	GetID() int64
}

type BaseModel struct {
	// TODO 分布式ID 雪花算法 https://www.itqiankun.com/article/1565747019
	ID        int64     `gorm:"<-:create;primarykey"`
	CreatedAt time.Time `gorm:"comment:创建时间"`
	UpdatedAt time.Time `gorm:"comment:更新时间"`
}

func (b *BaseModel) GenerateID() int64 {
	if b.ID == 0 {
		id := getSnowflakeNode().Generate().Int64()
		if id == 0 {
			panic("Error: getSnowflakeNode().Generate().Int64() == 0")
		}
		b.ID = id
	}
	return b.ID
}

func (b BaseModel) ParseID() snowflake.ID {
	return snowflake.ParseInt64(b.ID)
}

func (b BaseModel) GetID() int64 {
	return b.ID
}

func getSnowflakeNode() *snowflake.Node {
	if snode == nil {
		node, err := snowflake.NewNode(getNodeId())
		if err != nil {
			fmt.Println(err)
			snode = nil
		}
		snode = node
	}
	log.Println("---getSnowflakeNode---", snode)
	return snode
}

func getNodeId() int64 {
	return int64(config.GetNodeId())
}

func CreateModel(m IModel) (int64, error) {
	id := m.GenerateID()
	tx := database.GetDatabase().Create(m)
	err := tx.Error
	return id, err
}

func CreateModels(m []IModel) error {
	for i := range m {
		m[i].GenerateID()
	}
	tx := database.GetDatabase().Create(m)
	err := tx.Error
	return err
}

func GetModel(m IModel) {
	database.GetDatabase().Find(m)
}
