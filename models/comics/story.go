package comics

type Story struct {
	ID            int           `json:"id"`
	Title         string        `json:"title"`
	Description   string        `json:"description"`
	ResourceURI   string        `json:"resourceURI"`
	Type          string        `json:"type"`
	Modified      string        `json:"modified"`
	Thumbnail     Image         `json:"thumbnail"`
	Comics        ComicList     `json:"comics"`
	Series        SeriesList    `json:"series"`
	Events        EventList     `json:"events"`
	Characters    CharacterList `json:"characters"`
	Creators      CreatorList   `json:"creators"`
	Originalissue ComicSummary  `json:"originalissue"`
}
