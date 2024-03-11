package models

import (
	"gorm.io/gorm"
)

type Database struct {
	Host     string
	Pass     string
	Name     string
	Port     int
	User     string
	Psqlinfo string
	Db       *gorm.DB
}

type User struct {
	Id       string `gorm:"primaryKey" json:"Id"`
	Name     string `json:"Name"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

type Article struct {
	Id     string  `gorm:"primaryKey" json:"Id"`
	Name   string  `json:"Name"`
	Size   string  `json:"Size"`
	Price  float32 `json:"Price"`
	State  string  `json:"State"`
	Brand  string  `json:"Brand"`
	Color  string  `json:"Color"`
	UserId string  `json:"UserId"`
}
