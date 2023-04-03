package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"net/http"
	"regexp"
)

var headerRe = regexp.MustCompile(`<div>[\s\S]*?<div class="index_carousel_img__HbOWM"[\s\S]*?<a.*?alt="([\s\S]*?)"/>`)

func main() {
	url := "https://www.thepaper.cn/"
	body, err := Fetch(url)
	if err != nil {
		fmt.Println("read content failed:", err)
		return
	}

	matches := headerRe.FindAllSubmatch(body, -1)
	for _, m := range matches {
		fmt.Println("fetch card news:", string(m[1]))
	}
}

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error status code:", resp.StatusCode)
	}

	bodyReader := bufio.NewReader(resp.Body) //bufio.Reader
	e := DetermineEncoding(bodyReader)
	// 编码转换，将特定编码转为utf-8编码
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return io.ReadAll(utf8Reader)
}

// DetermineEncoding 猜测编码类型
func DetermineEncoding(r *bufio.Reader) encoding.Encoding {
	bytesBuf, err := r.Peek(1024)
	if err != nil {
		fmt.Println("fetch error:", err)
		return unicode.UTF8
	}

	e, name, certain := charset.DetermineEncoding(bytesBuf, "")
	fmt.Printf("编码类型为：%s\n是否确定：%v\n", name, certain)
	return e
}
