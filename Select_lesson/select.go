package Select_lesson

import (
	"fmt"
	"net/http"
	"time"
)

func getUrlResponseTime(url string) time.Duration {
	startTime := time.Now()
	http.Get(url)
	responseTime := time.Since(startTime)
	return responseTime
}

func WebsiteRacer(url_1, url_2 string) string {
	resp_1 := getUrlResponseTime(url_1)
	resp_2 := getUrlResponseTime(url_2)

	if resp_1 < resp_2 {
		return url_1
	} else {
		return url_2
	}

}

var tenSecondtimeout time.Duration = 10

func ConfigurableWebsiteRacer(url_1, url_2 string) (string, error) {
	return WebsiteSelectRacer(url_1, url_2, tenSecondtimeout)
}

// using select to do concurrent website pings instead of getting response time one by one
func WebsiteSelectRacer(url_1, url_2 string, timeout time.Duration) (s string, e error) {

	select {
	case <-ping(url_1):
		return url_1, e

	case <-ping(url_2):
		return url_2, e

	case <-time.After(timeout * time.Second):
		return "", fmt.Errorf("timeout %v sec exceeded", timeout)

	}
}

func ping(url string) (c chan struct{}) {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return

}
