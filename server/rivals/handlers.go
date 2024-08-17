package rivals

import (
	business "github.com/bchaillou003/marvel-family-backend/business/rivals"
	"github.com/bchaillou003/marvel-family-backend/models/rivals"
	"github.com/bchaillou003/marvel-family-backend/utils"
	"github.com/gin-gonic/gin"
)

const (
	paramID = "id"
)

func CreateAbility(c *gin.Context) (*rivals.Ability, error) {
	var value rivals.Ability
	err := c.BindJSON(&value)
	if err != nil {
		return nil, err
	}
	return business.CreateAbility(c.Request.Context(), &value)
}

func CreateAchievement(c *gin.Context) (*rivals.Achievement, error) {
	var value rivals.Achievement
	err := c.BindJSON(&value)
	if err != nil {
		return nil, err
	}
	return business.CreateAchievement(c.Request.Context(), &value)
}

func CreateCharacter(c *gin.Context) (*rivals.Character, error) {
	var value rivals.Character
	err := c.BindJSON(&value)
	if err != nil {
		return nil, err
	}
	return business.CreateCharacter(c.Request.Context(), &value)
}

func CreateCosmetic(c *gin.Context) (*rivals.Cosmetic, error) {
	var value rivals.Cosmetic
	err := c.BindJSON(&value)
	if err != nil {
		return nil, err
	}
	return business.CreateCosmetic(c.Request.Context(), &value)
}

func CreateLore(c *gin.Context) (*rivals.Lore, error) {
	var value rivals.Lore
	err := c.BindJSON(&value)
	if err != nil {
		return nil, err
	}
	return business.CreateLore(c.Request.Context(), &value)
}

func GetAllAbility(c *gin.Context) ([]rivals.Ability, error) {
	return business.GetAllAbility(c.Request.Context())
}

func GetAllAchievement(c *gin.Context) ([]rivals.Achievement, error) {
	return business.GetAllAchievement(c.Request.Context())
}

func GetAllCharacter(c *gin.Context) ([]rivals.Character, error) {
	return business.GetAllCharacter(c.Request.Context())
}

func GetAllCosmetic(c *gin.Context) ([]rivals.Cosmetic, error) {
	return business.GetAllCosmetic(c.Request.Context())
}

func GetAllLore(c *gin.Context) ([]rivals.Lore, error) {
	return business.GetAllLore(c.Request.Context())
}

func GetAbilityByID(c *gin.Context) (*rivals.Ability, error) {
	id, err := utils.StringToInt64(c.Param(paramID))
	if err != nil {
		return nil, err
	}
	return business.GetAbilityByID(c.Request.Context(), id)
}

func GetAchievementByID(c *gin.Context) (*rivals.Achievement, error) {
	id, err := utils.StringToInt64(c.Param(paramID))
	if err != nil {
		return nil, err
	}
	return business.GetAchievementByID(c.Request.Context(), id)
}

func GetCharacterByID(c *gin.Context) (*rivals.Character, error) {
	id, err := utils.StringToInt64(c.Param(paramID))
	if err != nil {
		return nil, err
	}
	return business.GetCharacterByID(c.Request.Context(), id)
}

func GetCosmeticByID(c *gin.Context) (*rivals.Cosmetic, error) {
	id, err := utils.StringToInt64(c.Param(paramID))
	if err != nil {
		return nil, err
	}
	return business.GetCosmeticByID(c.Request.Context(), id)
}

func GetLoreByID(c *gin.Context) (*rivals.Lore, error) {
	id, err := utils.StringToInt64(c.Param(paramID))
	if err != nil {
		return nil, err
	}
	return business.GetLoreByID(c.Request.Context(), id)
}
