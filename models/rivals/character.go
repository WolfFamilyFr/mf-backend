package rivals

import (
	"github.com/uptrace/bun"
)

type Character struct {
	bun.BaseModel `bun:"table:character"`

	ID          int64  `json:"id" bun:",pk,autoincrement"`
	Nickname    string `json:"nickname"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Role        string `json:"role"`
	Difficulty  int    `json:"difficulty"`
	Description string `json:"description"`
}

type CharacterRole string

const (
	Vanguard   CharacterRole = "Vanguard"
	Duelist    CharacterRole = "Duelist"
	Strategist CharacterRole = "Strategist"
)
