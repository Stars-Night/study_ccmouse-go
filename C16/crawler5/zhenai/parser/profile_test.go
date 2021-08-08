package parser

import (
	"ccmouse-go/C16/crawler5/engine"
	"ccmouse-go/C16/crawler5/model"
	"io/ioutil"
	"testing"
)

func TestParserProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParserProfile(contents, "http://album.zhenai.com/u/1106076361", "萍儿")
	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 "+"element; but was %v", result.Items)
	}
	profile := result.Items[0]

	expected := engine.Item{
		Url:  "http://album.zhenai.com/u/1106076361",
		Type: "zhenai",
		Id:   "1106076361",
		Payload: model.Profile{
			Age:       23,
			Height:    155,
			Weight:    45,
			Gender:    "女士",
			Income:    "5-8千",
			Marriage:  "未婚",
			Education: "高中及以下",
			Xinzuo:    "魔羯座",
			Car:       "未买车",
			Name:      "萍儿",
			House:     "租房",
			Hokou:     "湖南娄底",
		},
	}

	if expected != profile {
		t.Errorf("expected %v; but was %v", expected, profile)
	}
}
