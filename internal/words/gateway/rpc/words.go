package rpc

import (
	"context"
	"alexdenkk/words/model"
)

type Arguments struct {
	Text string
}

func (h *RpcHandler) Count(r Arguments, resp *model.Result) error {
	*resp = h.Service.CountWords(context.Background(), r.Text)

	return nil
}
