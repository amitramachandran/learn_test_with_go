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

func TestWebsiteSelectRacer(t *testing.T) {
	slowServer := makeHttpServer(2)

	fastServer := makeHttpServer(0)

	defer slowServer.Close()
	defer fastServer.Close()

	fastUrl := fastServer.URL
	slowUrl := slowServer.URL

	_, err := WebsiteSelectRacer(fastUrl, slowUrl, time.Duration(10))
	// want := fastUrl

	if err == nil {
		t.Errorf("expected an error %s", err)
	}
}        
