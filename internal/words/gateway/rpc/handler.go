package rpc

import (
	"alexdenkk/words/internal/words"
)

type RpcHandler struct {
	Service words.Service
}

func New(service words.Service) *RpcHandler {
	return &RpcHandler{
		Service: service,
	}
}
