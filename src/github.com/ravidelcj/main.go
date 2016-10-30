package main

import (

  "fmt"
  "github.com/ravidelcj/scraper"
  "github.com/ravidelcj/database"
)

func main() {

      database.InitDatabase()
      go scraper.ScrapeNotice("http://ggsipuresults.nic.in/ipu/examnotice/examnoticemain.htm")
      go scraper.ScrapeResults("http://ggsipuresults.nic.in/ipu/results/resultsmain.htm")
      go scraper.ScrapeDatesheet("http://ggsipuresults.nic.in/ipu/datesheet/datesheetmain.htm")

      var endInput string
      fmt.Scanln(&endInput)
}
