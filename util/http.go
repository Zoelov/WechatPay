package util

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

// PostXML post xml 数据请求
func PostXML(url string, postBody interface{}) ([]byte, error) {
	xmlData, err := xml.MarshalIndent(postBody, " ", "	")
	if err != nil {
		return nil, fmt.Errorf("post body is not xml format")
	}
	body := bytes.NewBuffer(xmlData)

	response, err := http.Post(url, "application/xml;charset=utf-8", body)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http post error: url=%v statusCode=%v", url, response.StatusCode)
	}

	return ioutil.ReadAll(response.Body)
}
