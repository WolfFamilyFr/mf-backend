package comics

import (
	business "github.com/bchaillou003/marvel-family-backend/business/comics"
	"github.com/bchaillou003/marvel-family-backend/models/comics"
	"github.com/bchaillou003/marvel-family-backend/utils"
	"github.com/gin-gonic/gin"
)

const (
	paramID = "id"
)

func GetAllCharacter(c *gin.Context) (*comics.CharacterDataWrapper, error) {
	return business.GetAllCharacter(c.Request.Context())
}

func GetCharacterByID(c *gin.Context) (*comics.CharacterDataWrapper, error) {
	id, err := utils.StringToInt64(c.Param(paramID))
	if err != nil {
		return nil, err
	}
	return business.GetCharacterByID(c.Request.Context(), id)
}

func GetAllComicByCharacterID(c *gin.Context) (*comics.ComicDataWrapper, error) {
	id, err := utils.StringToInt64(c.Param(paramID))
	if err != nil {
		return nil, err
	}
	return business.GetAllComicByCharacterID(c.Request.Context(), id)
}

func GetAllEventByCharacterID(c *gin.Context) (*comics.EventDataWrapper, error) {
	id, err := utils.StringToInt64(c.Param(paramID))
	if err != nil {
		return nil, err
	}
	return business.GetAllEventByCharacterID(c.Request.Context(), id)
}

func GetAllSeriesByCharacterID(c *gin.Context) (*comics.SeriesDataWrapper, error) {
	id, err := utils.StringToInt64(c.Param(paramID))
	if err != nil {
		return nil, err
	}
	return business.GetAllSeriesByCharacterID(c.Request.Context(), id)
}

func GetAllStoryByCharacterID(c *gin.Context) (*comics.StoryDataWrapper, error) {
	id, err := utils.StringToInt64(c.Param(paramID))
	if err != nil {
		return nil, err
	}
	return business.GetAllStoryByCharacterID(c.Request.Context(), id)
}

func GetAllComic(c *gin.Context) (*comics.ComicDataWrapper, error) {
	return business.GetAllComic(c.Request.Context())
}

func GetComicByID(c *gin.Context) (*comics.ComicDataWrapper, error) {
	id, err := utils.StringToInt64(c.Param(paramID))
	if err != nil {
		return nil, err
	}
	return business.GetComicByID(c.Request.Context(), id)
}

func GetAllCharacterByComicID(c *gin.Context) (*comics.CharacterDataWrapper, error) {
	id, err := utils.StringToInt64(c.Param(paramID))
	if err != nil {
		return nil, err
	}
	return business.GetAllCharacterByComicID(c.Request.Context(), id)
}

func GetAllCreatorByComicID(c *gin.Context) (*comics.CreatorDataWrapper, error) {
	id, err := utils.StringToInt64(c.Param(paramID))
	if err != nil {
		return nil, err
	}
	return business.GetAllCreatorByComicID(c.Request.Context(), id)
}

func GetAllEventByComicID(c *gin.Context) (*comics.EventDataWrapper, error) {
	id, err := utils.StringToInt64(c.Param(paramID))
	if err != nil {
		return nil, err
	}
	return business.GetAllEventByComicID(c.Request.Context(), id)
}

func GetAllStoryByComicID(c *gin.Context) (*comics.StoryDataWrapper, error) {
	id, err := utils.StringToInt64(c.Param(paramID))
	if err != nil {
		return nil, err
	}
	return business.GetAllStoryByComicID(c.Request.Context(), id)
}

func GetAllCreator(c *gin.Context) (*comics.CreatorDataWrapper, error) {
	return business.GetAllCreator(c.Request.Context())
}

func GetCreatorByID(c *gin.Context) (*comics.CreatorDataWrapper, error) {
	id, err := utils.StringToInt64(c.Param(paramID))
	if err != nil {
		return nil, err
	}
	return business.GetCreatorByID(c.Request.Context(), id)
}

func GetAllComicByCreatorID(c *gin.Context) (*comics.ComicDataWrapper, error) {
	id, err := utils.StringToInt64(c.Param(paramID))
	if err != nil {
		return nil, err
	}
	return business.GetAllComicByCreatorID(c.Request.Context(), id)
}

func GetAllEventByCreatorID(c *gin.Context) (*comics.EventDataWrapper, error) {
	id, err := utils.StringToInt64(c.Param(paramID))
	if err != nil {
		return nil, err
	}
	return business.GetAllEventByCreatorID(c.Request.Context(), id)
}

func GetAllSeriesByCreatorID(c *gin.Context) (*comics.SeriesDataWrapper, error) {
	id, err := utils.StringToInt64(c.Param(paramID))
	if err != nil {
		return nil, err
	}
	return business.GetAllSeriesByCreatorID(c.Request.Context(), id)
}

func GetAllStoryByCreatorID(c *gin.Context) (*comics.StoryDataWrapper, error) {
	id, err := utils.StringToInt64(c.Param(paramID))
	if err != nil {
		return nil, err
	}
	return business.GetAllStoryByCreatorID(c.Request.Context(), id)
}

func GetAllEvent(c *gin.Context) (*comics.EventDataWrapper, error) {
	return business.GetAllEvent(c.Request.Context())
}

func GetEventByID(c *gin.Context) (*comics.EventDataWrapper, error) {
	id, err := utils.StringToInt64(c.Param(paramID))
	if err != nil {
		return nil, err
	}
	return business.GetEventByID(c.Request.Context(), id)
}

func GetAllCharacterByEventID(c *gin.Context) (*comics.CharacterDataWrapper, error) {
	id, err := utils.StringToInt64(c.Param(paramID))
	if err != nil {
		return nil, err
	}
	return business.GetAllCharacterByEventID(c.Request.Context(), id)
}

func GetAllComicByEventID(c *gin.Context) (*comics.ComicDataWrapper, error) {
	id, err := utils.StringToInt64(c.Param(paramID))
	if err != nil {
		return nil, err
	}
	return business.GetAllComicByEventID(c.Request.Context(), id)
}

func GetAllCreatorByEventID(c *gin.Context) (*comics.CreatorDataWrapper, error) {
	id, err := utils.StringToInt64(c.Param(paramID))
	if err != nil {
		return nil, err
	}
	return business.GetAllCreatorByEventID(c.Request.Context(), id)
}

func GetAllSeriesByEventID(c *gin.Context) (*comics.SeriesDataWrapper, error) {
	id, err := utils.StringToInt64(c.Param(paramID))
	if err != nil {
		return nil, err
	}
	return business.GetAllSeriesByEventID(c.Request.Context(), id)
}

func GetAllStoryByEventID(c *gin.Context) (*comics.StoryDataWrapper, error) {
	id, err := utils.StringToInt64(c.Param(paramID))
	if err != nil {
		return nil, err
	}
	return business.GetAllStoryByEventID(c.Request.Context(), id)
}

func GetAllSeries(c *gin.Context) (*comics.SeriesDataWrapper, error) {
	return business.GetAllSeries(c.Request.Context())
}

func GetSeriesByID(c *gin.Context) (*comics.SeriesDataWrapper, error) {
	id, err := utils.StringToInt64(c.Param(paramID))
	if err != nil {
		return nil, err
	}
	return business.GetSeriesByID(c.Request.Context(), id)
}

func GetAllCharacterBySeriesID(c *gin.Context) (*comics.CharacterDataWrapper, error) {
	id, err := utils.StringToInt64(c.Param(paramID))
	if err != nil {
		return nil, err
	}
	return business.GetAllCharacterBySeriesID(c.Request.Context(), id)
}

func GetAllComicBySeriesID(c *gin.Context) (*comics.ComicDataWrapper, error) {
	id, err := utils.StringToInt64(c.Param(paramID))
	if err != nil {
		return nil, err
	}
	return business.GetAllComicBySeriesID(c.Request.Context(), id)
}

func GetAllCreatorBySeriesID(c *gin.Context) (*comics.CreatorDataWrapper, error) {
	id, err := utils.StringToInt64(c.Param(paramID))
	if err != nil {
		return nil, err
	}
	return business.GetAllCreatorBySeriesID(c.Request.Context(), id)
}

func GetAllEventBySeriesID(c *gin.Context) (*comics.EventDataWrapper, error) {
	id, err := utils.StringToInt64(c.Param(paramID))
	if err != nil {
		return nil, err
	}
	return business.GetAllEventBySeriesID(c.Request.Context(), id)
}

func GetAllStoryBySeriesID(c *gin.Context) (*comics.StoryDataWrapper, error) {
	id, err := utils.StringToInt64(c.Param(paramID))
	if err != nil {
		return nil, err
	}
	return business.GetAllStoryBySeriesID(c.Request.Context(), id)
}

func GetAllStory(c *gin.Context) (*comics.StoryDataWrapper, error) {
	return business.GetAllStory(c.Request.Context())
}

func GetStoryByID(c *gin.Context) (*comics.StoryDataWrapper, error) {
	id, err := utils.StringToInt64(c.Param(paramID))
	if err != nil {
		return nil, err
	}
	return business.GetStoryByID(c.Request.Context(), id)
}

func GetAllCharacterByStoryID(c *gin.Context) (*comics.CharacterDataWrapper, error) {
	id, err := utils.StringToInt64(c.Param(paramID))
	if err != nil {
		return nil, err
	}
	return business.GetAllCharacterByStoryID(c.Request.Context(), id)
}

func GetAllComicByStoryID(c *gin.Context) (*comics.ComicDataWrapper, error) {
	id, err := utils.StringToInt64(c.Param(paramID))
	if err != nil {
		return nil, err
	}
	return business.GetAllComicByStoryID(c.Request.Context(), id)
}

func GetAllCreatorByStoryID(c *gin.Context) (*comics.CreatorDataWrapper, error) {
	id, err := utils.StringToInt64(c.Param(paramID))
	if err != nil {
		return nil, err
	}
	return business.GetAllCreatorByStoryID(c.Request.Context(), id)
}

func GetAllEventByStoryID(c *gin.Context) (*comics.EventDataWrapper, error) {
	id, err := utils.StringToInt64(c.Param(paramID))
	if err != nil {
		return nil, err
	}
	return business.GetAllEventByStoryID(c.Request.Context(), id)
}

func GetAllSeriesByStoryID(c *gin.Context) (*comics.SeriesDataWrapper, error) {
	id, err := utils.StringToInt64(c.Param(paramID))
	if err != nil {
		return nil, err
	}
	return business.GetAllSeriesByStoryID(c.Request.Context(), id)
}
