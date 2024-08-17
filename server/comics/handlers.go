package comics

import (
	"strconv"

	"github.com/bchaillou003/marvel-family-backend/sdk"
	"github.com/gofiber/fiber/v2"
)

func GetAllCharacter(c *fiber.Ctx) error {
	result, err := sdk.GetAllCharacter()
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.Status(200).JSON(result)
}

func GetCharacterByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(500).JSON(err)
	}

	result, err := sdk.GetCharacterByID(id)
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.Status(200).JSON(result)
}

func GetAllComicByCharacterID(c *fiber.Ctx) error {
	characterID, err := strconv.Atoi(c.Params("characterID"))
	if err != nil {
		return c.Status(500).JSON(err)
	}

	result, err := sdk.GetAllComicByCharacterID(characterID)
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.Status(200).JSON(result)
}

func GetAllEventByCharacterID(c *fiber.Ctx) error {
	characterID, err := strconv.Atoi(c.Params("characterID"))
	if err != nil {
		return c.Status(500).JSON(err)
	}

	result, err := sdk.GetAllEventByCharacterID(characterID)
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.Status(200).JSON(result)
}

func GetAllSeriesByCharacterID(c *fiber.Ctx) error {
	characterID, err := strconv.Atoi(c.Params("characterID"))
	if err != nil {
		return c.Status(500).JSON(err)
	}

	result, err := sdk.GetAllSeriesByCharacterID(characterID)
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.Status(200).JSON(result)
}

func GetAllStoriesByCharacterID(c *fiber.Ctx) error {
	characterID, err := strconv.Atoi(c.Params("characterID"))
	if err != nil {
		return c.Status(500).JSON(err)
	}

	result, err := sdk.GetAllStoriesByCharacterID(characterID)
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.Status(200).JSON(result)
}

func GetAllComic(c *fiber.Ctx) error {
	result, err := sdk.GetAllComic()
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.Status(200).JSON(result)
}

func GetComicByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(500).JSON(err)
	}

	result, err := sdk.GetCharacterByID(id)
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.Status(200).JSON(result)
}

func GetAllCharacterByComicID(c *fiber.Ctx) error {
	comicID, err := strconv.Atoi(c.Params("comicID"))
	if err != nil {
		return c.Status(500).JSON(err)
	}

	result, err := sdk.GetAllCharacterByComicID(comicID)
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.Status(200).JSON(result)
}

func GetAllCreatorByComicID(c *fiber.Ctx) error {
	comicID, err := strconv.Atoi(c.Params("comicID"))
	if err != nil {
		return c.Status(500).JSON(err)
	}

	result, err := sdk.GetAllCreatorByComicID(comicID)
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.Status(200).JSON(result)
}

func GetAllEventByComicID(c *fiber.Ctx) error {
	comicID, err := strconv.Atoi(c.Params("comicID"))
	if err != nil {
		return c.Status(500).JSON(err)
	}

	result, err := sdk.GetAllEventByComicID(comicID)
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.Status(200).JSON(result)
}

func GetAllStoryByComicID(c *fiber.Ctx) error {
	comicID, err := strconv.Atoi(c.Params("comicID"))
	if err != nil {
		return c.Status(500).JSON(err)
	}

	result, err := sdk.GetAllStoryByComicID(comicID)
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.Status(200).JSON(result)
}

func GetAllCreator(c *fiber.Ctx) error {
	result, err := sdk.GetAllCreator()
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.Status(200).JSON(result)
}

func GetCreatorByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(500).JSON(err)
	}

	result, err := sdk.GetCreatorByID(id)
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.Status(200).JSON(result)
}

func GetAllComicByCreatorID(c *fiber.Ctx) error {
	creatorID, err := strconv.Atoi(c.Params("creatorID"))
	if err != nil {
		return c.Status(500).JSON(err)
	}

	result, err := sdk.GetAllComicByCreatorID(creatorID)
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.Status(200).JSON(result)
}

func GetAllEventByCreatorID(c *fiber.Ctx) error {
	creatorID, err := strconv.Atoi(c.Params("creatorID"))
	if err != nil {
		return c.Status(500).JSON(err)
	}

	result, err := sdk.GetAllEventByCreatorID(creatorID)
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.Status(200).JSON(result)
}

func GetAllSeriesByCreatorID(c *fiber.Ctx) error {
	creatorID, err := strconv.Atoi(c.Params("creatorID"))
	if err != nil {
		return c.Status(500).JSON(err)
	}

	result, err := sdk.GetAllSeriesByCreatorID(creatorID)
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.Status(200).JSON(result)
}

func GetAllStoryByCreatorID(c *fiber.Ctx) error {
	creatorID, err := strconv.Atoi(c.Params("creatorID"))
	if err != nil {
		return c.Status(500).JSON(err)
	}

	result, err := sdk.GetAllStoryByCreatorID(creatorID)
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.Status(200).JSON(result)
}

func GetAllEvent(c *fiber.Ctx) error {
	result, err := sdk.GetAllEvent()
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.Status(200).JSON(result)
}

func GetEventByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(500).JSON(err)
	}

	result, err := sdk.GetEventByID(id)
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.Status(200).JSON(result)
}

func GetAllCharacterByEventID(c *fiber.Ctx) error {
	eventID, err := strconv.Atoi(c.Params("eventID"))
	if err != nil {
		return c.Status(500).JSON(err)
	}

	result, err := sdk.GetAllCharacterByEventID(eventID)
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.Status(200).JSON(result)
}

func GetAllComicByEventID(c *fiber.Ctx) error {
	eventID, err := strconv.Atoi(c.Params("eventID"))
	if err != nil {
		return c.Status(500).JSON(err)
	}

	result, err := sdk.GetAllComicByEventID(eventID)
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.Status(200).JSON(result)
}

func GetAllCreatorByEventID(c *fiber.Ctx) error {
	eventID, err := strconv.Atoi(c.Params("eventID"))
	if err != nil {
		return c.Status(500).JSON(err)
	}

	result, err := sdk.GetAllCreatorByEventID(eventID)
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.Status(200).JSON(result)
}

func GetAllSeriesByEventID(c *fiber.Ctx) error {
	eventID, err := strconv.Atoi(c.Params("eventID"))
	if err != nil {
		return c.Status(500).JSON(err)
	}

	result, err := sdk.GetAllSeriesByEventID(eventID)
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.Status(200).JSON(result)
}

func GetAllStoryByEventID(c *fiber.Ctx) error {
	eventID, err := strconv.Atoi(c.Params("eventID"))
	if err != nil {
		return c.Status(500).JSON(err)
	}

	result, err := sdk.GetAllStoryByEventID(eventID)
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.Status(200).JSON(result)
}

func GetAllSeries(c *fiber.Ctx) error {
	result, err := sdk.GetAllSeries()
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.Status(200).JSON(result)
}

func GetSeriesByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(500).JSON(err)
	}

	result, err := sdk.GetSeriesByID(id)
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.Status(200).JSON(result)
}

func GetAllCharacterBySeriesID(c *fiber.Ctx) error {
	seriesID, err := strconv.Atoi(c.Params("seriesID"))
	if err != nil {
		return c.Status(500).JSON(err)
	}

	result, err := sdk.GetAllCharacterBySeriesID(seriesID)
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.Status(200).JSON(result)
}

func GetAllComicBySeriesID(c *fiber.Ctx) error {
	seriesID, err := strconv.Atoi(c.Params("seriesID"))
	if err != nil {
		return c.Status(500).JSON(err)
	}

	result, err := sdk.GetAllComicBySeriesID(seriesID)
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.Status(200).JSON(result)
}

func GetAllCreatorBySeriesID(c *fiber.Ctx) error {
	seriesID, err := strconv.Atoi(c.Params("seriesID"))
	if err != nil {
		return c.Status(500).JSON(err)
	}

	result, err := sdk.GetAllCreatorBySeriesID(seriesID)
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.Status(200).JSON(result)
}

func GetAllEventBySeriesID(c *fiber.Ctx) error {
	seriesID, err := strconv.Atoi(c.Params("seriesID"))
	if err != nil {
		return c.Status(500).JSON(err)
	}

	result, err := sdk.GetAllEventBySeriesID(seriesID)
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.Status(200).JSON(result)
}

func GetAllStoryBySeriesID(c *fiber.Ctx) error {
	seriesID, err := strconv.Atoi(c.Params("seriesID"))
	if err != nil {
		return c.Status(500).JSON(err)
	}

	result, err := sdk.GetAllStoryBySeriesID(seriesID)
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.Status(200).JSON(result)
}

func GetAllStory(c *fiber.Ctx) error {
	result, err := sdk.GetAllStory()
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.Status(200).JSON(result)
}

func GetStoryByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(500).JSON(err)
	}

	result, err := sdk.GetStoryByID(id)
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.Status(200).JSON(result)
}

func GetAllCharacterByStoryID(c *fiber.Ctx) error {
	storyID, err := strconv.Atoi(c.Params("storyID"))
	if err != nil {
		return c.Status(500).JSON(err)
	}

	result, err := sdk.GetAllCharacterByStoryID(storyID)
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.Status(200).JSON(result)
}

func GetAllComicByStoryID(c *fiber.Ctx) error {
	storyID, err := strconv.Atoi(c.Params("storyID"))
	if err != nil {
		return c.Status(500).JSON(err)
	}

	result, err := sdk.GetAllComicByStoryID(storyID)
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.Status(200).JSON(result)
}

func GetAllCreatorByStoryID(c *fiber.Ctx) error {
	storyID, err := strconv.Atoi(c.Params("storyID"))
	if err != nil {
		return c.Status(500).JSON(err)
	}

	result, err := sdk.GetAllCreatorByStoryID(storyID)
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.Status(200).JSON(result)
}

func GetAllEventByStoryID(c *fiber.Ctx) error {
	storyID, err := strconv.Atoi(c.Params("storyID"))
	if err != nil {
		return c.Status(500).JSON(err)
	}

	result, err := sdk.GetAllEventByStoryID(storyID)
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.Status(200).JSON(result)
}

func GetAllSeriesByStoryID(c *fiber.Ctx) error {
	storyID, err := strconv.Atoi(c.Params("storyID"))
	if err != nil {
		return c.Status(500).JSON(err)
	}

	result, err := sdk.GetAllSeriesByStoryID(storyID)
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.Status(200).JSON(result)
}
