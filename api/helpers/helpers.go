package helpers

import (
	"os"
	"strings"
)

func EnforceHTTP() string {
	if url[:4] != "http" { //String.Slice, check for HTTP Header
		return "http://" + url
	}
	return url
}

func RemoveDomainError(url string) bool {
	if  url == os.Getenv("DOMAIN"){
		return false
	}
	//All Possible URL Test Cases for Loop Manipulation
	newURL := strings.Replace(url, "http://", "", 1) //Prefixes
	newURL = strings.Replace(newURL, "http://", "", 1)
	newURL = strings.Replace(newURL, "www.", "", 1)
	newURL = strings.Split(newURL, "/")[0] //Path Separator

	if newURL == os.Getenv("DOMAIN"){
		return false
	}

	return true

}