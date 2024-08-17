package models

type SeriesDataContainer struct {
	Offset  int      `json:"offset"`
	Limit   int      `json:"limit"`
	Total   int      `json:"total"`
	Count   int      `json:"count"`
	Results []Series `json:"results"`
}
