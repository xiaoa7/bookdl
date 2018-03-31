package util

import (
	"bytes"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/axgle/mahonia"
)

//
var client *http.Client = &http.Client{
	Transport: &http.Transport{
		Dial: func(netw, addr string) (net.Conn, error) {
			c, err := net.DialTimeout(netw, addr, time.Second*5)
			if err != nil {
				return nil, err
			}
			return c, nil

		},
		MaxIdleConnsPerHost:   10,
		ResponseHeaderTimeout: time.Second * 4,
	},
}

//下载指定内容
func Download(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, bytes.NewReader([]byte{}))
	if err != nil {
		return nil, err
	}
	for k, v := range DownloadCfg.RequestHeader {
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if strings.ToLower(DownloadCfg.Encoding) == "gbk" {
		dec := mahonia.NewDecoder("gbk")
		rd := dec.NewReader(resp.Body)
		bs, err := ioutil.ReadAll(rd)
		if err != nil {
			return nil, err
		}
		resp.Body.Close()
		return bs, nil
	} else {
		bs, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		resp.Body.Close()
		return bs, nil
	}
}
