package scraper

import (
  "fmt",
  "github.com/PuerkitoBio/goquery"
  "github.com/ravidelcj/stack"
  "github.com/ravidelcj/download"
  "github.com/ravidelcj/database"
  "strings"
)

//ScrapeNotice from ggsipu
func ScrapeNotice(url string, db *DB)  {

    baseUrl := "http://ggsipuresults.nic.in/ipu/examnotice/"

    doc, err := goquery.NewDocument(url)
    if err != nil {
        fmt.Println("Error in Document Connection NoticeScraper")
        continue
    }
    while true {

        var lastElem element
        lastElem, err = getLastElement(db, "notice_ipu")
        if err != nil {
          fmt.Println("Error in getting last element notice_ipu")
          continue
        }
        var stackElement StackNode

        //parsing document
        doc.Find("table tr").EachWithBreak( func (indexTr int, tr *goquery.Selection)  bool {
          var elem element
          if indexTr >= 1 {
            tr.Find("td").Each( func (indexTd int, td *goquery.Selectiom) {
              //for title
              if indexTd == 1 {
                elem.title = strings.Trim(td.Text(), " ")
              }//titleIf

              //link to ggsipu server
              if indexTd == 2 {
                td.Find("a").Each(func(_ int, a *goquery.Selection){
                  url, exist := a.Attr("href")
                  if exist {
                    elem.url = baseUrl+url
                  } else {
                    fmt.Println("URL dont exist ")
                  }
                })
              }//index2
              //date
              if indexTd == 3 {
                  elem.date = td.Text()
              }

            })
            if elem.url == lastElem.url {
              return false
            }else {
              stackElement.push(elem)
              return true
            }
          }//indexIf
        })//Find

        addStackToDatabase(stackElement, "Results", db)
    }
}

//DateSheet Scraper
func ScrapeDatesheet(url string, db *DB)  {

    baseUrl := "http://ggsipuresults.nic.in/ipu/datesheet/"

    doc, err := goquery.NewDocument(url)
    if err != nil {
        fmt.Println("Error in Document Connection DatesheetScraper ")
        continue
    }
    while true {
        var lastElem element
        lastElem, err = getLastElement(db, "datesheet_ipu")
        if err != nil {
          fmt.Println("Error in getting last element notice_ipu")
          continue
        }
        var stackElement StackNode

        //parsing document
        doc.Find("table tr").EachWithBreak( func (indexTr int, tr *goquery.Selection)  bool {
          var elem element
          if indexTr >= 1 {
            tr.Find("td").Each( func (indexTd int, td *goquery.Selectiom) {
              //for title
              if indexTd == 1 {
                elem.title = strings.Trim(td.Text(), " ")
              }//titleIf

              //link to ggsipu server
              if indexTd == 2 {
                td.Find("a").Each(func(_ int, a *goquery.Selection){
                  url, exist := a.Attr("href")
                  if exist {
                    elem.url = baseUrl+url
                  } else {
                    fmt.Println("URL dont exist ")
                  }
                })
              }//index2
              //date
              if indexTd == 3 {
                  elem.date = td.Text()
              }

            })
            if elem.url == lastElem.url {
              return false
            }else {
              stackElement.push(elem)
              return true
            }
          }//indexIf
        })//Find

        addStackToDatabase(stackElement, "Datesheet", db)
    }
}


func ScrapeResults(url string, db *DB)  {

    baseUrl := "http://ggsipuresults.nic.in/ipu/results/"

    doc, err := goquery.NewDocument(url)
    if err != nil {
        fmt.Println("Error in Document Connection ResultsScraper ")
        continue
    }
    while true {
        var lastElem element
        lastElem, err = getLastElement(db, "results_ipu")
        if err != nil {
          fmt.Println("Error in getting last element notice_ipu")
          continue
        }
        var stackElement StackNode

        //parsing document
        doc.Find("table tr").EachWithBreak( func (indexTr int, tr *goquery.Selection)  bool {
          var elem element
          if indexTr >= 1 {
            tr.Find("td").Each( func (indexTd int, td *goquery.Selectiom) {
              //for title
              if indexTd == 1 {
                elem.title = strings.Trim(td.Text(), " ")
              }//titleIf

              //link to ggsipu server
              if indexTd == 3 {
                td.Find("a").Each(func(_ int, a *goquery.Selection){
                  url, exist := a.Attr("href")
                  if exist {
                    elem.url = baseUrl+url
                  } else {
                    fmt.Println("URL dont exist ")
                  }
                })
              }//index2
              //date
              if indexTd == 4 {
                  elem.date = td.Text()
              }

            })
            if elem.url == lastElem.url {
              return false
            }else {
              stackElement.push(elem)
              return true
            }
          }//indexIf
        })//Find

        addStackToDatabase(stackElement, "Datesheet", db)
    }
}

func addStackToDatabase(stack *StackNode, folder string, db *DB)  {

    while !stack.isEmpty() {
        elem := stack.top()
        path, success := DownloadFile(elem.title, folder, elem.url)
        if success == false {
          //if error try again
          fmt.Println("Error in downloading file")
          continue
        }else {
          elem.remoteUrl = path
          //TODO : database variavle db to be added
          insertNoticeData(db, elem, folder)
          stack.pop()
        }
    }
}
