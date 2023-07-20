package app

import (
	words_service "alexdenkk/words/internal/words/service"
	words_handler "alexdenkk/words/internal/words/gateway/rpc"
)

func (app *App) InitWordsService() {
	// service
	service := words_service.New()
	// handler
	handler := words_handler.New(service)
	app.WordsHandler = handler
}
