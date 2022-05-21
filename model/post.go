package model

import (
	"time"
)

type Post struct {
	ID          uint64    `gorm:"primary_key:auto_increment" json:"id"`
	UrlPhoto    string    `gorm:"type:varchar(255)" json:"urlPhoto"`
	Description string    `gorm:"type:text" json:"description"`
	Likes       uint32    `gorm:"not null;default:0" json:"likes"`
	CreatedAt   time.Time `gorm:"type:DATETIME DEFAULT CURRENT_TIMESTAMP" json:"createdAt"`
	Username    string    `gorm:"type:varchar(50)" json:"username"`
}

const (
	lenShortDescription = 15
)

func (p Post) GetShortDescription() string {

	if len(p.Description) > lenShortDescription {
		s := p.Description[:lenShortDescription+1] + "..."
		return s
	}

	return p.Description
}
