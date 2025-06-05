package urlshortener

import (
	"encoding/json"
	"math/rand"
	"regexp"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-"

type URLPair struct {
	OldURL   string `json:"oldURL"`
	ShortURL string `json:"shortURL"`
}

func RandURL(length int) string {
	randURL := make([]byte, length)
	for i := range randURL {
		randURL[i] = charset[rand.Intn(len(charset))]
	}

	return "http://localhost:8080/" + string(randURL)
}

func CreateJson(oldURL, newURL string) string {
	data := URLPair{
		OldURL:   oldURL,
		ShortURL: newURL,
	}
	jsonStr, _ := json.Marshal(data)
	return string(jsonStr)
}

func ParseJSON(jsonData string) (oldURL, newURL string, error error) {
	var data URLPair
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		return "", "", err
	}
	return data.OldURL, data.ShortURL, nil
}

func URLCheck(URL string) bool {
	re := regexp.MustCompile(`^((ftp|http|https):\/\/)?([a-zA-Z0-9\-]+\.)+[a-zA-Z]{2,}(:\d+)?(\/[\w\-._~:/?#[\]@!$&'()*+,;=]*)?$`)
	return re.MatchString(URL)
}
