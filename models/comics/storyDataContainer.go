package comics

type StoryDataContainer struct {
	Offset  int     `json:"offset"`
	Limit   int     `json:"limit"`
	Total   int     `json:"total"`
	Count   int     `json:"count"`
	Results []Story `json:"results"`
}
