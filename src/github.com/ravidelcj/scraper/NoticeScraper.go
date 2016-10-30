package scraper

import (
  "fmt"
  "github.com/PuerkitoBio/goquery"
  "github.com/ravidelcj/models"
  "github.com/ravidelcj/download"
  "github.com/ravidelcj/database"
  "github.com/golang-collections/collections/stack"
  "strings"
)

//ScrapeNotice from ggsipu
//This func is responsible for scrapping Notice section of ggsipuresults
//The scraping will always be on and whenever there is an update, the it will download the file and save it
//in the databse
func ScrapeNotice(url string) bool {

    //base url for the downloaded files
    baseUrl := "http://ggsipuresults.nic.in/ipu/examnotice/"
    doc, err := goquery.NewDocument(url)
    if err != nil {
        fmt.Println("Error in Document Connection NoticeScraper")
        return false
    }
    //forever running loop which keeps on checking for update and whenever there is an update
    //it will download it and save details in database
    for true {

        //Get the lastElement in the datbase to check whether there is an update or not
        lastElem, err1 := database.GetLastElement("notice_ipu")

        if err1 != nil {
          fmt.Println("Error in getting last element notice_ipu")
          continue
        }

       stackElement := stack.New()
        //scraping the page
        //parsing document
        doc.Find("table tr").EachWithBreak( func (indexTr int, tr *goquery.Selection)  bool {
          var elem models.Element
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

        success := addStackToDatabase(stackElement, "Notice")
        if !success {
          return false
        }
    }
    return true
}

//DateSheet Scraper
func ScrapeDatesheet(url string) bool {

    baseUrl := "http://ggsipuresults.nic.in/ipu/datesheet/"

    doc, err := goquery.NewDocument(url)
    if err != nil {
        fmt.Println("Error in Document Connection DatesheetScraper ")
        return false
    }
    for true {
        lastElem, err1 := database.GetLastElement("datesheet_ipu")
        if err1 != nil {
          fmt.Println("Error in getting last element notice_ipu")
          continue
        }
        stackElement := stack.New()
        //parsing document
        doc.Find("table tr").EachWithBreak( func (indexTr int, tr *goquery.Selection)  bool {
          var elem models.Element
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

        success := addStackToDatabase(stackElement, "Datesheet")
        if !success {
          return false
        }
    }
    return true
}


func ScrapeResults(url string)  bool {

    baseUrl := "http://ggsipuresults.nic.in/ipu/results/"

    doc, err := goquery.NewDocument(url)
    if err != nil {
        fmt.Println("Error in Document Connection ResultsScraper : ", err)
        return false
    }
    for true {
        lastElem, err1 := database.GetLastElement("results_ipu")
        if err1 != nil {
          fmt.Println("Error in getting last element notice_ipu")
          continue
        }
        stackElement := stack.New()

        //parsing document
        doc.Find("table").Each( func (indexTable int, table *goquery.Selection)  {
          if(indexTable == 1){

            table.Find("tr").EachWithBreak( func (indexTr int, tr *goquery.Selection)  bool {
              var elem models.Element
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
                  }//index3
                  //date
                  if indexTd == 4 {
                      elem.Date = td.Text()
                  }

                })
                //fmt.Println(elem)
                if elem.Url == lastElem.Url {
                  return false
                }else{
                  stackElement.Push(elem)
                  return true
                }
              }//indexIf
              return true
            })//Find

          }
        })
        success := addStackToDatabase(stackElement, "Results")

        if !success {
          return false
        }
    }
    return true
}


//Add the stack to the mysql database and download the required file
func addStackToDatabase(st *stack.Stack, folder string) bool  {

    for st.Len() != 0 {
        elem := st.Peek().(models.Element)
        fmt.Println(elem)
        path, success := download.DownloadFile(elem.Title, folder, elem.Url)
        if success == false {
          //if error try again
          fmt.Println("Error in downloading file")
          return false
        }else {
          elem.RemoteUrl = path
          database.InsertNoticeData(elem, folder)
        }
        st.Pop()
    }
    return true
}
