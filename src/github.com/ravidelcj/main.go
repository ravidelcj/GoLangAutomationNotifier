package main

import (

  "fmt"
  "database/sql"
	_ "github.com/go-sql-driver/mysql"
  "github.com/ravidelcj/scraper"
  "github.com/ravidelcj/database"
)

func main() {

      Database.initDatabase()
      go ScrapeNotice("http://ggsipuresults.nic.in/ipu/examnotice/examnoticemain.htm",db)
      go ScrapeResults("http://ggsipuresults.nic.in/ipu/results/resultsmain.htm", db)
      go ScrapeDatesheet("http://ggsipuresults.nic.in/ipu/datesheet/datesheetmain.htm", db)

      var endInput string
      fmt.Scanln(&endInput)
}
