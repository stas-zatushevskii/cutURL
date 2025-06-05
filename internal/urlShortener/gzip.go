package urlShortener

import (
	"encoding/json"
	"math/rand"
	"regexp"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-"

type UrlPair struct {
	OldURL   string `json:"oldURL"`
	ShortURL string `json:"shortURL"`
}

func RandUrl(length int) string {
	randUrl := make([]byte, length)
	for i := range randUrl {
		randUrl[i] = charset[rand.Intn(len(charset))]
	}

	return "http://localhost:8080/" + string(randUrl)
}

func CreateJson(oldUrl, newUrl string) string {
	data := UrlPair{
		OldURL:   oldUrl,
		ShortURL: newUrl,
	}
	jsonStr, _ := json.Marshal(data)
	return string(jsonStr)
}

func ParseJson(jsonData string) (oldUrl, newUrl string, error error) {
	var data UrlPair
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		return "", "", err
	}
	return data.OldURL, data.ShortURL, nil
}

func UrlCheck(url string) bool {
	re := regexp.MustCompile(`^((ftp|http|https):\/\/)?([a-zA-Z0-9\-]+\.)+[a-zA-Z]{2,}(:\d+)?(\/[\w\-._~:/?#[\]@!$&'()*+,;=]*)?$`)
	return re.MatchString(url)
}
