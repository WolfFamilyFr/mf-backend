package rivals

import "github.com/uptrace/bun"

type Achievement struct {
	bun.BaseModel `bun:"table:achievement"`

	ID          int64           `json:"id" bun:",pk,autoincrement"`
	Name        string          `json:"name"`
	Type        AchievementType `json:"type"`
	Description string          `json:"description"`
	NbTask      int64           `json:"nb_task"`
	CosmeticID  int64           `json:"cosmetic_id"`
}

type AchievementType string

const (
	Legend   AchievementType = "Legend"
	Bootcamp AchievementType = "Bootcamp"
)
