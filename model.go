package main

import "time"

type RecordLogs struct {
	Id        int    `gorm:"primary_key;AUTO_INCREMENT" form:"id" json:"id"`
	Type      string `gorm:"not null" form:"type" query:"type" json:"type"`
	Model     string `gorm:"not null" form:"model" query:"model" json:"model"`
	Content   string `gorm:"not null" form:"content" query:"content"  json:"content"`
	CreatedAt time.Time
}

type AuthKey struct {
	Id  int    `gorm:"primary_key;AUTO_INCREMENT" form:"id" json:"id"`
	Key string `form:"key" json:"key"`
}
