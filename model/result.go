package model

type Result struct {
	WordsTotal uint `json:"words_total"`
	WordsCount map[string]uint `json:"words_count"`
}
