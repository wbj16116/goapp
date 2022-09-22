package models

import (
	// "time"
)

type Guser struct {
	Id   int64
	Name string `json:"name"`
	Age int64
	CreatedAt int64 `gorm:"autoCreateTime"`
    UpdatedAt int64 `gorm:"autoCreateTime"`
}

//自定义操作的数据表
func (Guser) TableName() string {
	return "test"
}
