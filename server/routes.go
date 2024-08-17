package server

import (
	"github.com/bchaillou003/marvel-family-backend/server/comics"
	"github.com/bchaillou003/marvel-family-backend/server/rivals"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	setupComicsRoutes(app)
	setupRivalsRoutes(app)
}

func setupComicsRoutes(app *fiber.App) {
	app.Get("/characters", comics.GetAllCharacter)
	app.Get("/characters/:id", comics.GetCharacterByID)
	app.Get("/characters/:characterID/comics", comics.GetAllComicByCharacterID)
	app.Get("/characters/:characterID/events", comics.GetAllEventByCharacterID)
	app.Get("/characters/:characterID/series", comics.GetAllSeriesByCharacterID)
	app.Get("/characters/:characterID/stories", comics.GetAllStoriesByCharacterID)

	app.Get("/comics", comics.GetAllComic)
	app.Get("/comics/:id", comics.GetComicByID)
	app.Get("/comics/:comicID/characters", comics.GetAllCharacterByComicID)
	app.Get("/comics/:comicID/creators", comics.GetAllCreatorByComicID)
	app.Get("/comics/:comicID/events", comics.GetAllEventByComicID)
	app.Get("/comics/:comicID/stories", comics.GetAllStoryByComicID)

	app.Get("/creators", comics.GetAllCreator)
	app.Get("/creators/:id", comics.GetCreatorByID)
	app.Get("/creators/:creatorID/comics", comics.GetAllComicByCreatorID)
	app.Get("/creators/:creatorID/events", comics.GetAllEventByCreatorID)
	app.Get("/creators/:creatorID/series", comics.GetAllSeriesByCreatorID)
	app.Get("/creators/:creatorID/stories", comics.GetAllStoryByCreatorID)

	app.Get("/events", comics.GetAllEvent)
	app.Get("/events/:id", comics.GetEventByID)
	app.Get("/events/:eventID/characters", comics.GetAllCharacterByEventID)
	app.Get("/events/:eventID/comics", comics.GetAllComicByEventID)
	app.Get("/events/:eventID/creators", comics.GetAllCreatorByEventID)
	app.Get("/events/:eventID/series", comics.GetAllSeriesByEventID)
	app.Get("/events/:eventID/stories", comics.GetAllStoryByEventID)

	app.Get("/series", comics.GetAllSeries)
	app.Get("/series/:id", comics.GetSeriesByID)
	app.Get("/series/:seriesID/characters", comics.GetAllCharacterBySeriesID)
	app.Get("/series/:seriesID/comics", comics.GetAllComicBySeriesID)
	app.Get("/series/:seriesID/creators", comics.GetAllCreatorBySeriesID)
	app.Get("/series/:seriesID/events", comics.GetAllEventBySeriesID)
	app.Get("/series/:seriesID/stories", comics.GetAllStoryBySeriesID)

	app.Get("/stories", comics.GetAllStory)
	app.Get("/stories/:id", comics.GetStoryByID)
	app.Get("/stories/:storyID/characters", comics.GetAllCharacterByStoryID)
	app.Get("/stories/:storyID/comics", comics.GetAllComicByStoryID)
	app.Get("/stories/:storyID/creators", comics.GetAllCreatorByStoryID)
	app.Get("/stories/:storyID/events", comics.GetAllEventByStoryID)
	app.Get("/stories/:storyID/series", comics.GetAllSeriesByStoryID)
}

func setupRivalsRoutes(app *fiber.App) {
	app.Post("/rivals/ability", rivals.CreateAbility)
	app.Get("/rivals/ability", rivals.GetAllAbility)
	app.Get("/rivals/ability/:id", rivals.GetAbilityByID)
	app.Post("/rivals/achievement", rivals.CreateAchievement)
	app.Get("/rivals/achievement", rivals.GetAllAchievement)
	app.Get("/rivals/achievement/:id", rivals.GetAchievementByID)
	app.Post("/rivals/character", rivals.CreateCharacter)
	app.Get("/rivals/character", rivals.GetAllCharacter)
	app.Get("/rivals/character/:id", rivals.GetCharacterByID)
	app.Post("/rivals/cosmetic", rivals.CreateCosmetic)
	app.Get("/rivals/cosmetic", rivals.GetAllCosmetic)
	app.Get("/rivals/cosmetic/:id", rivals.GetCosmeticByID)
	app.Post("/rivals/lore", rivals.CreateLore)
	app.Get("/rivals/lore", rivals.GetAllLore)
	app.Get("/rivals/lore/:id", rivals.GetLoreByID)

}
