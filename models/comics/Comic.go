package comics

type Comic struct {
	ID                 int            `json:"id"`
	DigitalID          int            `json:"digitalId"`
	Title              string         `json:"title"`
	IssueNumber        float64        `json:"issueNumber"`
	VariantDescription string         `json:"variantDescription"`
	Description        string         `json:"description"`
	Modified           string         `json:"modified"`
	ISBN               string         `json:"isbn"`
	UPC                string         `json:"upc"`
	DiamondCode        string         `json:"diamondCode"`
	EAN                string         `json:"ean"`
	ISSN               string         `json:"issn"`
	Format             string         `json:"format"`
	PageCount          int            `json:"pageCount"`
	TextObjects        []TextObject   `json:"textObjects"`
	ResourceURI        string         `json:"resourceURI"`
	URLs               []URL          `json:"urls"`
	Series             SeriesSummary  `json:"series"`
	Variants           []ComicSummary `json:"variants"`
	Collections        []ComicSummary `json:"collections"`
	CollectedIssues    []ComicSummary `json:"collectedIssues"`
	Dates              []ComicDate    `json:"dates"`
	Prices             []ComicPrice   `json:"prices"`
	Thumbnail          Image          `json:"thumbnail"`
	Images             []Image        `json:"images"`
	Creators           CreatorList    `json:"creators"`
	Characters         CharacterList  `json:"characters"`
	Stories            StoryList      `json:"stories"`
	Events             EventList      `json:"events"`
}
