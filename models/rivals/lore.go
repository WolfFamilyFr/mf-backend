package rivals

import (
	"github.com/uptrace/bun"
)

type Lore struct {
	bun.BaseModel `bun:"table:lore"`

	ID        int64  `json:"id" bun:",pk,autoincrement"`
	Name      string `json:"name"`
	Paragraph string `json:"paragraph"`
}
