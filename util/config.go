package util

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

//配置
type downloadconfig struct {
	Encoding              string            `json:"encoding"`         //编码方式
	ListUrl               string            `json:"listurl"`          //列表页地址
	ListSelector          string            `json:"listselector"`     //列表中文章匹配规则
	ContentUrlPrefix      string            `json:"contenturlprefix"` //内容页地址前缀
	ContentSelector       string            `json:"contentselector"`  //内容匹配规则
	BookName              string            `json:"bookname"`
	SingleFileChapterSize int               `json:"singlefilechaptersize"` //单文件章节数
	RequestHeader         map[string]string `json:"requestheader"`
}

var DownloadCfg downloadconfig

//
func LoadDownloadCfg(cfg_file string) {
	fi, _ := os.Open(cfg_file)
	bs, _ := ioutil.ReadAll(fi)
	fi.Close()
	json.Unmarshal(bs, &DownloadCfg)
}
