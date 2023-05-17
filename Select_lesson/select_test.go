package Select_lesson

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func makeHttpServer(sleep_time int) *httptest.Server {
	newServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Duration(sleep_time) * time.Millisecond)
		w.WriteHeader(http.StatusOK)

	}))

	return newServer
}

func TestWebsiteRacer(t *testing.T) {
	slowServer := makeHttpServer(2)

	fastServer := makeHttpServer(0)

	defer slowServer.Close()
	defer fastServer.Close()

	fastUrl := fastServer.URL
	slowUrl := slowServer.URL

	got := WebsiteRacer(fastUrl, slowUrl)
	want := fastUrl

	if got != want {
		t.Errorf("Should get output %s but got %s", want, got)
	}
}
