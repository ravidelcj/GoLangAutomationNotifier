package database

import(
  "errors"
  "fmt"
  "database/sql"
	_ "github.com/go-sql-driver/mysql"
  "github.com/ravidelcj/stack"
)

//Data type that defines one identity
type Element struct {
  Title string
  Date string
  Url string
  RemoteUrl string
}


//global database object for every package
var (
    Db *sql.DB
)

func InitDatabase() bool {
  var err error
  Db, err = sql.Open("mysql", "root:admin@/ipuscraper")
  if err != nil {
    fmt.Println("Error in database connection")
    return false
  }
  //defer db.Close()

  err = Db.Ping()
  if err != nil {
    fmt.Println("Error in database call")
    return false
  }
  return true
}

// params  elem : element to be inserted , folder : folderName
func InsertNoticeData( elem stack.Element, folder string) bool  {

    var err error
    var stmt *sql.Stmt
    switch folder {
          case "Results" : stmt, err = Db.Prepare("INSERT results_ipu SET title=?, date=?, url=?, remoteUrl=?")
          case "Notices" : stmt, err = Db.Prepare("INSERT notice_ipu SET title=?, date=?, url=?, remoteUrl=?")
          case "Datesheets" : stmt, err = Db.Prepare("INSERT datesheet_ipu SET title=?, date=?, url=?, remoteUrl=?")
    }

    if err != nil {
      fmt.Println("Error in preparing insert statement")
      return false
    }
    _, err1 := stmt.Exec(elem.Title, elem.Date, elem.Url, elem.RemoteUrl)
    if err1 != nil {
      fmt.Println("Error inserting in database ")
      return false
    }
    return true
}

func GetLastElement( table string) (Element, error)  {
    var elem Element
    err := Db.QueryRow("Select title, date, url from datesheet_ipu order by id DESC LIMIT 1").Scan(&elem.Title, &elem.Date, &elem.Url)

    if err != nil && err != sql.ErrNoRows {
        fmt.Println("Error in getting last element : ",err)
    }else {
        return elem, nil
    }

    return elem, errors.New("Statement gives no results")
}
