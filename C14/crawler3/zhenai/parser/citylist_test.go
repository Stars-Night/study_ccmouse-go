package parser

import (
	"io/ioutil"
	"testing"
)

//14-6课程
func TestParserCityList(t *testing.T) {
	content, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}

	list := ParserCityList(content)
	const ResultSize = 470
	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba", "http://www.zhenai.com/zhenghun/akesu", "http://www.zhenai.com/zhenghun/alashanmeng",
	}
	expectedCities := []string{
		"阿坝", "阿克苏", "阿拉善盟",
	}

	if len(list.Requests) != ResultSize {
		t.Errorf("result should have %d"+"requests; but had %d", ResultSize, len(list.Requests))
	}
	for i, url := range expectedUrls {
		if list.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s; but"+"was %s", i, url, list.Requests[i].Url)
		}
	}

	if len(list.Items) != ResultSize {
		t.Errorf("result should have %d"+"Items; but had %d", ResultSize, len(list.Items))
	}
	for i, city := range expectedCities {
		if list.Items[i].(string) != city {
			t.Errorf("expected city #%d: %s; but"+"was %s", i, city, list.Items[i].(string))
		}
	}
}
