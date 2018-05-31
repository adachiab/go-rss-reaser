package main

import (
  "net/http"
  "io/ioutil"
  "fmt"
  "os"
  "encoding/xml"
)


type Group struct {
    Name      string    `xml:"name"`
    Companies []Company `xml:"company"`
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "hello,golang.")
}
func trumpHandler(w http.ResponseWriter, r *http.Request) {
    url := "https://omocoro.jp/feed"

  resp, _ := http.Get(url)
  defer resp.Body.Close()

  xmlStr, _ := ioutil.ReadAll(resp.Body)
  data := new(Group)
  if err := xml.Unmarshal([]byte(xmlStr), data); err != nil {
      fmt.Println("XML Unmarshal error:", err)
      return
  }
  fmt.Fprint(w, string(data)) // htmlをstringで取得
}

func main() {
    http.HandleFunc("/", viewHandler)
    http.HandleFunc("/trump", trumpHandler)
    http.ListenAndServe(os.Getenv("IP")+":"+os.Getenv("PORT"), nil)
}