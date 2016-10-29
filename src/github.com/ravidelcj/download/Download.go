package download

import(
   "os"
   "fmt"
   "net/http"
   "io"
)

func DownloadFile(title string, folder string, url string) (string, bool) {
      //file path
      path := "../../../../files/"+folder + "/" + title + ".pdf"

      out, err := os.Create(path)
      if err != nil {
          fmt.Println("Error creating path : " + path)
          return "", false
      }
      defer out.Close()

      resp, err := http.Get(url)
      if err != nil {
        fmt.Println("Error in connceting to : " + url)
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
