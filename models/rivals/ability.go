package rivals

import "github.com/uptrace/bun"

type Ability struct {
	bun.BaseModel `bun:"table:ability"`

	ID          int64       `json:"id" bun:",pk,autoincrement"`
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
