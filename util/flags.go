package util

import (
	"errors"
	"flag"
	"esdsc/consts"
	"esdsc/types"
)

func ParseFlags() map[string]*types.FlagsValue{
	flags := make(map[string]*types.FlagsValue)

	var wordlistPath *string = flag.String("w", "", "Subdomain names wordlist")
	var statusCode *int = flag.Int("s", 304, "Status code for indicating a non-existent subdomain")
	var isSSL *bool = flag.Bool("ssl", false, "Indicate whether it's SSL or plaintext url")
	flag.Parse()

	if (*wordlistPath) == ""{
		panic(errors.New("wordlist path cannot be empty"))
	}

	flags[consts.NON_EXISTENT_CODE] = &types.FlagsValue{Int: *statusCode}
	flags[consts.WORDLIST_PATH] = &types.FlagsValue{String: *wordlistPath}
	flags[consts.IS_SSL] =&types.FlagsValue{Bool: *isSSL}
	return flags
}