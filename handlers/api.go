package handlers

import (
	"esdsc/util"
	"net/http"
	"sync"
)

type SubdomainResponse struct{
	Exists bool
	FQDN *string
	StatusCode int
}
type SubdomainEnumerateApiClient struct{
	BaseDomain string
	Client http.Client
	isSSL bool
	StatusCodeIndicator int
}

func (client *SubdomainEnumerateApiClient) Init(baseDomain string, isSSL bool, StatusCodeIndicator int){
	/*
	Initializes the struct
	*/
	client.BaseDomain = baseDomain
	client.Client = http.Client{}
	client.isSSL = isSSL
	client.StatusCodeIndicator = StatusCodeIndicator
}

func (client *SubdomainEnumerateApiClient) executeRequest(url string) (*http.Response, error){
	/*
	Executes the requests and returns the response, along with an error if such exists
	*/
	req, err := http.NewRequest("GET", url, nil)
	if !(err == nil){
		return nil, err
	}
	resp, err := client.Client.Do(req)
	return resp, err
}

func (client *SubdomainEnumerateApiClient) statusCodesEqual(resp http.Response) bool{
	return resp.StatusCode == client.StatusCodeIndicator
}

func (client *SubdomainEnumerateApiClient) SubdomainExists(subdomain string) (*SubdomainResponse, error){
	/*
	Tests whether a subdomain exists or not by the status code found in the response
	*/
	fqdn := util.InsertSubdomain(&subdomain, &client.BaseDomain)
	url := util.ConstructURL(fqdn, client.isSSL)
	resp, err := client.executeRequest(url)
	if !(err == nil){
		return &SubdomainResponse{false, nil, 0}, err
	}
	return &SubdomainResponse{!client.statusCodesEqual(*resp), &fqdn, resp.StatusCode}, nil //If it evaluates to false, it indicates that the subdomain doesn't exist
}

func (client *SubdomainEnumerateApiClient) SubdomainExistsAsync(subdomain string, responseChannel chan *SubdomainResponse, responseWaitGroup *sync.WaitGroup){
	val, err := client.SubdomainExists(subdomain)
	if err != nil{
		responseWaitGroup.Done()
		return
	}
	responseChannel <- val
	responseWaitGroup.Done()
}