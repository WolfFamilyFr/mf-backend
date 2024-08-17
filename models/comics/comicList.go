package models

type ComicList struct {
	Available     int            `json:"available"`
	Returned      int            `json:"returned"`
	CollectionURI string         `json:"collectionURI"`
	Items         []ComicSummary `json:"items"`
}
