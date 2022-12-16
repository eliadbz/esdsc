package util

import (
	"errors"
	"flag"
	"esdsc/consts"
)

func ParseFlags() map[string]interface{}{
	flags := make(map[string]interface{})

	var wordlistPath *string = flag.String("w", "", "Subdomain names wordlist")
	var statusCode *int = flag.Int("s", 304, "Status code for indicating a non-existent subdomain")
	var isSSL *bool = flag.Bool("ssl", false, "Indicate whether it's SSL or plaintext url")
	flag.Parse()

	if (*wordlistPath) == ""{
		panic(errors.New("wordlist path cannot be empty"))
	}

	flags[consts.NON_EXISTENT_CODE] = statusCode
	flags[consts.WORDLIST_PATH] = wordlistPath
	flags[consts.IS_SSL] = isSSL
	return flags
}