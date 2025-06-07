package urlshortener

import (
	"strings"
	"testing"
)

func TestURLBuilder(t *testing.T) {
	wantURL := "http://localhost:8080"
	wantLength := 8
	builder := NewURLBuilder(wantLength)
	newURL := builder.CreateURL()
	if strings.HasPrefix(newURL, wantURL) != true {
		t.Errorf("URLBuilder.CreateURL() = %v, want = %v", newURL, wantURL)
	}
	if len([]rune(builder.StringID)) != 8 {
		t.Errorf("Wrong lenght of StringID = %v, want = %v", len(newURL), wantLength)
	}
}
