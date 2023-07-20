package app

import (
	words_handler "alexdenkk/words/internal/words/gateway/rpc"
	"net/rpc"
	"net/http"
)

type App struct {
	Address      string  
	WordsHandler *words_handler.RpcHandler
}

func (app *App) Run() error {
	rpc.Register(app.WordsHandler)
	rpc.HandleHTTP()

	return http.ListenAndServe(app.Address, nil)
}

func New(addr string) *App {
	app := &App{
		Address: addr,
	}

	app.InitWordsService()
	return app
}
