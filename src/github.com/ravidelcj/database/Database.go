package database

import(
  "database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// params db : database variable , elem : element to be inserted , folder : folderName
func insertNoticeData(db *DB, elem element, folder string) bool  {

    switch folder {
    case "Results":stmt, err := db.Prepare("INSERT results_ipu SET title=?, date=?, url=?, remoteUrl=?")
    case "Notices":stmt, err := db.Prepare("INSERT notice_ipu SET title=?, date=?, url=?, remoteUrl=?")
    case "Datesheets":stmt, err := db.Prepare("INSERT datesheet_ipu SET title=?, date=?, url=?, remoteUrl=?")
    }
    if err != nil {
      fmt.Println("Error in preparing statement Database.go")
      return false
    }

    res, err1 := stmt.Exec(elem.title, elem.date, elem.url, elem.remoteUrl)
    if err1 != nil {
      fmt.Println("Error inserting in database ")
      return false
    }
    return true
}


/*TODO add id to database*/
func getLastELemet(db *Db, table string) element  {
  row, err := db.Query("Select * from "+ table + "order by id DESC LIMIT 1")
  if err != nil {
    fmt.Println("Error in retreiving last element from " + table)
    return nil
  }
  defer row.Close()
  var elem element
  for row.Next() {
    var id int
    err := row.Scan(&id, &elem.title, &elem.date, &elem.url, &elem.remoteUrl)
      if err != nil {
        fmt.Println("Error in getting last element")
      }else {
        return elem
      }
  }
}
