package service

import (
	"context"
	"strings"
	"regexp"
	"golang.org/x/exp/maps"
	"alexdenkk/words/model"
)

func In(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

func (s *Service) CountWords(ctx context.Context, text string) model.Result {

	text = strings.ToLower(text)
	re := regexp.MustCompile(`[[:punct:]]`)
	text= re.ReplaceAllString(text, "")

	words := strings.Split(text, " ")

	wordsCount := map[string]uint{}

	for _, word := range words {
		if In(word, maps.Keys(wordsCount)) {
			wordsCount[word] = wordsCount[word] + 1
		} else {
			wordsCount[word] = 1
		}
	}


	return model.Result{
		WordsTotal: uint(len(words)),
		WordsCount: wordsCount,
	}
}
