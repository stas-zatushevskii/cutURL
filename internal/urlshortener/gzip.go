package urlshortener

import (
	"math/rand"
	"regexp"
)

const defaultCharset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-"

type URLBuilder struct {
	StringID string
	Length   int
	Charset  string
}

func NewURLBuilder(length int) *URLBuilder {
	builder := &URLBuilder{
		Length:  length,
		Charset: defaultCharset,
	}
	builder.randStringValues()
	return builder
}

func (u *URLBuilder) randStringValues() {
	randURL := make([]byte, u.Length)
	for i := range randURL {
		randURL[i] = u.Charset[rand.Intn(len(u.Charset))]
	}
	u.StringID = string(randURL)
}

func (u *URLBuilder) CreateURL(BaseURL string) string {
	return BaseURL + "/" + u.StringID
}

func URLCheck(URL string) bool {
	re := regexp.MustCompile(`^((ftp|http|https):\/\/)?([a-zA-Z0-9\-]+\.)+[a-zA-Z]{2,}(:\d+)?(\/[\w\-._~:/?#[\]@!$&'()*+,;=]*)?$`)
	return re.MatchString(URL)
}
