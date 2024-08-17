package comics

type ComicDataWrapper struct {
	Code            int                `json:"code"`
	Status          string             `json:"status"`
	Copyright       string             `json:"copyright"`
	AttributionText string             `json:"attributionText"`
	AttributionHTML string             `json:"attributionHTML"`
	Data            ComicDataContainer `json:"data"`
	Etag            string             `json:"etag"`
}
