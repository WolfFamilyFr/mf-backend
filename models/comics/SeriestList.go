package models

type SeriesList struct {
	Available     int             `json:"available"`
	Returned      int             `json:"returned"`
	CollectionURI string          `json:"collectionURI"`
	Items         []SeriesSummary `json:"items"`
}
