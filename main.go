package main

import (
	"github.com/MakeFang/GoUtility/interactor"
	"github.com/MakeFang/GoUtility/slackrtm"
	"github.com/MakeFang/GoUtility/sqldb"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

func main() {

	db := sqldb.SetupDB()
	interactor.SetDB(db)
	defer db.Close()

	botToken := os.Getenv("BOT_OAUTH_ACCESS_TOKEN")
	slackClient := slackrtm.CreateSlackClient(botToken)
	slackrtm.GetIncomingMsg(slackClient)

}
