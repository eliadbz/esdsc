package util

import (
	"errors"
	"flag"
	"esdsc/consts"
)

func ParseFlags() map[string]interface{}{
	/*
	Parses CLI parameters
	*/
	flags := make(map[string]interface{}) //holds the flags' values' pointers

	var wordlistPath *string = flag.String("w", "", "Subdomain names wordlist") //mandatory
	var baseDomain *string = flag.String("d", "", "The base domain") //mandatory
	var statusCode *int = flag.Int("s", 301, "Status code for indicating a non-existent subdomain")
	var isSSL *bool = flag.Bool("ssl", false, "Indicate whether it's SSL or plaintext url")
	flag.Parse()

	if (*wordlistPath) == ""{
		panic(errors.New("wordlist path cannot be empty"))
	}
	if (*baseDomain) == ""{
		panic(errors.New("domain cannot be empty"))
	}

	flags[consts.NON_EXISTENT_CODE] = statusCode
	flags[consts.WORDLIST_PATH] = wordlistPath
	flags[consts.IS_SSL] = isSSL
	flags[consts.BASE_DOMAIN] = baseDomain
	return flags
}