package scraper

import (

  "testing"
  "fmt"
  "github.com/ravidelcj/database"
)

func test(t *testing.T)  {

  fmt.Println("Asfasf")
  database.InitDatabase()
  success := ScrapeNotice("http://ggsipuresults.nic.in/ipu/examnotice/examnoticemain.htm")
  if ! success {
    t.Error("failed")
  }

  fmt.Println(success)
  //Output:
  //true

}
