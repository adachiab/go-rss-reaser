package main

import (
	"encoding/xml"
	//"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Rss2 struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	// Required
	Title       string `xml:"channel>title"`
	Link        string `xml:"channel>link"`
	Description string `xml:"channel>description"`
	// Optional
	PubDate  string `xml:"channel>pubDate"`
	ItemList []Item `xml:"channel>item"`
}

type Item struct {
	// Required
	Title       string        `xml:"title"`
	Link        string        `xml:"link"`
	Description template.HTML `xml:"description"`
	// Optional
	Content  template.HTML `xml:"encoded"`
	PubDate  string        `xml:"pubDate"`
	Comments string        `xml:"comments"`
}

type Atom1 struct {
	XMLName   xml.Name `xml:"http://www.w3.org/2005/Atom feed"`
	Title     string   `xml:"title"`
	Subtitle  string   `xml:"subtitle"`
	Id        string   `xml:"id"`
	Updated   string   `xml:"updated"`
	Rights    string   `xml:"rights"`
	Link      Link     `xml:"link"`
	Author    Author   `xml:"author"`
	EntryList []Entry  `xml:"entry"`
}

type Link struct {
	Href string `xml:"href,attr"`
}

type Author struct {
	Name  string `xml:"name"`
	Email string `xml:"email"`
}

type Entry struct {
	Title   string `xml:"title"`
	Summary string `xml:"summary"`
	Content string `xml:"content"`
	Id      string `xml:"id"`
	Updated string `xml:"updated"`
	Link    Link   `xml:"link"`
	Author  Author `xml:"author"`
}

type Page struct {
	Title string
	Count int
}

func parseFeedContent(content []byte) (Rss2, bool) {
	v := Rss2{}
	err := xml.Unmarshal(content, &v)
	if err != nil {

		log.Println(err)
		return v, false
	}

	if v.Version == "2.0" {
		// RSS 2.0
		for i, _ := range v.ItemList {
			if v.ItemList[i].Content != "" {
				v.ItemList[i].Description = v.ItemList[i].Content
			}
		}
		return v, true
	}

	log.Println("not RSS 2.0")
	return v, false
}

func TopHandler(w http.ResponseWriter, r *http.Request) {
	// parse sample rss feed
	// xmlContent1, _ := ioutil.ReadFile("feed.xml")
	url := "https://omocoro.jp/feed"
	resp, _ := http.Get(url)
	defer resp.Body.Close()

	xmlContent1, _ := ioutil.ReadAll(resp.Body)
	r1, ok1 := parseFeedContent(xmlContent1)
	if ok1 {
		// log.Println(r1.Title)
		//fmt.Fprint(w, r1.Title) // htmlをstringで取得
	} else {
		log.Println("fail to read feed")
	}

	//rendering

	page := Page{r1.Title, 1}
	tmpl, err := template.ParseFiles("template/index.html") // ParseFilesを使う
	if err != nil {
		panic(err)
	}
	tmpl.Execute(w, page)
}

func main() {
	http.HandleFunc("/", TopHandler)
	http.ListenAndServe(os.Getenv("IP")+":"+os.Getenv("PORT"), nil)
}
