package testcase

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestReadSlice(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("https://studygolang.com. \nIt is the home of gophers"))
	line, _ := reader.ReadSlice('\n')
	fmt.Printf("the line: %s\n", line)

	// 这里可以换上任意的 bufio 的 Read/Write 操作
	n, _ := reader.ReadSlice('\n')
	fmt.Printf("the line: %s\n", line)
	fmt.Println(string(n))
}

func TestReadString(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("https://studygolang.com. \nIt is the home of gophers"))
	line, _ := reader.ReadString('\n')
	fmt.Printf("the line: %s\n", line)

	// 这里可以换上任意的 bufio 的 Read/Write 操作
	n, _ := reader.ReadString('\n')
	fmt.Printf("the line: %s\n", line)
	fmt.Println(string(n))
}

func TestReadSliceV2(t *testing.T) {
	reader := bufio.NewReaderSize(strings.NewReader("http://studygolang.com"), 16)
	line, err := reader.ReadSlice('\n')
	fmt.Printf("line:%s\terror:%s\n", line, err)

	line, err = reader.ReadSlice('\n')
	fmt.Printf("line:%s\terror:%s\n", line, err)
}

func TestPeek(t *testing.T) {
	reader := bufio.NewReaderSize(strings.NewReader("http://studygolang.com.\t It is the home of gophers"), 14)
	go Peek(reader)
	go reader.ReadBytes('\t')
	time.Sleep(1e8)
}

func Peek(reader *bufio.Reader) {
	line, _ := reader.Peek(14)
	fmt.Printf("%s\n", line)
	time.Sleep(1)
	fmt.Printf("%s\n", line)
}
