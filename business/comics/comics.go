package comics

import (
	"context"

	"github.com/bchaillou003/marvel-family-backend/application"
	"github.com/bchaillou003/marvel-family-backend/models/comics"
)

func GetAllCharacter(ctx context.Context) (*comics.CharacterDataWrapper, error) {
	return application.ClientSDKFromContext(ctx).GetAllCharacter()
}

func GetCharacterByID(ctx context.Context, id int64) (*comics.CharacterDataWrapper, error) {
	return application.ClientSDKFromContext(ctx).GetCharacterByID(id)
}

func GetAllComicByCharacterID(ctx context.Context, characterID int64) (*comics.ComicDataWrapper, error) {
	return application.ClientSDKFromContext(ctx).GetAllComicByCharacterID(characterID)
}

func GetAllEventByCharacterID(ctx context.Context, characterID int64) (*comics.EventDataWrapper, error) {
	return application.ClientSDKFromContext(ctx).GetAllEventByCharacterID(characterID)
}

func GetAllSeriesByCharacterID(ctx context.Context, characterID int64) (*comics.SeriesDataWrapper, error) {
	return application.ClientSDKFromContext(ctx).GetAllSeriesByCharacterID(characterID)
}

func GetAllStoriesByCharacterID(ctx context.Context, characterID int64) (*comics.StoryDataWrapper, error) {
	return application.ClientSDKFromContext(ctx).GetAllStoriesByCharacterID(characterID)
}

func GetAllComic(ctx context.Context) (*comics.ComicDataWrapper, error) {
	return application.ClientSDKFromContext(ctx).GetAllComic()
}

func GetComicByID(ctx context.Context, id int64) (*comics.ComicDataWrapper, error) {
	return application.ClientSDKFromContext(ctx).GetComicByID(id)
}

func GetAllCharacterByComicID(ctx context.Context, comicID int64) (*comics.CharacterDataWrapper, error) {
	return application.ClientSDKFromContext(ctx).GetAllCharacterByComicID(comicID)
}

func GetAllCreatorByComicID(ctx context.Context, comicID int64) (*comics.CreatorDataWrapper, error) {
	return application.ClientSDKFromContext(ctx).GetAllCreatorByComicID(comicID)
}

func GetAllEventByComicID(ctx context.Context, comicID int64) (*comics.EventDataWrapper, error) {
	return application.ClientSDKFromContext(ctx).GetAllEventByComicID(comicID)
}

func GetAllStoryByComicID(ctx context.Context, comicID int64) (*comics.StoryDataWrapper, error) {
	return application.ClientSDKFromContext(ctx).GetAllStoryByComicID(comicID)
}

func GetAllCreator(ctx context.Context) (*comics.CreatorDataWrapper, error) {
	return application.ClientSDKFromContext(ctx).GetAllCreator()
}

func GetCreatorByID(ctx context.Context, id int64) (*comics.CreatorDataWrapper, error) {
	return application.ClientSDKFromContext(ctx).GetCreatorByID(id)
}

func GetAllComicByCreatorID(ctx context.Context, creatorID int64) (*comics.ComicDataWrapper, error) {
	return application.ClientSDKFromContext(ctx).GetAllComicByCreatorID(creatorID)
}

func GetAllEventByCreatorID(ctx context.Context, creatorID int64) (*comics.EventDataWrapper, error) {
	return application.ClientSDKFromContext(ctx).GetAllEventByCreatorID(creatorID)
}

func GetAllSeriesByCreatorID(ctx context.Context, creatorID int64) (*comics.SeriesDataWrapper, error) {
	return application.ClientSDKFromContext(ctx).GetAllSeriesByCreatorID(creatorID)
}

func GetAllStoryByCreatorID(ctx context.Context, creatorID int64) (*comics.StoryDataWrapper, error) {
	return application.ClientSDKFromContext(ctx).GetAllStoryByCreatorID(creatorID)
}

func GetAllEvent(ctx context.Context) (*comics.EventDataWrapper, error) {
	return application.ClientSDKFromContext(ctx).GetAllEvent()
}

func GetEventByID(ctx context.Context, id int64) (*comics.EventDataWrapper, error) {
	return application.ClientSDKFromContext(ctx).GetEventByID(id)
}

func GetAllCharacterByEventID(ctx context.Context, eventID int64) (*comics.CharacterDataWrapper, error) {
	return application.ClientSDKFromContext(ctx).GetAllCharacterByEventID(eventID)
}

func GetAllComicByEventID(ctx context.Context, eventID int64) (*comics.ComicDataWrapper, error) {
	return application.ClientSDKFromContext(ctx).GetAllComicByEventID(eventID)
}

func GetAllCreatorByEventID(ctx context.Context, eventID int64) (*comics.CreatorDataWrapper, error) {
	return application.ClientSDKFromContext(ctx).GetAllCreatorByEventID(eventID)
}

func GetAllSeriesByEventID(ctx context.Context, eventID int64) (*comics.SeriesDataWrapper, error) {
	return application.ClientSDKFromContext(ctx).GetAllSeriesByEventID(eventID)
}

func GetAllStoryByEventID(ctx context.Context, eventID int64) (*comics.StoryDataWrapper, error) {
	return application.ClientSDKFromContext(ctx).GetAllStoryByEventID(eventID)
}

func GetAllSeries(ctx context.Context) (*comics.SeriesDataWrapper, error) {
	return application.ClientSDKFromContext(ctx).GetAllSeries()
}

func GetSeriesByID(ctx context.Context, id int64) (*comics.SeriesDataWrapper, error) {
	return application.ClientSDKFromContext(ctx).GetSeriesByID(id)
}

func GetAllCharacterBySeriesID(ctx context.Context, seriesID int64) (*comics.CharacterDataWrapper, error) {
	return application.ClientSDKFromContext(ctx).GetAllCharacterBySeriesID(seriesID)
}

func GetAllComicBySeriesID(ctx context.Context, seriesID int64) (*comics.ComicDataWrapper, error) {
	return application.ClientSDKFromContext(ctx).GetAllComicBySeriesID(seriesID)
}

func GetAllCreatorBySeriesID(ctx context.Context, seriesID int64) (*comics.CreatorDataWrapper, error) {
	return application.ClientSDKFromContext(ctx).GetAllCreatorBySeriesID(seriesID)
}

func GetAllEventBySeriesID(ctx context.Context, seriesID int64) (*comics.EventDataWrapper, error) {
	return application.ClientSDKFromContext(ctx).GetAllEventBySeriesID(seriesID)
}

func GetAllStoryBySeriesID(ctx context.Context, seriesID int64) (*comics.StoryDataWrapper, error) {
	return application.ClientSDKFromContext(ctx).GetAllStoryBySeriesID(seriesID)
}

func GetAllStory(ctx context.Context) (*comics.StoryDataWrapper, error) {
	return application.ClientSDKFromContext(ctx).GetAllStory()
}

func GetStoryByID(ctx context.Context, id int64) (*comics.StoryDataWrapper, error) {
	return application.ClientSDKFromContext(ctx).GetStoryByID(id)
}

func GetAllCharacterByStoryID(ctx context.Context, storyID int64) (*comics.CharacterDataWrapper, error) {
	return application.ClientSDKFromContext(ctx).GetAllCharacterByStoryID(storyID)
}

func GetAllComicByStoryID(ctx context.Context, storyID int64) (*comics.ComicDataWrapper, error) {
	return application.ClientSDKFromContext(ctx).GetAllComicByStoryID(storyID)
}

func GetAllCreatorByStoryID(ctx context.Context, storyID int64) (*comics.CreatorDataWrapper, error) {
	return application.ClientSDKFromContext(ctx).GetAllCreatorByStoryID(storyID)
}

func GetAllEventByStoryID(ctx context.Context, storyID int64) (*comics.EventDataWrapper, error) {
	return application.ClientSDKFromContext(ctx).GetAllEventByStoryID(storyID)
}

func GetAllSeriesByStoryID(ctx context.Context, storyID int64) (*comics.SeriesDataWrapper, error) {
	return application.ClientSDKFromContext(ctx).GetAllSeriesByStoryID(storyID)
}
