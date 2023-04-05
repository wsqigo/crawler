package testcase

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"testing"
)

func TestElement(t *testing.T) {
	html := `<body>

				<div>DIV1</div>
				<div>DIV2</div>
				<span>SPAN</span>

			</body>
			`
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		fmt.Println(err)
	}

	dom.Find("div").Each(func(i int, s *goquery.Selection) {
		fmt.Println(s.Text())
	})
}

func TestIDSelection(t *testing.T) {
	html := `<body>

				<div id="div1">DIV1</div>
				<div>DIV2</div>
				<span>SPAN</span>

			</body>
			`

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		fmt.Println(err)
	}

	dom.Find("#div1").Each(func(i int, s *goquery.Selection) {
		fmt.Println(s.Text())
	})
}

func TestClassSelection(t *testing.T) {
	html := `<body>

				<div id="div1">DIV1</div>
				<div class="name">DIV2</div>
				<span>SPAN</span>

			</body>
			`
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		fmt.Println(err)
	}

	dom.Find(".name").Each(func(i int, s *goquery.Selection) {
		fmt.Println(s.Text())
	})
}

func TestElementAttr(t *testing.T) {
	html := `<body>

				<div>DIV1</div>
				<div class="name">DIV2</div>
				<span>SPAN</span>

			</body>
			`
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		fmt.Println(err)
	}

	dom.Find("div[class=name]").Each(func(i int, s *goquery.Selection) {
		fmt.Println(s.Text())
	})
}

func TestParentChild(t *testing.T) {
	html := `<body>

				<div lang="ZH">DIV1</div>
				<div lang="zh-cn">DIV2</div>
				<div lang="en">DIV3</div>
				<span>
					<div>DIV4</div>
				</span>

			</body>
			`
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		fmt.Println(err)
	}

	dom.Find("body>div").Each(func(i int, s *goquery.Selection) {
		fmt.Println(s.Text())
	})
}
