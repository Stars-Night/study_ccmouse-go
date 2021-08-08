package model

import "encoding/json"

type Profile struct {
	Name       string
	Gender     string //性别
	Age        int
	Height     int
	Weight     int
	Income     string //收入
	Marriage   string //婚姻
	Education  string //教育
	Occupation string //职业
	Hokou      string //户口
	Xinzuo     string
	House      string
	Car        string
}

func FromJsonObj(o interface{}) (Profile, error) {
	var profile Profile
	s, err := json.Marshal(o)
	if err != nil {
		return profile, err
	}

	err = json.Unmarshal(s, &profile)
	return profile, err
}
