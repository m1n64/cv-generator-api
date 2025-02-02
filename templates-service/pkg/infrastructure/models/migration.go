package models

type Migration struct {
	Version int `gorm:"primaryKey"`
	Name    string
}
