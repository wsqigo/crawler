package testcase

import (
	"fmt"
	"regexp"
	"testing"
)

func TestMatchString(t *testing.T) {
	str := "Golang regular expressions example"
	match, err := regexp.MatchString("^Golang", str)
	fmt.Println("Match:", match, " Error:", err)
}

func TestFindString(t *testing.T) {
	str := "Golang expressions example"
	re := regexp.MustCompile("Gola([a-z]+)g")
	fmt.Println(re.FindString(str))
}

func TestFindStringIndex(t *testing.T) {
	str := "Golang regular expressions example"
	re := regexp.MustCompile(`exp`)
	match := re.FindStringIndex(str)
	fmt.Println("Match:", match)
}

func TestFindStringSubMatch(t *testing.T) {
	str := "Golang regular expressions example"
	re := regexp.MustCompile("p([a-z]+)e")

	match := re.FindStringSubmatch(str)
	fmt.Println(match)
}

func TestFindAllString(t *testing.T) {
	str := "Golang regular expressions example"
	re := regexp.MustCompile("p([a-z]+)e")

	matches := re.FindAllString(str, 2)
	fmt.Println(matches)
}

func TestFindAll(t *testing.T) {
	str := "aaabb"
	re := regexp.MustCompile("a*")
	fmt.Println(len(re.FindAllString(str, -1)))
}
