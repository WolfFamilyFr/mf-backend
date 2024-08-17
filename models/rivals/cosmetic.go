package rivals

import "gorm.io/gorm"

type Cosmetic struct {
	gorm.Model
	ID      uint         `json:"id"`
	Name    string       `json:"name"`
	Type    CosmeticType `json:"type"`
	Default bool         `json:"default"`
}

type CosmeticType string

const (
	Costume CosmeticType = "Costume"
	MVP     CosmeticType = "MVP"
	Emote   CosmeticType = "Emote"
	Spray   CosmeticType = "Spray"
)
