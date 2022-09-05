package appranking

type Cfg struct {
	Database string
}

var cfg *Cfg

func init() {
	cfg = &Cfg{Database: "root:123456@tcp(127.0.0.1:3306)/app_ranking"}
}
