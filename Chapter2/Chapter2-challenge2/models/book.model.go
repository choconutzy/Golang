package models

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Title     string    `json:"title" gorm:"type:varchar(255);not null"`
	Author    string    `json:"author" gorm:"type:varchar(255);not null"`
	Descs     string    `json:"descs" gorm:"type:text;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoUpdateTime:false;type:time"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime:true;type:time"`
}
