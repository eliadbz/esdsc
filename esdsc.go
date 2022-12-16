package main

import (
	"esdsc/consts"
	"esdsc/util"
	"fmt"
)

func main(){
	flags := util.ParseFlags()
	data := util.ReadWordlist(&flags[consts.WORDLIST_PATH].String)
	domain := "test.com"
	subdomain := "test"
	println(util.InsertSubdomain(&subdomain, &domain))
	fmt.Println(data)
}