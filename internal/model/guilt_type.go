package model

type GuiltType struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Name      string `json:"name" binding:"required"`
	OtherInfo string `json:"otherInfo,omitempty"`
}