package util

import "fmt"

func InsertSubdomain(subdomain *string, domain *string) string{
	return fmt.Sprintf("%s.%s", *subdomain, *domain)
}