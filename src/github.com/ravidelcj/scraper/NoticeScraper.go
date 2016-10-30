package scraper

import (
  "fmt"
  "github.com/PuerkitoBio/goquery"
  "github.com/ravidelcj/stack"
  "github.com/ravidelcj/download"
  "github.com/ravidelcj/database"
  "strings"
)


type Element struct {
  Title string
  Date string
  Url string
  RemoteUrl string
}

//ScrapeNotice from ggsipu
func ScrapeNotice(url string)  {

    baseUrl := "http://ggsipuresults.nic.in/ipu/examnotice/"

    doc, err := goquery.NewDocument(url)
    if err != nil {
        fmt.Println("Error in Document Connection NoticeScraper")
        return
    }
    for true {


        lastElem, err1 := database.GetLastElement("notice_ipu")
        if err1 != nil {
          fmt.Println("Error in getting last element notice_ipu")
          continue
        }
       stackElement := stack.NewStack()

        //parsing document
        doc.Find("table tr").EachWithBreak( func (indexTr int, tr *goquery.Selection)  bool {
          var elem stack.Element
          if indexTr >= 1 {
            tr.Find("td").Each( func (indexTd int, td *goquery.Selection) {
              //for title
              if indexTd == 1 {
                elem.Title = strings.Trim(td.Text(), " ")
              }//titleIf

              //link to ggsipu server
              if indexTd == 2 {
                td.Find("a").Each(func(_ int, a *goquery.Selection){
                  url, exist := a.Attr("href")
                  if exist {
                    elem.Url = baseUrl+url
                  } else {
                    fmt.Println("URL dont exist ")
                  }
                })
              }//index2
              //date
              if indexTd == 3 {
                  elem.Date = td.Text()
              }

            })
            if elem.Url == lastElem.Url {
              return false
            }else {
              stackElement.Push(elem)
              return true
            }
          }//indexIf
          return true
        })//Find

        addStackToDatabase(stackElement, "Results")
    }
}

//DateSheet Scraper
func ScrapeDatesheet(url string)  {

    baseUrl := "http://ggsipuresults.nic.in/ipu/datesheet/"

    doc, err := goquery.NewDocument(url)
    if err != nil {
        fmt.Println("Error in Document Connection DatesheetScraper ")
        return
    }
    for true {
        lastElem, err1 := database.GetLastElement("datesheet_ipu")
        if err1 != nil {
          fmt.Println("Error in getting last element notice_ipu")
          continue
        }
        stackElement := stack.NewStack()

        //parsing document
        doc.Find("table tr").EachWithBreak( func (indexTr int, tr *goquery.Selection)  bool {
          var elem stack.Element
          if indexTr >= 1 {
            tr.Find("td").Each( func (indexTd int, td *goquery.Selection) {
              //for title
              if indexTd == 1 {
                elem.Title = strings.Trim(td.Text(), " ")
              }//titleIf

              //link to ggsipu server
              if indexTd == 2 {
                td.Find("a").Each(func(_ int, a *goquery.Selection){
                  url, exist := a.Attr("href")
                  if exist {
                    elem.Url = baseUrl+url
                  } else {
                    fmt.Println("URL dont exist ")
                  }
                })
              }//index2
              //date
              if indexTd == 3 {
                  elem.Date = td.Text()
              }

            })
            if elem.Url == lastElem.Url {
              return false
            }else {
              stackElement.Push(elem)
              return true
            }
          }//indexIf
          return true
        })//Find

        addStackToDatabase(stackElement, "Datesheet")
    }
}


func ScrapeResults(url string)  {

    baseUrl := "http://ggsipuresults.nic.in/ipu/results/"

    doc, err := goquery.NewDocument(url)
    if err != nil {
        fmt.Println("Error in Document Connection ResultsScraper ")
        return
    }
    for true {
        //var lastElem Element
        lastElem, err1 := database.GetLastElement("results_ipu")
        if err1 != nil {
          fmt.Println("Error in getting last element notice_ipu")
          continue
        }
        stackElement := stack.NewStack()

        //parsing document
        doc.Find("table tr").EachWithBreak( func (indexTr int, tr *goquery.Selection)  bool {
          var elem stack.Element
          if indexTr >= 1 {
            tr.Find("td").Each( func (indexTd int, td *goquery.Selection) {
              //for title
              if indexTd == 1 {
                elem.Title = strings.Trim(td.Text(), " ")
              }//titleIf

              //link to ggsipu server
              if indexTd == 3 {
                td.Find("a").Each(func(_ int, a *goquery.Selection){
                  url, exist := a.Attr("href")
                  if exist {
                    elem.Url = baseUrl+url
                  } else {
                    fmt.Println("URL dont exist ")
                  }
                })
              }//index2
              //date
              if indexTd == 4 {
                  elem.Date = td.Text()
              }

            })
            if elem.Url == lastElem.Url {
              return false
            }else {
              stackElement.Push(elem)
              return true
            }
          }//indexIf
          return true
        })//Find

        addStackToDatabase(stackElement, "Datesheet")
    }
}

func addStackToDatabase(stack *stack.StackNode, folder string)  {

    for !stack.IsEmpty() {
        elem, _ := stack.Top()
        path, success := download.DownloadFile(elem.Title, folder, elem.Url)
        if success == false {
          //if error try again
          fmt.Println("Error in downloading file")
          continue
        }else {
          elem.RemoteUrl = path
          database.InsertNoticeData(elem, folder)
          stack.Pop()
        }
    }
}
