package comics

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	group := router.Group("/comics")
	{
		setupCharactersRoutes(group)
		setupComicsRoutes(group)
		setupCreatorsRoutes(group)
		setupEventsRoutes(group)
		setupSeriesRoutes(group)
		setupStoriesRoutes(group)
	}
}

func setupCharactersRoutes(group *gin.RouterGroup) {
	charactersGroup := group.Group("/characters")
	{
		charactersGroup.GET("", func(c *gin.Context) {
			result, err := GetAllCharacter(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": err})
				return
			}
			c.JSON(http.StatusOK, result)
		})
		charactersGroup.GET("/:id", func(c *gin.Context) {
			result, err := GetCharacterByID(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": err})
				return
			}
			c.JSON(http.StatusOK, result)
		})
		charactersGroup.GET("/:id/comics", func(c *gin.Context) {
			result, err := GetAllComicByCharacterID(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": err})
				return
			}
			c.JSON(http.StatusOK, result)
		})
		charactersGroup.GET("/:id/events", func(c *gin.Context) {
			result, err := GetAllEventByCharacterID(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": err})
				return
			}
			c.JSON(http.StatusOK, result)
		})
		charactersGroup.GET("/:id/series", func(c *gin.Context) {
			result, err := GetAllSeriesByCharacterID(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": err})
				return
			}
			c.JSON(http.StatusOK, result)
		})
		charactersGroup.GET("/:id/stories", func(c *gin.Context) {
			result, err := GetAllStoriesByCharacterID(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": err})
				return
			}
			c.JSON(http.StatusOK, result)
		})
	}
}

func setupComicsRoutes(group *gin.RouterGroup) {
	comicsGroup := group.Group("/comics")
	{
		comicsGroup.GET("", func(c *gin.Context) {
			result, err := GetAllComic(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": err})
				return
			}
			c.JSON(http.StatusOK, result)
		})
		comicsGroup.GET("/:id", func(c *gin.Context) {
			result, err := GetComicByID(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": err})
				return
			}
			c.JSON(http.StatusOK, result)
		})
		comicsGroup.GET("/:id/characters", func(c *gin.Context) {
			result, err := GetAllCharacterByComicID(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": err})
				return
			}
			c.JSON(http.StatusOK, result)
		})
		comicsGroup.GET("/:id/creators", func(c *gin.Context) {
			result, err := GetAllCreatorByComicID(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": err})
				return
			}
			c.JSON(http.StatusOK, result)
		})
		comicsGroup.GET("/:id/events", func(c *gin.Context) {
			result, err := GetAllEventByComicID(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": err})
				return
			}
			c.JSON(http.StatusOK, result)
		})
		comicsGroup.GET("/:id/stories", func(c *gin.Context) {
			result, err := GetAllStoryByComicID(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": err})
				return
			}
			c.JSON(http.StatusOK, result)
		})
	}
}

func setupCreatorsRoutes(group *gin.RouterGroup) {
	creatorsGroup := group.Group("/creators")
	{
		creatorsGroup.GET("", func(c *gin.Context) {
			result, err := GetAllCreator(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": err})
				return
			}
			c.JSON(http.StatusOK, result)
		})
		creatorsGroup.GET("/:id", func(c *gin.Context) {
			result, err := GetCreatorByID(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": err})
				return
			}
			c.JSON(http.StatusOK, result)
		})
		creatorsGroup.GET("/:id/comics", func(c *gin.Context) {
			result, err := GetAllComicByCreatorID(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": err})
				return
			}
			c.JSON(http.StatusOK, result)
		})
		creatorsGroup.GET("/:id/events", func(c *gin.Context) {
			result, err := GetAllEventByCreatorID(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": err})
				return
			}
			c.JSON(http.StatusOK, result)
		})
		creatorsGroup.GET("/:id/series", func(c *gin.Context) {
			result, err := GetAllSeriesByCreatorID(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": err})
				return
			}
			c.JSON(http.StatusOK, result)
		})
		creatorsGroup.GET("/:id/stories", func(c *gin.Context) {
			result, err := GetAllStoryByCreatorID(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": err})
				return
			}
			c.JSON(http.StatusOK, result)
		})
	}
}

func setupEventsRoutes(group *gin.RouterGroup) {
	eventsGroup := group.Group("/events")
	{
		eventsGroup.GET("", func(c *gin.Context) {
			result, err := GetAllEvent(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": err})
				return
			}
			c.JSON(http.StatusOK, result)
		})
		eventsGroup.GET("/:id", func(c *gin.Context) {
			result, err := GetEventByID(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": err})
				return
			}
			c.JSON(http.StatusOK, result)
		})
		eventsGroup.GET("/:id/characters", func(c *gin.Context) {
			result, err := GetAllCharacterByEventID(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": err})
				return
			}
			c.JSON(http.StatusOK, result)
		})
		eventsGroup.GET("/:id/comics", func(c *gin.Context) {
			result, err := GetAllComicByEventID(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": err})
				return
			}
			c.JSON(http.StatusOK, result)
		})
		eventsGroup.GET("/:id/creators", func(c *gin.Context) {
			result, err := GetAllCreatorByEventID(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": err})
				return
			}
			c.JSON(http.StatusOK, result)
		})
		eventsGroup.GET("/:id/series", func(c *gin.Context) {
			result, err := GetAllSeriesByEventID(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": err})
				return
			}
			c.JSON(http.StatusOK, result)
		})
		eventsGroup.GET("/:id/stories", func(c *gin.Context) {
			result, err := GetAllStoryByEventID(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": err})
				return
			}
			c.JSON(http.StatusOK, result)
		})
	}

}

func setupSeriesRoutes(group *gin.RouterGroup) {
	seriesGroup := group.Group("/series")
	{
		seriesGroup.GET("", func(c *gin.Context) {
			result, err := GetAllSeries(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": err})
				return
			}
			c.JSON(http.StatusOK, result)
		})
		seriesGroup.GET("/:id", func(c *gin.Context) {
			result, err := GetSeriesByID(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": err})
				return
			}
			c.JSON(http.StatusOK, result)
		})
		seriesGroup.GET("/:id/characters", func(c *gin.Context) {
			result, err := GetAllCharacterBySeriesID(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": err})
				return
			}
			c.JSON(http.StatusOK, result)
		})
		seriesGroup.GET("/:id/comics", func(c *gin.Context) {
			result, err := GetAllComicBySeriesID(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": err})
				return
			}
			c.JSON(http.StatusOK, result)
		})
		seriesGroup.GET("/:id/creators", func(c *gin.Context) {
			result, err := GetAllCreatorBySeriesID(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": err})
				return
			}
			c.JSON(http.StatusOK, result)
		})
		seriesGroup.GET("/:id/events", func(c *gin.Context) {
			result, err := GetAllEventBySeriesID(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": err})
				return
			}
			c.JSON(http.StatusOK, result)
		})
		seriesGroup.GET("/:id/stories", func(c *gin.Context) {
			result, err := GetAllStoryBySeriesID(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": err})
				return
			}
			c.JSON(http.StatusOK, result)
		})
	}
}

func setupStoriesRoutes(group *gin.RouterGroup) {
	storiesGroup := group.Group("/stories")
	{
		storiesGroup.GET("", func(c *gin.Context) {
			result, err := GetAllStory(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": err})
				return
			}
			c.JSON(http.StatusOK, result)
		})
		storiesGroup.GET("/:id", func(c *gin.Context) {
			result, err := GetStoryByID(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": err})
				return
			}
			c.JSON(http.StatusOK, result)
		})
		storiesGroup.GET("/:id/characters", func(c *gin.Context) {
			result, err := GetAllCharacterByStoryID(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": err})
				return
			}
			c.JSON(http.StatusOK, result)
		})
		storiesGroup.GET("/:id/comics", func(c *gin.Context) {
			result, err := GetAllComicByStoryID(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": err})
				return
			}
			c.JSON(http.StatusOK, result)
		})
		storiesGroup.GET("/:id/creators", func(c *gin.Context) {
			result, err := GetAllCreatorByStoryID(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": err})
				return
			}
			c.JSON(http.StatusOK, result)
		})
		storiesGroup.GET("/:id/events", func(c *gin.Context) {
			result, err := GetAllEventByStoryID(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": err})
				return
			}
			c.JSON(http.StatusOK, result)
		})
		storiesGroup.GET("/:id/series", func(c *gin.Context) {
			result, err := GetAllSeriesByStoryID(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": err})
				return
			}
			c.JSON(http.StatusOK, result)
		})
	}

}
