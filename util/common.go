package util

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

func GenTestURL(url string) string {
	oldStr := "https://api.mch.weixin.qq.com/"
	newStr := "https://api.mch.weixin.qq.com/sandboxnew/"

	return strings.Replace(url, oldStr, newStr, 1)
}

type KeyValue struct {
	Key   string
	Value string
}
type KeyValueSlice []KeyValue

func (k KeyValueSlice) Len() int {
	return len(k)
}

func (k KeyValueSlice) Less(i, j int) bool {
	v := strings.Compare(k[i].Key, k[j].Key)
	if v < 0 {
		return true
	}
	return false
}

func (k KeyValueSlice) Swap(i, j int) {
	k[i], k[j] = k[j], k[i]
}

func GenWechatSign(apiKey string, param interface{}) string {
	v := reflect.ValueOf(param)
	t := reflect.TypeOf(param)
	count := v.NumField()

	dataList := make([]KeyValue, 0)
	for i := 0; i < count; i++ {
		f := v.Field(i)
		data := new(KeyValue)
		data.Key = t.Field(i).Name
		if data.Key == "sign" {
			continue
		}
		var shoudAdd bool
		fmt.Print(t.Field(i).Name)
		switch f.Kind() {
		case reflect.String:
			if f.String() != "" {
				data.Value = f.String()
				shoudAdd = true
			}
		case reflect.Int:
			if f.Int() != 0 {
				data.Value = strconv.Itoa(int(f.Int()))
				shoudAdd = true
			}
		}
		if shoudAdd {
			dataList = append(dataList, *data)
		}
	}

	sort.Sort(KeyValueSlice(dataList))

	for i := 0; i < len(dataList); i++ {

	}

	return ""
}
