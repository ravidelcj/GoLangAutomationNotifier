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
    // table, _ := "github.com/PuerkitoBio/goquery";
    doc, err := goquery.NewDucmnet(url)
    while true {
        var lastElem element
        lastElem = getLastElement()
        doc, err := goquery.NewDocument("http://ggsipuresults.nic.in/ipu/datesheet/datesheetmain.htm")
        if err != nil {
            fmt.Println("Error in Document Connection NoticeScraper line 16 ")
            continue
        }

        var stackElement StackNode

        //parsing document
        doc.Find("table tr").EachWithBreak( func (indexTr int, tr *goquery.Selection)  bool {
          var elem element
          if indexTr >= 1 {
            tr.Find("td").EachWithBreak( func (indexTd int, td *goquery.Selectiom) bool {
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
              /*TODO add to stack*/
              stackElement.push(elem)
              return true
            }
          }//indexIf
        })//Find

        addStackToDatabase(stackElement)
    }
}

//TODO : add database variable to params
func addStackToDatabase(stack *StackNode, folder string)  {

    while !stack.isEmpty() {
        elem := stack.top()
        path, success := DownloadFile(elem.title, folder, elem.url)
        if success == false {
          //if error try again
          fmt.Println("Error in downloading file at line 75 NoticeScraper")
          continue
        }else {
          elem.remoteUrl = path
          insertNoticeData(db, elem, folder)
          stack.pop()
        }
    }
}
