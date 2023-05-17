package Select_lesson

import (
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
