package util

import "fmt"

func ConstructURL(fqdn string, isSSL bool) string{
	var url string
	if isSSL{
		url = fmt.Sprintf("https://%s", fqdn)
	} else {
		url = fmt.Sprintf("http://%s", fqdn)
	}
	return url
}