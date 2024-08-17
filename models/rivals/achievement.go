package rivals

import "gorm.io/gorm"

type Achievement struct {
	gorm.Model
	ID          uint            `json:"id"`
	Name        string          `json:"name"`
	Type        AchievementType `json:"type"`
	Description string          `json:"description"`
	NbTask      uint            `json:"nb_task"`
	Cosmetic    Cosmetic        `gorm:"foreignKey:ID"`
}

type AchievementType string

const (
	Legend   AchievementType = "Legend"
	Bootcamp AchievementType = "Bootcamp"
)
