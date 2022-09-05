package main

import (
	appranking "github.com/cjinle/app_ranking"
)

func main() {
	// appranking.Request("game", "TH", "2022-09-03")
	appranking.HandleRanking("2022-09-03")
}
