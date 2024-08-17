package comics

type CharacterDataContainer struct {
	Offset  int         `json:"offset"`
	Limit   int         `json:"limit"`
	Total   int         `json:"total"`
	Count   int         `json:"count"`
	Results []Character `json:"results"`
}
