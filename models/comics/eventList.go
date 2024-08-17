package comics

type EventList struct {
	Available     int            `json:"available"`
	Returned      int            `json:"returned"`
	CollectionURI string         `json:"collectionURI"`
	Items         []EventSummary `json:"items"`
}
