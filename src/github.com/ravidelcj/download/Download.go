package download

import(
   "os"
   "fmt"
   "net/http"
   "io"
   "strings"
)


//Download the file and save it with the name title
//@params titleC : title of the file , folder : which folder to save the file in , url : the url to be downloaded
//@return returns the local path of the file which is downloaded
func DownloadFile(titleC string, folder string, url string) (string, bool) {
      //file path
      title := trimAndModify(titleC)
      path := "/home/ezio/Programming/TechbitsScrapper/files/"+folder + "/" + title + ".pdf"

      out, err := os.Create(path)
      if err != nil {
          fmt.Println("Error creating path : " + path)
          return "", false
      }
      defer out.Close()

      resp, err := http.Get(url)
      if err != nil {
        fmt.Println("Error in connecting to : " , err)
        return "", false
      }
      defer resp.Body.Close()

      _, err = io.Copy(out, resp.Body)
      if err != nil {
        fmt.Println("Error in downlading : " + url)
        return "", false;
      }

      fmt.Println("Download successful : " + url)
      return path, true
}


func trimAndModify(str string) string  {

  str = strings.Trim(str, " ")
  str = strings.Replace(str, " ", "_", -1)
  str = strings.Replace(str, "/", "_", -1)
  str = strings.Replace(str, "\n", "_", -1)
  return str
}
