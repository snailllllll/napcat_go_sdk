package napcat_go_sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type HttpClient struct {
	url    string
	token  *string
	client *http.Client
}

func NewHttpClient(url string, token *string) *HttpClient {
	return &HttpClient{
		url:    url,
		token:  token,
		client: &http.Client{},
	}
}

type HttpResponse[T any] struct {
	Status  string      `json:"status"`
	Retcode int         `json:"retcode"`
	Data    T           `json:"data"`
	Message string      `json:"message"`
	Wording string      `json:"wording"`
	Echo    interface{} `json:"echo"`
}

func (h *HttpResponse[T]) ReceiveResponseMessage(bytes []byte) error {

	return json.Unmarshal(bytes, h)
}

func (msg *Message[any]) SendHttpMsg() interface{} {
	return msg.Params
}

func (c *HttpClient) SendMessage(message SendMsg, receiveMessage ReceiveResponseMessage) error {
	marshal, err := json.Marshal(message.SendHttpMsg())
	if err != nil {
		log.Fatal(err)
		return err
	}
	req, err := http.NewRequest("POST", c.url, bytes.NewBuffer(marshal))

	req.Header.Set("Content-Type", "application/json")
	if c.token != nil {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *c.token))
	}

	do, err := c.client.Do(req)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(do.Body)

	body, err := ioutil.ReadAll(do.Body)
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = receiveMessage.ReceiveResponseMessage(body)
	if err != nil {
		return err
	}
	return nil
}
