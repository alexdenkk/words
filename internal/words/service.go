package words

import (
	"context"
	"alexdenkk/words/model"
)

type Service interface {
	CountWords(context.Context, string) model.Result
}
