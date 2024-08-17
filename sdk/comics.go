package sdk

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/bchaillou003/marvel-family-backend/models/comics"
)

func buildURL(path string) string {
	ts := time.Now().Unix()
	newmd5 := md5.New()
	io.WriteString(newmd5, fmt.Sprintf("%d%s%s", ts, os.Getenv("MARVEL_PRIVATEKEY"), os.Getenv("MARVEL_PUBLICKEY")))
	hash := fmt.Sprintf("%x", newmd5.Sum(nil))

	return os.Getenv("MARVEL_URL") + path +
		"?ts=" + fmt.Sprintf("%d", ts) +
		"&apikey=" + os.Getenv("MARVEL_PUBLICKEY") +
		"&hash=" + hash
}

type IComicsSDK interface {
	GetAllCharacter() (*comics.CharacterDataWrapper, error)
	GetCharacterByID(id int64) (*comics.CharacterDataWrapper, error)
	GetAllComicByCharacterID(characterID int64) (*comics.ComicDataWrapper, error)
	GetAllEventByCharacterID(characterID int64) (*comics.EventDataWrapper, error)
	GetAllSeriesByCharacterID(characterID int64) (*comics.SeriesDataWrapper, error)
	GetAllStoriesByCharacterID(characterID int64) (*comics.StoryDataWrapper, error)

	GetAllComic() (*comics.ComicDataWrapper, error)
	GetComicByID(id int64) (*comics.ComicDataWrapper, error)
	GetAllCharacterByComicID(comicID int64) (*comics.CharacterDataWrapper, error)
	GetAllCreatorByComicID(comicID int64) (*comics.CreatorDataWrapper, error)
	GetAllEventByComicID(comicID int64) (*comics.EventDataWrapper, error)
	GetAllStoryByComicID(comicID int64) (*comics.StoryDataWrapper, error)

	GetAllCreator() (*comics.CreatorDataWrapper, error)
	GetCreatorByID(id int64) (*comics.CreatorDataWrapper, error)
	GetAllComicByCreatorID(creatorID int64) (*comics.ComicDataWrapper, error)
	GetAllEventByCreatorID(creatorID int64) (*comics.EventDataWrapper, error)
	GetAllSeriesByCreatorID(creatorID int64) (*comics.SeriesDataWrapper, error)
	GetAllStoryByCreatorID(creatorID int64) (*comics.StoryDataWrapper, error)
	GetAllEvent() (*comics.EventDataWrapper, error)
	GetEventByID(id int64) (*comics.EventDataWrapper, error)
	GetAllCharacterByEventID(eventID int64) (*comics.CharacterDataWrapper, error)
	GetAllComicByEventID(eventID int64) (*comics.ComicDataWrapper, error)
	GetAllCreatorByEventID(eventID int64) (*comics.CreatorDataWrapper, error)
	GetAllSeriesByEventID(eventID int64) (*comics.SeriesDataWrapper, error)
	GetAllStoryByEventID(eventID int64) (*comics.StoryDataWrapper, error)
	GetAllSeries() (*comics.SeriesDataWrapper, error)
	GetSeriesByID(id int64) (*comics.SeriesDataWrapper, error)
	GetAllCharacterBySeriesID(seriesID int64) (*comics.CharacterDataWrapper, error)
	GetAllComicBySeriesID(seriesID int64) (*comics.ComicDataWrapper, error)
	GetAllCreatorBySeriesID(seriesID int64) (*comics.CreatorDataWrapper, error)
	GetAllEventBySeriesID(seriesID int64) (*comics.EventDataWrapper, error)
	GetAllStoryBySeriesID(seriesID int64) (*comics.StoryDataWrapper, error)
	GetAllStory() (*comics.StoryDataWrapper, error)
	GetStoryByID(id int64) (*comics.StoryDataWrapper, error)
	GetAllCharacterByStoryID(storyID int64) (*comics.CharacterDataWrapper, error)
	GetAllComicByStoryID(storyID int64) (*comics.ComicDataWrapper, error)
	GetAllCreatorByStoryID(storyID int64) (*comics.CreatorDataWrapper, error)
	GetAllEventByStoryID(storyID int64) (*comics.EventDataWrapper, error)
	GetAllSeriesByStoryID(storyID int64) (*comics.SeriesDataWrapper, error)
}

func (c ClientSDK) GetAllCharacter() (*comics.CharacterDataWrapper, error) {
	url := buildURL("/v1/public/characters")
	var out comics.CharacterDataWrapper
	err := c.get(url, &out)
	return &out, err
}

func (c ClientSDK) GetCharacterByID(id int64) (*comics.CharacterDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/characters/%d", id))
	var out comics.CharacterDataWrapper
	err := c.get(url, &out)
	return &out, err
}

func (c ClientSDK) GetAllComicByCharacterID(characterID int64) (*comics.ComicDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/characters/%d/comics", characterID))
	var out comics.ComicDataWrapper
	err := c.get(url, &out)
	return &out, err
}

func (c ClientSDK) GetAllEventByCharacterID(characterID int64) (*comics.EventDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/characters/%d/events", characterID))
	var out comics.EventDataWrapper
	err := c.get(url, &out)
	return &out, err
}

func (c ClientSDK) GetAllSeriesByCharacterID(characterID int64) (*comics.SeriesDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/characters/%d/series", characterID))
	var out comics.SeriesDataWrapper
	err := c.get(url, &out)
	return &out, err
}

func (c ClientSDK) GetAllStoriesByCharacterID(characterID int64) (*comics.StoryDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/characters/%d/stories", characterID))
	var out comics.StoryDataWrapper
	err := c.get(url, &out)
	return &out, err
}

func (c ClientSDK) GetAllComic() (*comics.ComicDataWrapper, error) {
	url := buildURL("/v1/public/comics")
	var out comics.ComicDataWrapper
	err := c.get(url, &out)
	return &out, err
}

func (c ClientSDK) GetComicByID(id int64) (*comics.ComicDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/comics/%d", id))
	var out comics.ComicDataWrapper
	err := c.get(url, &out)
	return &out, err
}

func (c ClientSDK) GetAllCharacterByComicID(comicID int64) (*comics.CharacterDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/comics/%d/characters", comicID))
	var out comics.CharacterDataWrapper
	err := c.get(url, &out)
	return &out, err
}

func (c ClientSDK) GetAllCreatorByComicID(comicID int64) (*comics.CreatorDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/comics/%d/creators", comicID))
	var out comics.CreatorDataWrapper
	err := c.get(url, &out)
	return &out, err
}

func (c ClientSDK) GetAllEventByComicID(comicID int64) (*comics.EventDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/comics/%d/events", comicID))
	var out comics.EventDataWrapper
	err := c.get(url, &out)
	return &out, err
}

func (c ClientSDK) GetAllStoryByComicID(comicID int64) (*comics.StoryDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/comics/%d/stories", comicID))
	var out comics.StoryDataWrapper
	err := c.get(url, &out)
	return &out, err
}

func (c ClientSDK) GetAllCreator() (*comics.CreatorDataWrapper, error) {
	url := buildURL("/v1/public/creators")
	var out comics.CreatorDataWrapper
	err := c.get(url, &out)
	return &out, err
}

func (c ClientSDK) GetCreatorByID(id int64) (*comics.CreatorDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/creators/%d", id))
	var out comics.CreatorDataWrapper
	err := c.get(url, &out)
	return &out, err
}

func (c ClientSDK) GetAllComicByCreatorID(creatorID int64) (*comics.ComicDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/creators/%d/comics", creatorID))
	var out comics.ComicDataWrapper
	err := c.get(url, &out)
	return &out, err
}

func (c ClientSDK) GetAllEventByCreatorID(creatorID int64) (*comics.EventDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/creators/%d/events", creatorID))
	var out comics.EventDataWrapper
	err := c.get(url, &out)
	return &out, err
}

func (c ClientSDK) GetAllSeriesByCreatorID(creatorID int64) (*comics.SeriesDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/creators/%d/series", creatorID))
	var out comics.SeriesDataWrapper
	err := c.get(url, &out)
	return &out, err
}

func (c ClientSDK) GetAllStoryByCreatorID(creatorID int64) (*comics.StoryDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/creators/%d/stories", creatorID))
	var out comics.StoryDataWrapper
	err := c.get(url, &out)
	return &out, err
}

func (c ClientSDK) GetAllEvent() (*comics.EventDataWrapper, error) {
	url := buildURL("/v1/public/events")
	var out comics.EventDataWrapper
	err := c.get(url, &out)
	return &out, err
}

func (c ClientSDK) GetEventByID(id int64) (*comics.EventDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/events/%d", id))
	var out comics.EventDataWrapper
	err := c.get(url, &out)
	return &out, err
}

func (c ClientSDK) GetAllCharacterByEventID(eventID int64) (*comics.CharacterDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/events/%d/characters", eventID))
	var out comics.CharacterDataWrapper
	err := c.get(url, &out)
	return &out, err
}

func (c ClientSDK) GetAllComicByEventID(eventID int64) (*comics.ComicDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/events/%d/comics", eventID))
	var out comics.ComicDataWrapper
	err := c.get(url, &out)
	return &out, err
}

func (c ClientSDK) GetAllCreatorByEventID(eventID int64) (*comics.CreatorDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/events/%d/creators", eventID))
	var out comics.CreatorDataWrapper
	err := c.get(url, &out)
	return &out, err
}

func (c ClientSDK) GetAllSeriesByEventID(eventID int64) (*comics.SeriesDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/events/%d/series", eventID))
	var out comics.SeriesDataWrapper
	err := c.get(url, &out)
	return &out, err
}

func (c ClientSDK) GetAllStoryByEventID(eventID int64) (*comics.StoryDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/events/%d/stories", eventID))
	var out comics.StoryDataWrapper
	err := c.get(url, &out)
	return &out, err
}

func (c ClientSDK) GetAllSeries() (*comics.SeriesDataWrapper, error) {
	url := buildURL("/v1/public/series")
	var out comics.SeriesDataWrapper
	err := c.get(url, &out)
	return &out, err
}

func (c ClientSDK) GetSeriesByID(id int64) (*comics.SeriesDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/series/%d", id))
	var out comics.SeriesDataWrapper
	err := c.get(url, &out)
	return &out, err
}

func (c ClientSDK) GetAllCharacterBySeriesID(seriesID int64) (*comics.CharacterDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/series/%d/characters", seriesID))
	var out comics.CharacterDataWrapper
	err := c.get(url, &out)
	return &out, err
}

func (c ClientSDK) GetAllComicBySeriesID(seriesID int64) (*comics.ComicDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/series/%d/comics", seriesID))
	var out comics.ComicDataWrapper
	err := c.get(url, &out)
	return &out, err
}

func (c ClientSDK) GetAllCreatorBySeriesID(seriesID int64) (*comics.CreatorDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/series/%d/creators", seriesID))
	var out comics.CreatorDataWrapper
	err := c.get(url, &out)
	return &out, err
}

func (c ClientSDK) GetAllEventBySeriesID(seriesID int64) (*comics.EventDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/series/%d/events", seriesID))
	var out comics.EventDataWrapper
	err := c.get(url, &out)
	return &out, err
}

func (c ClientSDK) GetAllStoryBySeriesID(seriesID int64) (*comics.StoryDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/series/%d/stories", seriesID))
	var out comics.StoryDataWrapper
	err := c.get(url, &out)
	return &out, err
}

func (c ClientSDK) GetAllStory() (*comics.StoryDataWrapper, error) {
	url := buildURL("/v1/public/stories")
	var out comics.StoryDataWrapper
	err := c.get(url, &out)
	return &out, err
}

func (c ClientSDK) GetStoryByID(id int64) (*comics.StoryDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/stories/%d", id))
	var out comics.StoryDataWrapper
	err := c.get(url, &out)
	return &out, err
}

func (c ClientSDK) GetAllCharacterByStoryID(storyID int64) (*comics.CharacterDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/stories/%d/characters", storyID))
	var out comics.CharacterDataWrapper
	err := c.get(url, &out)
	return &out, err
}

func (c ClientSDK) GetAllComicByStoryID(storyID int64) (*comics.ComicDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/stories/%d/comics", storyID))
	var out comics.ComicDataWrapper
	err := c.get(url, &out)
	return &out, err
}

func (c ClientSDK) GetAllCreatorByStoryID(storyID int64) (*comics.CreatorDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/stories/%d/creators", storyID))
	var out comics.CreatorDataWrapper
	err := c.get(url, &out)
	return &out, err
}

func (c ClientSDK) GetAllEventByStoryID(storyID int64) (*comics.EventDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/stories/%d/events", storyID))
	var out comics.EventDataWrapper
	err := c.get(url, &out)
	return &out, err
}

func (c ClientSDK) GetAllSeriesByStoryID(storyID int64) (*comics.SeriesDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/stories/%d/series", storyID))
	var out comics.SeriesDataWrapper
	err := c.get(url, &out)
	return &out, err
}
