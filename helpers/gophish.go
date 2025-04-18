package helpers

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type options struct {
	base_url string
	api_key  string
}

type Sending_Profile struct {
	Profile_name string `json:"profile_name"`
	From_address string `json:"from_address"`
	Host         string `json:"host"` // this should be <hostname/ip>:<port>
	Username     string `json:"username"`
	Password     string `json:"password"`
}

func gophish_api_request(http_method string, request_body []byte, api_url *url.URL) []byte {

	//proxyUrl, err := url.Parse("http://192.168.2.168:8083")

	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
		//Proxy:           http.ProxyURL(proxyUrl),
	}

	httpClient := &http.Client{Transport: transCfg}

	reqUrl := fmt.Sprintf("%s/api/smtp", api_url)

	fmt.Println(reqUrl)

	req, err := http.NewRequest(http_method, reqUrl, nil)

	switch http_method {
	case http.MethodPost, http.MethodPut, http.MethodPatch:
		req, err = http.NewRequest(http_method, reqUrl, bytes.NewReader(request_body))
	default:
	}

	if err != nil {
		panic(err)
	}

	req.Header.Set("Authorization", fmt.Sprint(os.Getenv("GOPHISH_API_KEY")))
	req.Header.Set("Content-Type", "application/json")

	res, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	return body

}

func Create_gophish_sending_profile(profile Sending_Profile) {

	api_url, err := url.Parse("https://localhost:3333")
	jsonBody, err := json.Marshal(profile)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(gophish_api_request(http.MethodPost, jsonBody, api_url)))

}

func Get_gophish_sending_profiles() {
	api_url, err := url.Parse("https://localhost:3333")
	jsonBody, err := json.Marshal("")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(gophish_api_request(http.MethodGet, jsonBody, api_url)))
}
