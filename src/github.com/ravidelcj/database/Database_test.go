package database


import "fmt"
import "testing"



func Test(t *testing.T){

      success := initDatabase()

      if success == false {
        t.Error("Error in database Connection")
      }

      elem, err1 := getLastElement("datesheet_ipu")
      if err1 != nil {
        t.Error("nil returned from table")
      }
      fmt.Println(elem.title)
}
