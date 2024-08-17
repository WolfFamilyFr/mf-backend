package comics

type ComicDataContainer struct {
	Offset  int     `json:"offset"`
	Limit   int     `json:"limit"`
	Total   int     `json:"total"`
	Count   int     `json:"count"`
	Results []Comic `json:"results"`
}
