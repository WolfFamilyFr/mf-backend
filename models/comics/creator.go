package models

type Creator struct {
	ID          int        `json:"id"`
	FirstName   string     `json:"firstName"`
	MiddleName  string     `json:"middleName"`
	LastName    string     `json:"lastName"`
	Suffix      string     `json:"suffix"`
	FullName    string     `json:"fullName"`
	Modified    string     `json:"modified"`
	ResourceURI string     `json:"resourceURI"`
	URLs        []URL      `json:"urls"`
	Thumbnail   Image      `json:"thumbnail"`
	Series      SeriesList `json:"series"`
	Stories     StoryList  `json:"stories"`
	Comics      ComicList  `json:"comics"`
	Events      EventList  `json:"events"`
}
