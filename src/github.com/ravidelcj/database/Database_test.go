package database


import "fmt"
import "testing"



func Test(t *testing.T){

      success := InitDatabase()

      if success == false {
        t.Error("Error in database Connection")
      }

      elem, err1 := GetLastElement("datesheet_ipu")
      if err1 != nil {
        t.Error("nil returned from table")
      }
      fmt.Println(elem.Title)
}
