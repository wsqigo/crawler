package testcase

import (
	"context"
	"github.com/chromedp/chromedp"
	"log"
	"testing"
	"time"
)

func TestClickElement(t *testing.T) {
	// 创建实例
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		// chromedp.WithDebugf(log.Printf), //是否显示调试信息
	)
	defer cancel()

	// 创建超时
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// 导航到页面，等待元素，点击
	var example string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://pkg.go.dev/time`),
		// 等待页脚元素课件（即页面已加载）
		chromedp.WaitVisible(`body > footer`),
		// 查找并单击“示例”链接
		chromedp.Click(`#example-After`, chromedp.NodeVisible),
		// 检索textarea的文本
		chromedp.Value(`#example-After textarea`, &example),
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("获取到示例数据：\n%s", example)
}
