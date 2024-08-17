package comics

type CreatorDataContainer struct {
	Offset  int       `json:"offset"`
	Limit   int       `json:"limit"`
	Total   int       `json:"total"`
	Count   int       `json:"count"`
	Results []Creator `json:"results"`
}
