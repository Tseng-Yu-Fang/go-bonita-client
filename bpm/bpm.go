package bpm

import (
	"github.com/go-resty/resty/v2"
	"log"
)

var bc *BPMClient

type BPMClient struct {
	server   string
	token    string
	username string
	password string
	client   *resty.Client
}

func init() {
	const server_addr = "http://1.2.3.4:8080" + "/bonita/"
	// sources := fmt.Sprintf(server_addr,
	// 	// os.Getenv("BPM_SERVER_ADDR"),
	// 	os.Getenv("b.server"),
	// )
	bc = &BPMClient{
		server:   server_addr,
		token:    "",
		username: "",
		// password: "123456",
		client: resty.New(),
	}
}

// Login
// /bonita/loginservice
func (b *BPMClient) Login(username string) {

	url := b.server + "loginservice"

	resp, err := b.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"username":"` + username + `", "password":"123456"}`).
		Post(url)
	if err != nil {
		log.Fatal(err)
	}

	b.token = resp.Header().Get("X-Bonita-API-Token")
}

// Start-Form
// /bonita/API/bpm/process/[ProcessId]/instantiation
// [ProcessId] == 表單編號
// return caseId
func (b *BPMClient) StartForm(processID string, body string) []byte {

	url := b.server + "API/bpm/process/" + processID + "/instantiation"

	resp, err := b.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(url)
	if err != nil {
		log.Fatal(err)
	}

	return resp.Body()
}
