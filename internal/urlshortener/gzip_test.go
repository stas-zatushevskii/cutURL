package urlshortener

import "testing"

func TestCreateJson(t *testing.T) {
	oldURL := "https://www.google.com"
	newURL := "https" + RandUrl(8)
	want := "{\"" + oldURL + "\":\"" + newURL + "\"}"
	result := CreateJson(oldURL, newURL)

	if result != want {
		t.Errorf("CreateJson(oldURL, newURL) = %s, want %s", result, want)
	}
}
