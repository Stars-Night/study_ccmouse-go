package view

import (
	"ccmouse-go/C16/crawler6/engine"
	"ccmouse-go/C16/crawler6/frontend/model"
	common "ccmouse-go/C16/crawler6/model"
	"os"
	"testing"
)

func TestSearchResultView_Render(t *testing.T) {
	view := CreateSearchResultView("template.html")

	out, err := os.Create("template_test.html")
	page := model.SearchResult{}
	page.Hits = 356
	item := engine.Item{
		Url:  "http://album.zhenai.com/u/1106076361",
		Type: "zhenai",
		Id:   "1106076361",
		Payload: common.Profile{
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
	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, item)
	}

	err = view.Render(out, page)
	if err != nil {
		panic(err)
	}
}
