package main

import (
	"esdsc/consts"
	"esdsc/handlers"
	"esdsc/util"
	"sync"
	"fmt"
)

func initObjects() (*handlers.SubdomainEnumerateApiClient, chan string, error){
	flags := util.ParseFlags() //read flags
	dataChannel, err := util.ReadWordlist(flags[consts.WORDLIST_PATH].(*string)) //prepare to read file lines
	if err != nil{
		return nil, nil, err
	}
	domain := flags[consts.BASE_DOMAIN].(*string)
	isSSL := flags[consts.IS_SSL].(*bool)
	statusCode := flags[consts.NON_EXISTENT_CODE].(*int)

	client := handlers.SubdomainEnumerateApiClient{} // init the api client struct
	client.Init(*domain, *isSSL, *statusCode)
	return &client, dataChannel, nil
}

func main(){
	client, dataChannel, err := initObjects()
	if err != nil{
		panic(err)
	}
	responseChannel := make(chan *handlers.SubdomainResponse)
	var responseWaitGroup sync.WaitGroup
	for name := range dataChannel{ //launch all goroutines
		responseWaitGroup.Add(1)
		go client.SubdomainExistsAsync(name, responseChannel, &responseWaitGroup)
	}
	go func(){ //closes the channel after all goroutines are done
		responseWaitGroup.Wait()
		close(responseChannel)
	}()
	for resp := range responseChannel{ //processes the information
		if resp.Exists{
			fmt.Printf("%s -> %d\n",*(resp.FQDN), resp.StatusCode)
		}
	}
}