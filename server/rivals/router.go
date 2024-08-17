package rivals

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	group := router.Group("/rivals")
	{
		setupAbilityRoutes(group)
		setupAchievementRoutes(group)
		setupCharacterRoutes(group)
		setupCosmeticRoutes(group)
		setupLoreRoutes(group)
	}
}

func setupAbilityRoutes(group *gin.RouterGroup) {
	abilityGroup := group.Group("/ability")
	{
		abilityGroup.POST("", func(c *gin.Context) {
			result, err := CreateAbility(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": err})
				return
			}
			c.JSON(http.StatusOK, result)
		})
		abilityGroup.GET("", func(c *gin.Context) {
			result, err := GetAllAbility(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": err})
				return
			}
			c.JSON(http.StatusOK, result)
		})
		abilityGroup.GET("/:id", func(c *gin.Context) {
			result, err := GetAbilityByID(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": err})
				return
			}
			c.JSON(http.StatusOK, result)
		})
	}
}

func setupAchievementRoutes(group *gin.RouterGroup) {
	achievementGroup := group.Group("/achievement")
	{
		achievementGroup.POST("", func(c *gin.Context) {
			result, err := CreateAchievement(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": err})
				return
			}
			c.JSON(http.StatusOK, result)
		})
		achievementGroup.GET("", func(c *gin.Context) {
			result, err := GetAllAchievement(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": err})
				return
			}
			c.JSON(http.StatusOK, result)
		})
		achievementGroup.GET("/:id", func(c *gin.Context) {
			result, err := GetAchievementByID(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": err})
				return
			}
			c.JSON(http.StatusOK, result)
		})
	}
}

func setupCharacterRoutes(group *gin.RouterGroup) {
	characterGroup := group.Group("/character")
	{
		characterGroup.POST("", func(c *gin.Context) {
			result, err := CreateCharacter(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": err})
				return
			}
			c.JSON(http.StatusOK, result)
		})
		characterGroup.GET("", func(c *gin.Context) {
			result, err := GetAllCharacter(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": err})
				return
			}
			c.JSON(http.StatusOK, result)
		})
		characterGroup.GET("/:id", func(c *gin.Context) {
			result, err := GetCharacterByID(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": err})
				return
			}
			c.JSON(http.StatusOK, result)
		})
	}
}

func setupCosmeticRoutes(group *gin.RouterGroup) {
	cosmeticGroup := group.Group("/cosmetic")
	{
		cosmeticGroup.POST("", func(c *gin.Context) {
			result, err := CreateCosmetic(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": err})
				return
			}
			c.JSON(http.StatusOK, result)
		})
		cosmeticGroup.GET("", func(c *gin.Context) {
			result, err := GetAllCosmetic(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": err})
				return
			}
			c.JSON(http.StatusOK, result)
		})
		cosmeticGroup.GET("/:id", func(c *gin.Context) {
			result, err := GetCosmeticByID(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": err})
				return
			}
			c.JSON(http.StatusOK, result)
		})
	}
}

func setupLoreRoutes(group *gin.RouterGroup) {
	loreGroup := group.Group("/lore")
	{
		loreGroup.POST("", func(c *gin.Context) {
			result, err := CreateLore(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": err})
				return
			}
			c.JSON(http.StatusOK, result)
		})
		loreGroup.GET("", func(c *gin.Context) {
			result, err := GetAllLore(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": err})
				return
			}
			c.JSON(http.StatusOK, result)
		})
		loreGroup.GET("/:id", func(c *gin.Context) {
			result, err := GetLoreByID(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": err})
				return
			}
			c.JSON(http.StatusOK, result)
		})
	}
}
