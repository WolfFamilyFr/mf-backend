package comics

type EventDataWrapper struct {
	Code            int                `json:"code"`
	Status          string             `json:"status"`
	Copyright       string             `json:"copyright"`
	AttributionText string             `json:"attributionText"`
	AttributionHTML string             `json:"attributionHTML"`
	Data            EventDataContainer `json:"data"`
	Etag            string             `json:"etag"`
}
