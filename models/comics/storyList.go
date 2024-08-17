package models

type StoryList struct {
	Available     int            `json:"available"`
	Returned      int            `json:"returned"`
	CollectionURI string         `json:"collectionURI"`
	Items         []StorySummary `json:"items"`
}
