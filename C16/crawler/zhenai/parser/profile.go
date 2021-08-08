package parser

import (
	"ccmouse-go/C16/crawler/engine"
	"ccmouse-go/C16/crawler/model"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([\d]+)岁</div>`)
var heightRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([\d]+)cm</div>`)
var incomeRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>月收入:([^>]+)</div>`)
var weightRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([\d]+)kg</div>`)
var genderRe = regexp.MustCompile(`"genderString":"([^"]*)"`)
var xinzuoRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([^>]+座)[^>]*</div>`)
var marriageRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([离异|丧偶|未婚]+)</div>`)
var educationRe = regexp.MustCompile(`"educationString":"([^"]*)"`)
var occupationRe = regexp.MustCompile(``)
var hokouRe = regexp.MustCompile(`<div class="m-btn pink"[^>]*>籍贯:([^>]+)</div>`)
var houseRe = regexp.MustCompile(`<div class="m-btn pink"[^>]*>([^>]+房)</div>`)
var carRe = regexp.MustCompile(`<div class="m-btn pink"[^>]*>([^>]+车)</div>`)

func ParserProfile(contents []byte, name string) engine.ParserResult {
	profile := model.Profile{}
	profile.Name = name

	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err == nil {
		profile.Age = age
	}
	height, err := strconv.Atoi(extractString(contents, heightRe))
	if err == nil {
		profile.Height = height
	}
	profile.Income = extractString(contents, incomeRe)
	weight, err := strconv.Atoi(extractString(contents, weightRe))
	if err == nil {
		profile.Weight = weight
	}
	profile.Gender = extractString(contents, genderRe)
	profile.Xinzuo = extractString(contents, xinzuoRe)
	profile.Marriage = extractString(contents, marriageRe)
	profile.Education = extractString(contents, educationRe)
	profile.Hokou = extractString(contents, hokouRe)
	profile.House = extractString(contents, houseRe)
	profile.Car = extractString(contents, carRe)

	result := engine.ParserResult{
		Items: []interface{}{profile},
		// Requests不传
	}

	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
