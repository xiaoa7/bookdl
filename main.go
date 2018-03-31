package main

import (
	"util"
)

func init() {
	util.LoadDownloadCfg("./config.json")
}
func main() {
	util.DownloadBook()
}
