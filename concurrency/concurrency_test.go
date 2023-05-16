package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(s string) bool {
	if s == "https:/hckighekng.com/" {
		return false
	} else {
		return true
	}
}

func slowWebsiteChecker(_ string) bool {
	time.Sleep(10 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urlss := make([]string, 100)

	for i := 0; i < len(urlss); i++ {
		urlss[i] = "an url"
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowWebsiteChecker, urlss)
	}
}

func TestCheckWebsites(t *testing.T) {
	urlss := []string{"https:/google.com",
		"https:/bigo.com",
		"https:/hckighekng.com/",
		"https:/tumor.com"}

	got := CheckWebsites(mockWebsiteChecker, urlss)
	want := map[string]bool{"https:/google.com": true,
		"https:/bigo.com":        true,
		"https:/hckighekng.com/": false,
		"https:/tumor.com":       true}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v but wanted %v", got, want)
	}
}
