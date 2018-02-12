package util

import (
	"encoding/xml"
	"fmt"
	"testing"
)

type Address struct {
	City, State string
}
type Person struct {
	XMLName   xml.Name `xml:"person"`
	Id        int      `xml:"id,attr"`
	FirstName string   `xml:"name>first"`
	LastName  string   `xml:"name>last"`
	Age       int      `xml:"age"`
	Height    float32  `xml:"height,omitempty"`
	Married   bool
	Address
	Comment string `xml:"comment"`
}

// func TestXML(t *testing.T) {
// 	v := &Person{Id: 13, FirstName: "John", LastName: "Doe", Age: 42, Height: 12.4}
// 	v.Comment = " Need more details. "
// 	v.Address = Address{"Hanga Roa", "Easter Island"}

// 	// PostXML("test", v)
// }

func TestGenTestURL(t *testing.T) {
	url := "https://api.mch.weixin.qq.com/pay/micropay"
	testURL := GenTestURL(url)
	if testURL == "https://api.mch.weixin.qq.com/sandboxnew/pay/micropay" {
		t.Log("sucess ....")
	} else {
		t.Error("error")
	}
}

func TestReflect(t *testing.T) {
	v := Person{Id: 13, FirstName: "John", LastName: "Doe", Height: 12.4}
	v.Comment = " Need more details. "
	v.Address = Address{"Hanga Roa", "Easter Island"}
	GenWechatSign("", v)
}

type a struct {
	Id int
}

type b struct {
	Name string
}

type c struct {
	a
	b
}

func TestHeir(t *testing.T) {
	t1 := a{
		Id: 2,
	}
	t2 := b{
		Name: "string",
	}
	test := c{
		b: t2,
		a: t1,
	}

	fmt.Println(test)
}
