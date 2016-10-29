package scraper

import (
  "fmt",
  "github.com/PuerkitoBio/goquery"
  "github.com/ravidelcj/stack"
  "github.com/ravidelcj/download"
  "github.com/ravidelcj/database"
)

//ScrapeNotice from ggsipu
func ScrapeNotice(url string)  {

    baseUrl := "http://ggsipuresults.nic.in/ipu/examnotice/"

    doc, err := goquery.NewDucmnet(url)
    if err != nil {
        fmt.Println("Error in Document Connection NoticeScraper")
        continue
    }
    while true {
        var lastElem element
        lastElem = getLastElement()
        var stackElement StackNode

        //parsing document
        doc.Find("table tr").EachWithBreak( func (indexTr int, tr *goquery.Selection)  bool {
          var elem element
          if indexTr >= 1 {
            tr.Find("td").Each( func (indexTd int, td *goquery.Selectiom) {
              //for title
              if indexTd == 1 {
                elem.title = td.Text()
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
            if elem.title == lastElem.title {
              return false
            }else {
              stackElement.push(elem)
              return true
            }
          }//indexIf
        })//Find

        addStackToDatabase(stackElement, "Results")
    }
}

//DateSheet Scraper
func ScrapeDatesheet(url string)  {

    baseUrl := "http://ggsipuresults.nic.in/ipu/datesheet/"

    doc, err := goquery.NewDucmnet(url)
    if err != nil {
        fmt.Println("Error in Document Connection DatesheetScraper ")
        continue
    }
    while true {
        var lastElem element
        lastElem = getLastElement()
        var stackElement StackNode

        //parsing document
        doc.Find("table tr").EachWithBreak( func (indexTr int, tr *goquery.Selection)  bool {
          var elem element
          if indexTr >= 1 {
            tr.Find("td").Each( func (indexTd int, td *goquery.Selectiom) {
              //for title
              if indexTd == 1 {
                elem.title = td.Text()
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
            if elem.title == lastElem.title {
              return false
            }else {
              stackElement.push(elem)
              return true
            }
          }//indexIf
        })//Find

        addStackToDatabase(stackElement, "Datesheet")
    }
}


func ScrapeResults(url string)  {

    baseUrl := "http://ggsipuresults.nic.in/ipu/results/"

    doc, err := goquery.NewDucmnet(url)
    if err != nil {
        fmt.Println("Error in Document Connection ResultsScraper ")
        continue
    }
    while true {
        var lastElem element
        lastElem = getLastElement()
        var stackElement StackNode

        //parsing document
        doc.Find("table tr").EachWithBreak( func (indexTr int, tr *goquery.Selection)  bool {
          var elem element
          if indexTr >= 1 {
            tr.Find("td").Each( func (indexTd int, td *goquery.Selectiom) {
              //for title
              if indexTd == 1 {
                elem.title = td.Text()
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
            if elem.title == lastElem.title {
              return false
            }else {
              stackElement.push(elem)
              return true
            }
          }//indexIf
        })//Find

        addStackToDatabase(stackElement, "Datesheet")
    }
}

//TODO : add database variable to params
func addStackToDatabase(stack *StackNode, folder string)  {

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
