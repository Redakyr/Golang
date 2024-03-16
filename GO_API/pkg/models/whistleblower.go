package models

import (
	"github.com/Redakyr/Golang/GO_API/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Agent struct {
	gorm.Model
	Name       string `gorm:""json:"name"`
	Birth      string `gorm:""json:"birth"`
	Imprisoned bool   `gorm:""json:"imprisoned"`
}

func init() {
	config.Connection()
	db = config.GetDB()
	db.AutoMigrate(&Agent{})
}

func (b *Agent) AddAgent() *Agent {
	db.NewRecord(b)
	db.Create(&b)
	return b
}