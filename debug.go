package main

import (
	"fmt"
	"net/http"

	"encoding/xml"
	// "github.com/gorilla/mux"
	"io/ioutil"
	// 	"os"
)

const port = "8080"

type Channel struct {
	// Name string `xml:"name"`
	// Item []Item `xml:"item"`
	Title string `xml:"title"`
}

func main() {
	http.HandleFunc("/", handler)
}

func top_page() {

	url := "https://omocoro.jp/feed"

	resp, _ := http.Get(url)
	defer resp.Body.Close()

	xmlStr, _ := ioutil.ReadAll(resp.Body)
	data := Channel{}
	if err := xml.Unmarshal([]byte(xmlStr), data); err != nil {
		fmt.Println("XML Unmarshal error:", err)
		return
	}
	// fmt.Fprint(w, data.Title) // htmlをstringで取得
	// fmt.Fprint(w, "hello,golang.")
	// fmt.Println("XML Unmarshal error:", err)
}
