package database

import(
  "errors"
  "fmt"
  "database/sql"
	_ "github.com/go-sql-driver/mysql"
)

//Data type that defines one identity
type element struct {
  title string
  date string
  url string
  remoteUrl string
}


//global database object for every package
var (
    db *sql.DB
)

func initDatabase() bool {
  var err error
  db, err = sql.Open("mysql", "root:admin@/ipuscraper")
  if err != nil {
    fmt.Println("Error in database connection")
    return false
  }
  //defer db.Close()

  err = db.Ping()
  if err != nil {
    fmt.Println("Error in database call")
    return false
  }
  return true
}

// params db : database variable , elem : element to be inserted , folder : folderName
func insertNoticeData( elem element, folder string) bool  {

    var err error
    var stmt *sql.Stmt
    switch folder {
    case "Results" : stmt, err = db.Prepare("INSERT results_ipu SET title=?, date=?, url=?, remoteUrl=?")
    case "Notices" : stmt, err = db.Prepare("INSERT notice_ipu SET title=?, date=?, url=?, remoteUrl=?")
    case "Datesheets" : stmt, err = db.Prepare("INSERT datesheet_ipu SET title=?, date=?, url=?, remoteUrl=?")
    }
    if err != nil {
      fmt.Println("Error in preparing insert statement")
      return false
    }
    _, err1 := stmt.Exec(elem.title, elem.date, elem.url, elem.remoteUrl)
    if err1 != nil {
      fmt.Println("Error inserting in database ")
      return false
    }
    return true
}

func getLastElement( table string) (element, error)  {
    var elem element
  //    var str string
    err := db.QueryRow("Select title, date, url from datesheet_ipu order by id DESC LIMIT 1").Scan(&elem.title, &elem.date, &elem.url)
    // if err != nil {
    //      fmt.Println("Error in retreiving last element from " + table)
    //      return elem, errors.New("error in retreiving")
    // }
    //defer row.Close()
    // for row.Next() {

        if err != nil && err != sql.ErrNoRows {
            fmt.Println("Error in getting last element : ",err)
        }else {
           return elem, nil
        }
    // }
    return elem, errors.New("Statement gives no results")
}
