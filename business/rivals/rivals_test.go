package rivals

import (
	"context"
	"testing"

	"github.com/bchaillou003/marvel-family-backend/application"
	"github.com/bchaillou003/marvel-family-backend/application/mocks"
	models "github.com/bchaillou003/marvel-family-backend/models/rivals"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_CreateAbility(t *testing.T) {
	// Setup
	ctx := context.Background()
	app, dbm, _ := mocks.NewMockedApplication()
	ctx = application.ToContext(ctx, app)

	ability := models.Ability{
		ID: 1,
	}

	dbm.On("CreateAbility", mock.Anything, mock.Anything).Return(&ability, nil)

	// Exercise
	result, err := CreateAbility(ctx, &ability)

	// Verify
	require.NoError(t, err)
	require.Equal(t, ability.ID, result.ID)
}

func Test_CreateAchievement(t *testing.T) {
	// Setup
	ctx := context.Background()
	app, dbm, _ := mocks.NewMockedApplication()
	ctx = application.ToContext(ctx, app)

	achievement := models.Achievement{
		ID: 1,
	}

	dbm.On("CreateAchievement", mock.Anything, mock.Anything).Return(&achievement, nil)

	// Exercise
	result, err := CreateAchievement(ctx, &achievement)

	// Verify
	require.NoError(t, err)
	require.Equal(t, achievement.ID, result.ID)
}

func Test_CreateCharacter(t *testing.T) {
	// Setup
	ctx := context.Background()
	app, dbm, _ := mocks.NewMockedApplication()
	ctx = application.ToContext(ctx, app)

	character := models.Character{
		ID: 1,
	}

	dbm.On("CreateCharacter", mock.Anything, mock.Anything).Return(&character, nil)

	// Exercise
	result, err := CreateCharacter(ctx, &character)

	// Verify
	require.NoError(t, err)
	require.Equal(t, character.ID, result.ID)
}

func Test_CreateCosmetic(t *testing.T) {
	// Setup
	ctx := context.Background()
	app, dbm, _ := mocks.NewMockedApplication()
	ctx = application.ToContext(ctx, app)

	cosmetic := models.Cosmetic{
		ID: 1,
	}

	dbm.On("CreateCosmetic", mock.Anything, mock.Anything).Return(&cosmetic, nil)

	// Exercise
	result, err := CreateCosmetic(ctx, &cosmetic)

	// Verify
	require.NoError(t, err)
	require.Equal(t, cosmetic.ID, result.ID)
}

func Test_CreateLore(t *testing.T) {
	// Setup
	ctx := context.Background()
	app, dbm, _ := mocks.NewMockedApplication()
	ctx = application.ToContext(ctx, app)

	lore := models.Lore{
		ID: 1,
	}

	dbm.On("CreateLore", mock.Anything, mock.Anything).Return(&lore, nil)

	// Exercise
	result, err := CreateLore(ctx, &lore)

	// Verify
	require.NoError(t, err)
	require.Equal(t, lore.ID, result.ID)
}

func Test_GetAllAbility(t *testing.T) {
	// Setup
	ctx := context.Background()
	app, dbm, _ := mocks.NewMockedApplication()
	ctx = application.ToContext(ctx, app)

	ability := models.Ability{
		ID: 1,
	}

	dbm.On("GetAllAbility", mock.Anything).Return([]models.Ability{ability}, nil)

	// Exercise
	result, err := GetAllAbility(ctx)

	// Verify
	require.NoError(t, err)
	require.Equal(t, 1, len(result))
	require.Equal(t, ability.ID, result[0].ID)
}

func Test_GetAllAchievement(t *testing.T) {
	// Setup
	ctx := context.Background()
	app, dbm, _ := mocks.NewMockedApplication()
	ctx = application.ToContext(ctx, app)

	achievement := models.Achievement{
		ID: 1,
	}

	dbm.On("GetAllAchievement", mock.Anything).Return([]models.Achievement{achievement}, nil)

	// Exercise
	result, err := GetAllAchievement(ctx)

	// Verify
	require.NoError(t, err)
	require.Equal(t, 1, len(result))
	require.Equal(t, achievement.ID, result[0].ID)
}

func Test_GetAllCharacter(t *testing.T) {
	// Setup
	ctx := context.Background()
	app, dbm, _ := mocks.NewMockedApplication()
	ctx = application.ToContext(ctx, app)

	character := models.Character{
		ID: 1,
	}

	dbm.On("GetAllCharacter", mock.Anything).Return([]models.Character{character}, nil)

	// Exercise
	result, err := GetAllCharacter(ctx)

	// Verify
	require.NoError(t, err)
	require.Equal(t, 1, len(result))
	require.Equal(t, character.ID, result[0].ID)
}

func Test_GetAllCosmetic(t *testing.T) {
	// Setup
	ctx := context.Background()
	app, dbm, _ := mocks.NewMockedApplication()
	ctx = application.ToContext(ctx, app)

	cosmetic := models.Cosmetic{
		ID: 1,
	}

	dbm.On("GetAllCosmetic", mock.Anything).Return([]models.Cosmetic{cosmetic}, nil)

	// Exercise
	result, err := GetAllCosmetic(ctx)

	// Verify
	require.NoError(t, err)
	require.Equal(t, 1, len(result))
	require.Equal(t, cosmetic.ID, result[0].ID)
}

func Test_GetAllLore(t *testing.T) {
	// Setup
	ctx := context.Background()
	app, dbm, _ := mocks.NewMockedApplication()
	ctx = application.ToContext(ctx, app)

	lore := models.Lore{
		ID: 1,
	}

	dbm.On("GetAllLore", mock.Anything).Return([]models.Lore{lore}, nil)

	// Exercise
	result, err := GetAllLore(ctx)

	// Verify
	require.NoError(t, err)
	require.Equal(t, 1, len(result))
	require.Equal(t, lore.ID, result[0].ID)
}

func Test_GetAbilityByID(t *testing.T) {
	// Setup
	ctx := context.Background()
	app, dbm, _ := mocks.NewMockedApplication()
	ctx = application.ToContext(ctx, app)

	ability := models.Ability{
		ID: 1,
	}

	dbm.On("GetAbilityByID", mock.Anything, mock.Anything).Return(&ability, nil)

	// Exercise
	result, err := GetAbilityByID(ctx, 1)

	// Verify
	require.NoError(t, err)
	require.Equal(t, ability.ID, result.ID)
}

func Test_GetAchievementByID(t *testing.T) {
	// Setup
	ctx := context.Background()
	app, dbm, _ := mocks.NewMockedApplication()
	ctx = application.ToContext(ctx, app)

	achievement := models.Achievement{
		ID: 1,
	}

	dbm.On("GetAchievementByID", mock.Anything, mock.Anything).Return(&achievement, nil)

	// Exercise
	result, err := GetAchievementByID(ctx, 1)

	// Verify
	require.NoError(t, err)
	require.Equal(t, achievement.ID, result.ID)
}

func Test_GetCharacterByID(t *testing.T) {
	// Setup
	ctx := context.Background()
	app, dbm, _ := mocks.NewMockedApplication()
	ctx = application.ToContext(ctx, app)

	character := models.Character{
		ID: 1,
	}

	dbm.On("GetCharacterByID", mock.Anything, mock.Anything).Return(&character, nil)

	// Exercise
	result, err := GetCharacterByID(ctx, 1)

	// Verify
	require.NoError(t, err)
	require.Equal(t, character.ID, result.ID)
}

func Test_GetCosmeticByID(t *testing.T) {
	// Setup
	ctx := context.Background()
	app, dbm, _ := mocks.NewMockedApplication()
	ctx = application.ToContext(ctx, app)

	cosmetic := models.Cosmetic{
		ID: 1,
	}

	dbm.On("GetCosmeticByID", mock.Anything, mock.Anything).Return(&cosmetic, nil)

	// Exercise
	result, err := GetCosmeticByID(ctx, 1)

	// Verify
	require.NoError(t, err)
	require.Equal(t, cosmetic.ID, result.ID)
}

func Test_GetLoreByID(t *testing.T) {
	// Setup
	ctx := context.Background()
	app, dbm, _ := mocks.NewMockedApplication()
	ctx = application.ToContext(ctx, app)

	lore := models.Lore{
		ID: 1,
	}

	dbm.On("GetLoreByID", mock.Anything, mock.Anything).Return(&lore, nil)

	// Exercise
	result, err := GetLoreByID(ctx, 1)

	// Verify
	require.NoError(t, err)
	require.Equal(t, lore.ID, result.ID)
}
