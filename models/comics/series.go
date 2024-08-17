package models

type Series struct {
	ID          int           `json:"id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	ResourceURI string        `json:"resourceURI"`
	URLs        []URL         `json:"urls"`
	StartYear   int           `json:"startYear"`
	EndYear     int           `json:"endYear"`
	Rating      string        `json:"rating"`
	Modified    string        `json:"modified"`
	Thumbnail   Image         `json:"thumbnail"`
	Comics      ComicList     `json:"comics"`
	Stories     StoryList     `json:"stories"`
	Events      EventList     `json:"events"`
	Characters  CharacterList `json:"characters"`
	Creators    CreatorList   `json:"creators"`
	Next        SeriesSummary `json:"next"`
	Previous    SeriesSummary `json:"previous"`
}
