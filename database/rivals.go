package database

import (
	"context"

	"github.com/bchaillou003/marvel-family-backend/models/rivals"
)

type IRivals interface {
	CreateAbility(ctx context.Context, ability *rivals.Ability) (*rivals.Ability, error)
	CreateAchievement(ctx context.Context, achievement *rivals.Achievement) (*rivals.Achievement, error)
	CreateCharacter(ctx context.Context, character *rivals.Character) (*rivals.Character, error)
	CreateCosmetic(ctx context.Context, cosmetic *rivals.Cosmetic) (*rivals.Cosmetic, error)
	CreateLore(ctx context.Context, lore *rivals.Lore) (*rivals.Lore, error)

	GetAllAbility(ctx context.Context) ([]rivals.Ability, error)
	GetAllAchievement(ctx context.Context) ([]rivals.Achievement, error)
	GetAllCharacter(ctx context.Context) ([]rivals.Character, error)
	GetAllCosmetic(ctx context.Context) ([]rivals.Cosmetic, error)
	GetAllLore(ctx context.Context) ([]rivals.Lore, error)

	GetAbilityByID(ctx context.Context, id int64) (*rivals.Ability, error)
	GetAchievementByID(ctx context.Context, id int64) (*rivals.Achievement, error)
	GetCharacterByID(ctx context.Context, id int64) (*rivals.Character, error)
	GetCosmeticByID(ctx context.Context, id int64) (*rivals.Cosmetic, error)
	GetLoreByID(ctx context.Context, id int64) (*rivals.Lore, error)
}

func (db Database) CreateAbility(ctx context.Context, ability *rivals.Ability) (*rivals.Ability, error) {
	_, err := db.NewInsert().Model(ability).Returning("id").Exec(ctx)
	return ability, err
}

func (db Database) CreateAchievement(ctx context.Context, achievement *rivals.Achievement) (*rivals.Achievement, error) {
	_, err := db.NewInsert().Model(achievement).Returning("id").Exec(ctx)
	return achievement, err
}

func (db Database) CreateCharacter(ctx context.Context, character *rivals.Character) (*rivals.Character, error) {
	_, err := db.NewInsert().Model(character).Returning("id").Exec(ctx)
	return character, err
}

func (db Database) CreateCosmetic(ctx context.Context, cosmetic *rivals.Cosmetic) (*rivals.Cosmetic, error) {
	_, err := db.NewInsert().Model(cosmetic).Returning("id").Exec(ctx)
	return cosmetic, err
}

func (db Database) CreateLore(ctx context.Context, lore *rivals.Lore) (*rivals.Lore, error) {
	_, err := db.NewInsert().Model(lore).Returning("id").Exec(ctx)
	return lore, err
}

func (db Database) GetAllAbility(ctx context.Context) ([]rivals.Ability, error) {
	records := []rivals.Ability{}
	err := db.NewSelect().Model(&records).Scan(ctx)
	return records, err
}

func (db Database) GetAllAchievement(ctx context.Context) ([]rivals.Achievement, error) {
	records := []rivals.Achievement{}
	err := db.NewSelect().Model(&records).Scan(ctx)
	return records, err
}

func (db Database) GetAllCharacter(ctx context.Context) ([]rivals.Character, error) {
	records := []rivals.Character{}
	err := db.NewSelect().Model(&records).Scan(ctx)
	return records, err
}

func (db Database) GetAllCosmetic(ctx context.Context) ([]rivals.Cosmetic, error) {
	records := []rivals.Cosmetic{}
	err := db.NewSelect().Model(&records).Scan(ctx)
	return records, err
}

func (db Database) GetAllLore(ctx context.Context) ([]rivals.Lore, error) {
	records := []rivals.Lore{}
	err := db.NewSelect().Model(&records).Scan(ctx)
	return records, err
}

func (db Database) GetAbilityByID(ctx context.Context, id int64) (*rivals.Ability, error) {
	var record rivals.Ability
	err := db.NewSelect().Model(&record).Where("id = ?", id).Scan(ctx)
	return &record, err
}

func (db Database) GetAchievementByID(ctx context.Context, id int64) (*rivals.Achievement, error) {
	var record rivals.Achievement
	err := db.NewSelect().Model(&record).Where("id = ?", id).Scan(ctx)
	return &record, err
}

func (db Database) GetCharacterByID(ctx context.Context, id int64) (*rivals.Character, error) {
	var record rivals.Character
	err := db.NewSelect().Model(&record).Where("id = ?", id).Scan(ctx)
	return &record, err
}

func (db Database) GetCosmeticByID(ctx context.Context, id int64) (*rivals.Cosmetic, error) {
	var record rivals.Cosmetic
	err := db.NewSelect().Model(&record).Where("id = ?", id).Scan(ctx)
	return &record, err
}

func (db Database) GetLoreByID(ctx context.Context, id int64) (*rivals.Lore, error) {
	var record rivals.Lore
	err := db.NewSelect().Model(&record).Where("id = ?", id).Scan(ctx)
	return &record, err
}
