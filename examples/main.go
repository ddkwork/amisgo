package main

import (
	"github.com/zrcoder/amisgo"
)

func main() {
	ServeApi()

	panic(amisgo.ListenAndServe(":8080", page))
}
