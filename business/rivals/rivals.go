package rivals

import (
	"context"

	"github.com/bchaillou003/marvel-family-backend/application"
	"github.com/bchaillou003/marvel-family-backend/models/rivals"
)

func CreateAbility(ctx context.Context, ability *rivals.Ability) (*rivals.Ability, error) {
	return application.DatabaseFromContext(ctx).CreateAbility(ctx, ability)
}

func CreateAchievement(ctx context.Context, achievement *rivals.Achievement) (*rivals.Achievement, error) {
	return application.DatabaseFromContext(ctx).CreateAchievement(ctx, achievement)
}

func CreateCharacter(ctx context.Context, character *rivals.Character) (*rivals.Character, error) {
	return application.DatabaseFromContext(ctx).CreateCharacter(ctx, character)
}

func CreateCosmetic(ctx context.Context, cosmetic *rivals.Cosmetic) (*rivals.Cosmetic, error) {
	return application.DatabaseFromContext(ctx).CreateCosmetic(ctx, cosmetic)
}

func CreateLore(ctx context.Context, lore *rivals.Lore) (*rivals.Lore, error) {
	return application.DatabaseFromContext(ctx).CreateLore(ctx, lore)
}

func GetAllAbility(ctx context.Context) ([]rivals.Ability, error) {
	return application.DatabaseFromContext(ctx).GetAllAbility(ctx)
}

func GetAllAchievement(ctx context.Context) ([]rivals.Achievement, error) {
	return application.DatabaseFromContext(ctx).GetAllAchievement(ctx)
}

func GetAllCharacter(ctx context.Context) ([]rivals.Character, error) {
	return application.DatabaseFromContext(ctx).GetAllCharacter(ctx)
}

func GetAllCosmetic(ctx context.Context) ([]rivals.Cosmetic, error) {
	return application.DatabaseFromContext(ctx).GetAllCosmetic(ctx)
}

func GetAllLore(ctx context.Context) ([]rivals.Lore, error) {
	return application.DatabaseFromContext(ctx).GetAllLore(ctx)
}

func GetAbilityByID(ctx context.Context, id int64) (*rivals.Ability, error) {
	return application.DatabaseFromContext(ctx).GetAbilityByID(ctx, id)
}

func GetAchievementByID(ctx context.Context, id int64) (*rivals.Achievement, error) {
	return application.DatabaseFromContext(ctx).GetAchievementByID(ctx, id)
}

func GetCharacterByID(ctx context.Context, id int64) (*rivals.Character, error) {
	return application.DatabaseFromContext(ctx).GetCharacterByID(ctx, id)
}

func GetCosmeticByID(ctx context.Context, id int64) (*rivals.Cosmetic, error) {
	return application.DatabaseFromContext(ctx).GetCosmeticByID(ctx, id)
}

func GetLoreByID(ctx context.Context, id int64) (*rivals.Lore, error) {
	return application.DatabaseFromContext(ctx).GetLoreByID(ctx, id)
}
