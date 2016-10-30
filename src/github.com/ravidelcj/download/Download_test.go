package download

import "testing"

func Test(t *testing.T)  {

      url := "http://ggsipuresults.nic.in/ipu/results/2016/250116/016_B.Arch._SUPPLEMENTARY_september_2015_Rechecking_Result_With%20NO%20Change.pdf"

      _, err := DownloadFile("a", "Results", url)
      if err == false {
        t.Error("Error")
      }

      str := "RAvi prasad / ravi   "
      title := trimAndModify(str)
      if title != "RAvi_prasad___ravi" {
        t.Error("String not parsed")
      }
}
