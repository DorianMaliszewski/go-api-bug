package models

import (
	"encoding/json"
	"io"
	"time"

	"gorm.io/gorm"
)

//Bug New struct
type Bug struct {
	ID          int32          `json:"id" gorm:"autoIncrement;primaryKey"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	CreatedBy   string         `json:"createdBy"`
	Status      string         `json:"status"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

func (Bug) TableName() string {
	return "bugs"
}

func (b *Bug) ToJSON(writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	return encoder.Encode(b)
}

func (b *Bug) FromJSON(reader io.Reader) error {
	decoder := json.NewDecoder(reader)
	return decoder.Decode(b)
}
