package main

import (

  "fmt"
  "database/sql"
	_ "github.com/go-sql-driver/mysql"
  "github.com/ravidelcj/scraper"
)

func main() {

      //Database call
      db, err := sql.Open("mysql", "root:admin@/ipuscraper")
      if err != nil {
        fmt.Println("Error in database connection")
        return
      }
      defer db.Close()

      err = db.Ping()
      if err != nil {
        fmt.Println("Error in database call")
        return
      }

      go ScrapeNotice("http://ggsipuresults.nic.in/ipu/examnotice/examnoticemain.htm",db)
      go ScrapeResults("http://ggsipuresults.nic.in/ipu/results/resultsmain.htm", db)
      go ScrapeDatesheet("http://ggsipuresults.nic.in/ipu/datesheet/datesheetmain.htm", db)


      var endInput string
      fmt.Scanln(&endInput)
}
