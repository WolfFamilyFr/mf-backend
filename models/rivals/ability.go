package rivals

import "gorm.io/gorm"

type Ability struct {
	gorm.Model
	ID          uint        `json:"id"`
	Name        string      `json:"name"`
	Type        AbilityType `json:"type"`
	Description string      `json:"description"`
	Icon        string      `json:"icon"`
	Video       string      `json:"video"`
	ShortcutKey int         `json:"shortcut_key"`
}

type AbilityType string

const (
	NormalAttack    AbilityType = "Normal Attack"
	Abilities       AbilityType = "Abilities"
	TeamUpAbilities AbilityType = "Team-Up Abilities"
)
