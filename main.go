package main

import (
	"github.com/Charles-Leo-G/tss-lib/cmd"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6062", nil))
	}()

	cmd.Execute()
}
