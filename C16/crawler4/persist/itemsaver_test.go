package persist

import (
	"ccmouse-go/C16/crawler4/engine"
	"ccmouse-go/C16/crawler4/model"
	"context"
	"encoding/json"
	"gopkg.in/olivere/elastic.v5"
	"testing"
)

func TestItemSaver(t *testing.T) {
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

	// TODO: Try to start up elastic search
	// here using docker go client.
	client, err := elastic.NewClient(
		elastic.SetURL("http://192.168.99.100:9200"),
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err)
	}

	const index = "dating_test"
	err = save(client, index, expected)
	if err != nil {
		panic(err)
	}

	resp, err := client.Get().Index(index).Type(expected.Type).Id(expected.Id).Do(context.Background())
	if err != nil {
		panic(err)
	}

	t.Logf("%s", resp.Source)

	var actual engine.Item
	json.Unmarshal(*resp.Source, &actual)

	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile

	if expected != actual {
		t.Errorf("expectd %v;\n got %v", expected, actual)
	}
}
