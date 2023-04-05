package testcase

import (
	"fmt"
	"reflect"
	"testing"
)

func createQuery(q any) string {
	// 判断类型为结构体
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		// 获取结构体名字
		t := reflect.TypeOf(q).Name()
		// 查询语句
		query := fmt.Sprintf("insert into %s values(", t)
		v := reflect.ValueOf(q)
		// 遍历结构体字段
		for i := 0; i < v.NumField(); i++ {
			// 判断结构体类型
			switch v.Field(i).Kind() {
			case reflect.Int:
				if i == 0 {
					query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
				} else {
					query = fmt.Sprintf("%s, %d", query, v.Field(i).Int())
				}
			case reflect.String:
				if i == 0 {
					query = fmt.Sprintf("%s\"%s\"", query, v.Field(i).String())
				} else {
					query = fmt.Sprintf("%s, \"%s\"", query, v.Field(i).String())
				}
			}
		}

		query = fmt.Sprintf("%s)", query)
		fmt.Println(query)
		return query
	}

	return ""
}

type Trade struct {
	tradeId int
	Price   int
}

func TestCreateQuery(t *testing.T) {
	createQuery(Trade{tradeId: 123, Price: 456})
}
