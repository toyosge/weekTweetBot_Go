package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Auth struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

var (
	clientId     = "xxxxxxxxxxxxxxxxxx"
	clientSecret = "xxxxxxxxxxxxxxxxxx"
	topicId      = "xxxxxxxxxxxxxxxxxx"
	message      = "hello"
)

func main() {
	message = createMessage()
	// postMessage()
	fmt.Println(message)

}

func createMessage() string {
	m := make([]byte, 0, 100)
	m = append(b, "massanのへっぽこbotからお送りします\n"...)
	m = append(b, sayDayInfo()...)
	m = append(b, "\n"...)
	m = append(b, sayWeek()...)
	m = append(b, "...(σ･∀･)σゲッツ!!"...)
	return string(b)
}

func sayDayInfo() string {
	t := time.Now()
	month := t.Month().String()
	day := strconv.Itoa(t.Day())
	return "Today is " + month + " " + day + "."
}

func sayWeek() string {
	week := setWeekday()
	switch week {
	case "Monday":
		return "月曜日です！！今週もがんばりましょう。OK牧場"
	case "Tuesday":
		return "火曜日でございまする！！"
	case "Wednesday":
		return "すいよーーーーーび"
	case "Thursday":
		return "も、もくようび..."
	case "Friday":
		return "今日は金曜日です..."
	case "Saturday", "Sunday":
		return "おやすみだよね〜"
	}
	return ""
}
func setWeekday() string {
	t := time.Now()
	return t.Weekday().String()
}

func postMessage() {
	resp, err := http.PostForm(
		"https://typetalk.in/oauth2/access_token",
		url.Values{
			"client_id":     {clientId},
			"client_secret": {clientSecret},
			"grant_type":    {"client_credentials"},
			"scope":         {"topic.post"}})
	if err != nil {
		panic(err)
	}
	var d Auth
	err = json.NewDecoder(resp.Body).Decode(&d)
	if err != nil {
		panic(err)
	}
	resp, err = http.PostForm(
		fmt.Sprintf("https://typetalk.in/api/v1/topics/%s", topicId),
		url.Values{
			"access_token": {d.AccessToken},
			"message":      {message}})
	if err != nil {
		panic(err)
	}
}
