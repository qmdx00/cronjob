package rpc

import (
	"gorm.io/gorm"
	"time"
)

// gormCreate ...
func gormCreate(db *gorm.DB, now time.Time) error {
	at := now.Format("2006-01-02 15:04:05")
	db.Statement.SetColumn("created_at", at)
	db.Statement.SetColumn("updated_at", at)
	return nil
}

// TableName ...
func (x *Task_Model) TableName() string {
	return "task"
}

// BeforeCreate ...
func (x *Task_Model) BeforeCreate(scope *gorm.DB) error {
	return gormCreate(scope, time.Now().Local())
}
