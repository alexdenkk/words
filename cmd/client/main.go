package main

import (
	"fmt"
	"log"
	"net/rpc"
	"alexdenkk/words/model"
)

type Arguments struct {
	Text string
}

func main() {
	client, err := rpc.DialHTTP("tcp", ":8080")
	if err != nil {
		log.Fatal(err.Error())
	}
	arg := Arguments{Text: "Well, whatever happens, happens..."}
	var resp model.Result

	err = client.Call("RpcHandler.Count", arg, &resp)
	fmt.Println(resp.WordsTotal)
	fmt.Println(resp.WordsCount)
}
