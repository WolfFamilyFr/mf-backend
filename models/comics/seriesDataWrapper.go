package models

type SeriesDataWrapper struct {
	Code            int                 `json:"code"`
	Status          string              `json:"status"`
	Copyright       string              `json:"copyright"`
	AttributionText string              `json:"attributionText"`
	AttributionHTML string              `json:"attributionHTML"`
	Data            SeriesDataContainer `json:"data"`
	Etag            string              `json:"etag"`
}
