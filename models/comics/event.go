package comics

type Event struct {
	ID          int           `json:"id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	ResourceURI string        `json:"resourceURI"`
	URLs        []URL         `json:"urls"`
	Modified    string        `json:"modified"`
	Start       string        `json:"start"`
	End         string        `json:"end"`
	Thumbnail   Image         `json:"thumbnail"`
	Comics      ComicList     `json:"comics"`
	Series      SeriesList    `json:"series"`
	Characters  CharacterList `json:"characters"`
	Creators    CreatorList   `json:"creators"`
	Next        EventSummary  `json:"next"`
	Previous    EventSummary  `json:"previous"`
}
