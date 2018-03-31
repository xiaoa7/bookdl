package util

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
)
import (
	gq "github.com/PuerkitoBio/goquery"
)

//
type Item struct {
	Index int
	Title string
	Href  string
}

const (
	BOOK_DIR = "./mybooks"
)

//
func downloadlist() []*Item {
	bs, _ := Download(DownloadCfg.ListUrl)
	doc, _ := gq.NewDocumentFromReader(bytes.NewReader(bs))
	ret := make([]*Item, 0, 0)
	doc.Find(DownloadCfg.ListSelector).Each(func(index int, gs *gq.Selection) {
		title := gs.AttrOr("title", "")
		text := gs.Text()
		href := gs.AttrOr("href", "")
		if title == "" {
			title = text
		}
		ret = append(ret, &Item{index, title, href})
	})
	return ret
}

//
func downcontent(i *Item) string {
	bs, _ := Download(DownloadCfg.ContentUrlPrefix + i.Href)
	doc, _ := gq.NewDocumentFromReader(bytes.NewReader(bs))
	html, _ := doc.Find(DownloadCfg.ContentSelector).Html()
	return Clean(html)
}

//存入文件
func DownloadBook() {
	//创建目录
	if _, err := os.Stat(BOOK_DIR); err != nil {
		os.Mkdir(BOOK_DIR, 0777)
	}
	pos, index := 0, 1
	fi, _ := os.OpenFile(fmt.Sprintf("%s%c%s%d.html", BOOK_DIR, filepath.Separator, DownloadCfg.BookName, index), os.O_CREATE|os.O_TRUNC|os.O_SYNC|os.O_RDWR, 0777)
	for _, v := range downloadlist() {
		if pos == DownloadCfg.SingleFileChapterSize {
			pos = 0
			index = index + 1
			fi.Close()
			fi, _ = os.OpenFile(fmt.Sprintf("%s%c%s%d.html", BOOK_DIR, filepath.Separator, DownloadCfg.BookName, index), os.O_CREATE|os.O_TRUNC|os.O_SYNC|os.O_RDWR, 0777)
		}
		content := downcontent(v)
		fmt.Print(".")
		fmt.Fprint(fi, "<h3>"+v.Title+"</h3>"+CRLF, content)
		pos = pos + 1
	}
	if pos > 0 {
		fi.Close()
	}
}
