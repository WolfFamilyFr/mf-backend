package comics

type Character struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Modified    string     `json:"modified"`
	ResourceURI string     `json:"resourceURI"`
	Urls        []URL      `json:"urls"`
	Thumbnail   Image      `json:"thumbnail"`
	Comics      ComicList  `json:"comics"`
	Stories     StoryList  `json:"stories"`
	Events      EventList  `json:"events"`
	Series      SeriesList `json:"series"`
}
