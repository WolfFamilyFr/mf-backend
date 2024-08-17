package models

type CharacterList struct {
	Available     int                `json:"available"`
	Returned      int                `json:"returned"`
	CollectionURI string             `json:"collectionURI"`
	Items         []CharacterSummary `json:"items"`
}
