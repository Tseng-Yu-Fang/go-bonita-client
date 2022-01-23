package main

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

func main() {
	// just mentioning about POST as an example with simple flow
	// User Login
	client := resty.New()
	APIToken := ""
	resp, _ := client.R().
		SetFormData(map[string]string{
			"username": "isabelle_wu",
			"password": "12345",
		}).
		Post("http://54.169.182.165:8080/bonita/loginservice")

	for _, cookie := range resp.Cookies() {
		if cookie.Name == "X-Bonita-API-Token" {
			APIToken = cookie.Value
		}
	}
	resp, _ = client.R().SetHeaders(map[string]string{
		"Content-Type":       "application/json",
		"X-Bonita-API-Token": APIToken,
	}).
		SetBody(`{
			"modelInput":
			{
					"assistant":"choc",
					"recipient":"kevin_lin" 
			}
	}`).
		Post("http://54.169.182.165:8080/bonita/API/bpm/process/8759976868088592450/instantiation")

	fmt.Println(resp)
}
