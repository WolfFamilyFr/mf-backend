package comics

type EventDataContainer struct {
	Offset  int     `json:"offset"`
	Limit   int     `json:"limit"`
	Total   int     `json:"total"`
	Count   int     `json:"count"`
	Results []Event `json:"results"`
}
