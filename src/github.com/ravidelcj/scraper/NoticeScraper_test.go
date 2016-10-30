package scraper

import (

  "testing"
  //"fmt"
  "github.com/ravidelcj/database"
)

func test(t *testing.T)  {

  database.InitDatabase()
  defer database.Db.Close()
  ScrapeNotice("http://ggsipuresults.nic.in/ipu/examnotice/examnoticemain.htm")


}
