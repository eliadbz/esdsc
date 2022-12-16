package handlers

import (
	"errors"
	"esdsc/util"
	"fmt"
	"net/http"
)

type ApiClient struct{
	BaseDomain string
	Client http.Client
	isSSL bool
	StatusCodeIndicator
}

func (client *ApiClient) Init(baseDomain string, isSSL bool, StatusCodeIndicator int){
	/*
	Initializes the struct
	*/
	client.BaseDomain = baseDomain
	client.Client = http.Client{}
	client.isSSL = isSSL
	client.StatusCodeIndicator = StatusCodeIndicator
}

func (client *ApiClient) executeRequest(url string) (*http.Response, error){
	/*
	Executes the requests and returns the response, along with an error if such exists
	*/
	req, err := http.NewRequest("GET", url, nil)
	if !(err == nil){
		panic(err)
	}
	resp, err := client.Client.Do(req)
	return resp, err
}

func (client *ApiClient) statusCodesEqual(resp http.Response) bool{
	return resp.StatusCode == client.StatusCodeIndicator
}

func (client *ApiClient) SubdomainExists(subdomain string) bool{
	/*
	Tests whether a subdomain exists or not by the status code found in the response
	*/
	fqdn := util.InsertSubdomain(&subdomain, &client.BaseDomain)
	url := util.ConstructURL(fqdn, client.isSSL)
	resp, err := client.executeRequest(url)
	if !(err == nil){
		panic(err)
	}
	return client.statusCodesEqual(*resp) //If it evaluates to true, it indicates that the subdomain doesn't exist
}