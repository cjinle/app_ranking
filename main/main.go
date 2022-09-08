package main

import (
	"flag"
	"fmt"

	appranking "github.com/cjinle/app_ranking"
)

func main() {
	// appranking.Request("game", "TH", "2022-09-03")
	dt := flag.String("d", "", "input datetime like 2022-09-03")
	flag.Parse()
	if len(*dt) == 0 {
		fmt.Println("param error", *dt)
		return
	}
	appranking.HandleRanking(*dt)
}
