package helpers

import (
	"os"
	"strings"
)

func EnforceHTTP(url string) string {
	if url[:4] != "http" {
		return "http://" + url
	}
	return url
}

func RemoveDomainError(url string) bool {
	if url == "" {
		return false
	}
	if url == os.Getenv("DOMAIN") {
		return false
	}
	newURL := strings.Replace(url, "http://", "", 1)
	newURL = strings.Replace(newURL, "https://", "", 1)
	newURL = strings.Replace(newURL, "www.", "", 1)
	newURL = strings.TrimSuffix(newURL, "/")
	newURL = strings.Split(newURL, "/")[0]

	return newURL != os.Getenv("DOMAIN")

	// const expression = `^(?:f|ht)tps?:\/\/(\w+:{0,1}\w*@)?(\S+)(:[0-9]+)?(\/|\/([\w#!:.?+=&%@!\-\/]))?$`
	// var r = regexp.MustCompile(expression)
	// return r.MatchString(url)
}
