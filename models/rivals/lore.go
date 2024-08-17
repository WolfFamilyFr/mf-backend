package rivals

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Lore struct {
	gorm.Model
	ID         uint           `json:"id"`
	Name       string         `json:"name"`
	Paragraphs pq.StringArray `json:"paragraphs" gorm:"type:text[]"`
}
