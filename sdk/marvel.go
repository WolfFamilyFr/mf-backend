package sdk

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	models "github.com/bchaillou003/marvel-family-backend/models/comics"
	"github.com/gofiber/fiber/v2/log"
)

func GetAllCharacter() (*models.CharacterDataWrapper, error) {
	url := buildURL("/v1/public/characters")
	var out models.CharacterDataWrapper
	err := get(url, &out)
	return &out, err
}

func GetCharacterByID(id int) (*models.CharacterDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/characters/%d", id))
	var out models.CharacterDataWrapper
	err := get(url, &out)
	return &out, err
}

func GetAllComicByCharacterID(characterID int) (*models.ComicDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/characters/%d/comics", characterID))
	var out models.ComicDataWrapper
	err := get(url, &out)
	return &out, err
}

func GetAllEventByCharacterID(characterID int) (*models.EventDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/characters/%d/events", characterID))
	var out models.EventDataWrapper
	err := get(url, &out)
	return &out, err
}

func GetAllSeriesByCharacterID(characterID int) (*models.SeriesDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/characters/%d/series", characterID))
	var out models.SeriesDataWrapper
	err := get(url, &out)
	return &out, err
}

func GetAllStoriesByCharacterID(characterID int) (*models.StoryDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/characters/%d/stories", characterID))
	var out models.StoryDataWrapper
	err := get(url, &out)
	return &out, err
}

func GetAllComic() (*models.ComicDataWrapper, error) {
	url := buildURL("/v1/public/comics")
	var out models.ComicDataWrapper
	err := get(url, &out)
	return &out, err
}

func GetComicByID(id int) (*models.ComicDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/comics/%d", id))
	var out models.ComicDataWrapper
	err := get(url, &out)
	return &out, err
}

func GetAllCharacterByComicID(comicID int) (*models.CharacterDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/comics/%d/characters", comicID))
	var out models.CharacterDataWrapper
	err := get(url, &out)
	return &out, err
}

func GetAllCreatorByComicID(comicID int) (*models.CreatorDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/comics/%d/creators", comicID))
	var out models.CreatorDataWrapper
	err := get(url, &out)
	return &out, err
}

func GetAllEventByComicID(comicID int) (*models.EventDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/comics/%d/events", comicID))
	var out models.EventDataWrapper
	err := get(url, &out)
	return &out, err
}

func GetAllStoryByComicID(comicID int) (*models.StoryDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/comics/%d/stories", comicID))
	var out models.StoryDataWrapper
	err := get(url, &out)
	return &out, err
}

func GetAllCreator() (*models.CreatorDataWrapper, error) {
	url := buildURL("/v1/public/creators")
	var out models.CreatorDataWrapper
	err := get(url, &out)
	return &out, err
}

func GetCreatorByID(id int) (*models.CreatorDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/creators/%d", id))
	var out models.CreatorDataWrapper
	err := get(url, &out)
	return &out, err
}

func GetAllComicByCreatorID(creatorID int) (*models.ComicDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/creators/%d/comics", creatorID))
	var out models.ComicDataWrapper
	err := get(url, &out)
	return &out, err
}

func GetAllEventByCreatorID(creatorID int) (*models.EventDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/creators/%d/events", creatorID))
	var out models.EventDataWrapper
	err := get(url, &out)
	return &out, err
}

func GetAllSeriesByCreatorID(creatorID int) (*models.SeriesDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/creators/%d/series", creatorID))
	var out models.SeriesDataWrapper
	err := get(url, &out)
	return &out, err
}

func GetAllStoryByCreatorID(creatorID int) (*models.StoryDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/creators/%d/stories", creatorID))
	var out models.StoryDataWrapper
	err := get(url, &out)
	return &out, err
}

func GetAllEvent() (*models.EventDataWrapper, error) {
	url := buildURL("/v1/public/events")
	var out models.EventDataWrapper
	err := get(url, &out)
	return &out, err
}

func GetEventByID(id int) (*models.EventDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/events/%d", id))
	var out models.EventDataWrapper
	err := get(url, &out)
	return &out, err
}

func GetAllCharacterByEventID(eventID int) (*models.CharacterDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/events/%d/characters", eventID))
	var out models.CharacterDataWrapper
	err := get(url, &out)
	return &out, err
}

func GetAllComicByEventID(eventID int) (*models.ComicDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/events/%d/comics", eventID))
	var out models.ComicDataWrapper
	err := get(url, &out)
	return &out, err
}

func GetAllCreatorByEventID(eventID int) (*models.CreatorDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/events/%d/creators", eventID))
	var out models.CreatorDataWrapper
	err := get(url, &out)
	return &out, err
}

func GetAllSeriesByEventID(eventID int) (*models.SeriesDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/events/%d/series", eventID))
	var out models.SeriesDataWrapper
	err := get(url, &out)
	return &out, err
}

func GetAllStoryByEventID(eventID int) (*models.StoryDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/events/%d/stories", eventID))
	var out models.StoryDataWrapper
	err := get(url, &out)
	return &out, err
}

func GetAllSeries() (*models.SeriesDataWrapper, error) {
	url := buildURL("/v1/public/series")
	var out models.SeriesDataWrapper
	err := get(url, &out)
	return &out, err
}

func GetSeriesByID(id int) (*models.SeriesDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/series/%d", id))
	var out models.SeriesDataWrapper
	err := get(url, &out)
	return &out, err
}

func GetAllCharacterBySeriesID(seriesID int) (*models.CharacterDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/series/%d/characters", seriesID))
	var out models.CharacterDataWrapper
	err := get(url, &out)
	return &out, err
}

func GetAllComicBySeriesID(seriesID int) (*models.ComicDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/series/%d/comics", seriesID))
	var out models.ComicDataWrapper
	err := get(url, &out)
	return &out, err
}

func GetAllCreatorBySeriesID(seriesID int) (*models.CreatorDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/series/%d/creators", seriesID))
	var out models.CreatorDataWrapper
	err := get(url, &out)
	return &out, err
}

func GetAllEventBySeriesID(seriesID int) (*models.EventDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/series/%d/events", seriesID))
	var out models.EventDataWrapper
	err := get(url, &out)
	return &out, err
}

func GetAllStoryBySeriesID(seriesID int) (*models.StoryDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/series/%d/stories", seriesID))
	var out models.StoryDataWrapper
	err := get(url, &out)
	return &out, err
}

func GetAllStory() (*models.StoryDataWrapper, error) {
	url := buildURL("/v1/public/stories")
	var out models.StoryDataWrapper
	err := get(url, &out)
	return &out, err
}

func GetStoryByID(id int) (*models.StoryDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/stories/%d", id))
	var out models.StoryDataWrapper
	err := get(url, &out)
	return &out, err
}

func GetAllCharacterByStoryID(storyID int) (*models.CharacterDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/stories/%d/characters", storyID))
	var out models.CharacterDataWrapper
	err := get(url, &out)
	return &out, err
}

func GetAllComicByStoryID(storyID int) (*models.ComicDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/stories/%d/comics", storyID))
	var out models.ComicDataWrapper
	err := get(url, &out)
	return &out, err
}

func GetAllCreatorByStoryID(storyID int) (*models.CreatorDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/stories/%d/creators", storyID))
	var out models.CreatorDataWrapper
	err := get(url, &out)
	return &out, err
}

func GetAllEventByStoryID(storyID int) (*models.EventDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/stories/%d/events", storyID))
	var out models.EventDataWrapper
	err := get(url, &out)
	return &out, err
}

func GetAllSeriesByStoryID(storyID int) (*models.SeriesDataWrapper, error) {
	url := buildURL(fmt.Sprintf("/v1/public/stories/%d/series", storyID))
	var out models.SeriesDataWrapper
	err := get(url, &out)
	return &out, err
}

func get(url string, out interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		log.Error("err : ", err.Error())
		return fmt.Errorf("ooopsss an error occurred, please try again")
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
		log.Error("err : ", err.Error())
		return fmt.Errorf("ooopsss! an error occurred, please try again")
	}
	return nil
}

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
