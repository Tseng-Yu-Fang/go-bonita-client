package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

func main() {
	// just mentioning about POST as an example with simple flow
	// User Login
	client := resty.New()

	resp, _ := client.R().
		SetFormData(map[string]string{
			"username": "isabelle_wu",
			"password": "12345",
		}).
		Post("http://localhost:8080/bonita/loginservice")

	fmt.Println(resp.Header())

}
