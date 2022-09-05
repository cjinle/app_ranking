package appranking

type Cfg struct {
	Database string
}

var cfg *Cfg

func init() {
	cfg = &Cfg{Database: "app_ranking"}
}
