package models

type CreatorList struct {
	Available     int              `json:"available"`
	Returned      int              `json:"returned"`
	CollectionURI string           `json:"collectionURI"`
	Items         []CreatorSummary `json:"items"`
}
