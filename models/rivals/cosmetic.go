package rivals

import "github.com/uptrace/bun"

type Cosmetic struct {
	bun.BaseModel `bun:"table:cosmetic"`

	ID      int64        `json:"id" bun:",pk,autoincrement"`
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
