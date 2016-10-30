package database

import(
  "errors"
  "fmt"
  "database/sql"
	_ "github.com/go-sql-driver/mysql"
  "github.com/ravidelcj/models"
)

//global database object for every package
var (
    Db *sql.DB
)

//initialize global datbase
func InitDatabase() bool {
  var err error
  Db, err = sql.Open("mysql", "root:admin@/ipuscraper")
  if err != nil {
    fmt.Println("Error in database connection")
    return false
  }
  err = Db.Ping()
  if err != nil {
    fmt.Println("Error in database call")
    return false
  }
  return true
}

//Inserts and element to the required table in the Database
//@params  elem : element to be inserted , folder : folderName
func InsertNoticeData( elem models.Element, folder string) bool  {

    var err error
    var stmt *sql.Stmt
    switch folder {
          case "Results" : stmt, err = Db.Prepare("INSERT results_ipu SET title=?, date=?, url=?, remoteUrl=?")
          case "Notice" : stmt, err = Db.Prepare("INSERT notice_ipu SET title=?, date=?, url=?, remoteUrl=?")
          case "Datesheet" : stmt, err = Db.Prepare("INSERT datesheet_ipu SET title=?, date=?, url=?, remoteUrl=?")
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


//Return the last element in the database
//@params table : table name
func GetLastElement( table string) (models.Element, error)  {
    var elem models.Element
    err := Db.QueryRow("Select title, date, url from " + table + " order by id DESC LIMIT 1").Scan(&elem.Title, &elem.Date, &elem.Url)

    if err != nil && err != sql.ErrNoRows {
        fmt.Println("Error in getting last element : ",err)
    }else {
        return elem, nil
    }

    return elem, errors.New("Statement gives no results")
}
