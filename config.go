package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const API_BASE_URL = "https://{{site_url}}.atlassian.net/rest/api/{{api_version}}"

type Config struct {
	SiteUrl    string
	ApiUser    string
	ApiKey     string
	ApiVersion string
}

type ApiClient struct {
	BaseUrl string
	ApiUser string
	ApiKey  string
}

// Client configures and returns a fully initialized AWSClient
func (c *Config) Client() (interface{}, error) {

	client := &ApiClient{
		ApiKey:  c.ApiKey,
		ApiUser: c.ApiUser,
		BaseUrl: strings.ReplaceAll(strings.ReplaceAll(API_BASE_URL, "{{site_url}}", c.SiteUrl), "{{api_version}}", c.ApiVersion),
	}

	return client, nil
}

func (a *ApiClient) request(method, resource, body string) (map[string]interface{}, error) {
	reqBody := []byte(body)
	httpClient := &http.Client{}
	log.Printf("REQUEST URL: %s", a.BaseUrl+"/"+resource)
	log.Printf("REQUEST BODY: %s", body)
	req, err := http.NewRequest(method, a.BaseUrl+"/"+resource, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	} else {
		req.SetBasicAuth(a.ApiUser, a.ApiKey)
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Accept", "application/json")
		resp, err := httpClient.Do(req)

		if err != nil {
			return nil, err
		} else {
			defer resp.Body.Close()
			if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusCreated || resp.StatusCode == http.StatusAccepted || resp.StatusCode == http.StatusNoContent {
				bodyBytes, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					return nil, err
				} else {
					var result map[string]interface{}
					json.Unmarshal(bodyBytes, &result)
					return result, nil
				}
			} else {
				return nil, errors.New(resp.Status)
			}
		}
	}
}
