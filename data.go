package appranking

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const CAT = "game_card"
const COUNTRY = "TH"

type Connect struct {
	DB *sql.DB
}

var connect *Connect
var err error

func HandleRanking(date string) {
	fmt.Println("start handle date", date)
	connect = new(Connect)
	connect.DB, err = sql.Open("mysql", cfg.Database)
	if err != nil {
		panic(err)
	}
	appDatas, err := Request(CAT, COUNTRY, date)
	if err != nil {
		fmt.Println("HandleRanking", err)
		return
	}
	if len(appDatas) == 0 {
		return
	}
	CleanRank(0, date, COUNTRY, CAT)
	CleanRank(1, date, COUNTRY, CAT)
	CleanRank(2, date, COUNTRY, CAT)
	for _, v := range appDatas {
		if len(v) == 0 {
			continue
		}
		for rankType, vv := range v {
			app := &App{
				AppId:            vv.AppId,
				Name:             vv.Name,
				PublisherName:    vv.PublisherName,
				Icon:             vv.Icon,
				OS:               vv.OS,
				ReleaseDate:      vv.ReleaseDate,
				UpdatedDate:      vv.UpdatedDate,
				Rating:           vv.Rating,
				RatingCount:      vv.RatingCount,
				Price:            vv.Price,
				Version:          vv.Version,
				Downloads:        vv.Downloads,
				Revenue:          vv.Revenue,
				SupportUrl:       vv.SupportUrl,
				WebsiteUrl:       vv.WebsiteUrl,
				PrivacyPolicyUrl: vv.PrivacyPolicyUrl,
				FeatureGraphic:   vv.FeatureGraphic,
				ShortDescription: vv.ShortDescription,
			}
			app.Update()
			AddRank(int32(rankType), vv.Rank, date, COUNTRY, CAT, vv.AppId)
		}
	}
}

func CleanRank(rank_type int32, rank_date, country, cat string) error {
	tb := fmt.Sprintf("rank_%s_%s%d", country, cat, rank_type)
	query := fmt.Sprintf("DELETE FROM %s WHERE rank_type=? AND rank_date=? AND country=? AND cat=?", tb)
	_, err := connect.DB.Exec(query, rank_type, rank_date, country, cat)
	return err
}

func AddRank(rank_type, rank int32, rank_date, country, cat, app_id string) error {
	tb := fmt.Sprintf("rank_%s_%s%d", country, cat, rank_type)
	query := fmt.Sprintf("INSERT INTO %s SET rank_date=?,rank=?,app_id=?", tb)
	fmt.Println(rank_type, rank, rank_date, country, cat, app_id)
	_, err := connect.DB.Exec(query, rank_date, rank, app_id)
	return err
}

type App struct {
	AppId            string
	Name             string
	PublisherName    string
	Icon             string
	OS               string
	ReleaseDate      string
	UpdatedDate      string
	Rating           float32
	RatingCount      int32
	Price            float32
	Version          string
	Downloads        int32
	Revenue          int32
	SupportUrl       string
	WebsiteUrl       string
	PrivacyPolicyUrl string
	FeatureGraphic   string
	ShortDescription string
}

func (app *App) Update() error {
	setStr := "name=?,publisher_name=?,icon=?,os=?,release_date=?,updated_date=?,rating=?,rating_count=?,price=?,version=?," +
		"downloads=?,revenue=?,support_url=?,website_url=?,privacy_policy_url=?,feature_graphic=?,short_description=?"
	query := fmt.Sprintf("INSERT INTO app SET app_id=?,%s ON DUPLICATE KEY UPDATE %s", setStr, setStr)
	_, err := connect.DB.Exec(query, app.AppId,
		app.Name, app.PublisherName, app.Icon, app.OS, app.ReleaseDate, app.UpdatedDate,
		app.Rating, app.RatingCount, app.Price, app.Version, app.Downloads, app.Revenue, app.SupportUrl,
		app.WebsiteUrl, app.PrivacyPolicyUrl, app.FeatureGraphic, app.ShortDescription,
		app.Name, app.PublisherName, app.Icon, app.OS, app.ReleaseDate, app.UpdatedDate,
		app.Rating, app.RatingCount, app.Price, app.Version, app.Downloads, app.Revenue, app.SupportUrl,
		app.WebsiteUrl, app.PrivacyPolicyUrl, app.FeatureGraphic, app.ShortDescription)
	return err
}
