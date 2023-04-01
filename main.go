package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	url := "https://www.thepaper.cn/"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("fetch url error:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error status code:", resp.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read content failed:", err)
		return
	}

	numLinks := strings.Count(string(body), "<a")
	fmt.Printf("homepage has %d links!\n", numLinks)

	numLinks = bytes.Count(body, []byte("<1"))
	fmt.Printf("homepage has %d links!\n", numLinks)

	exists := strings.Contains(string(body), "疫情")
	fmt.Printf("是否存在疫情：%v\n", exists)

	exists = bytes.Contains(body, []byte("疫情"))
	fmt.Printf("是否存在疫情: %v\n", exists)
}
