package rivals

import "gorm.io/gorm"

type Character struct {
	gorm.Model
	ID          uint   `json:"id"`
	Nickname    string `json:"nickname"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Role        string `json:"role"`
	Difficulty  int    `json:"difficulty"`
	Description string `json:"description"`
	// Cosmetics   []Cosmetic    `gorm:"foreignKey:ID"`
	// Abilities   []Ability     `gorm:"foreignKey:ID"`
	// Lore        []Lore        `gorm:"foreignKey:ID"`
	// Achievement []Achievement `gorm:"foreignKey:ID"`
}

type CharacterRole string

const (
	Vanguard   CharacterRole = "Vanguard"
	Duelist    CharacterRole = "Duelist"
	Strategist CharacterRole = "Strategist"
)
