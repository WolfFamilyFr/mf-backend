package models

type TextObject struct {
	Type     string `json:"type"`
	Language string `json:"language"`
	Text     string `json:"text"`
}
