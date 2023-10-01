package client

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Client struct {
	*http.Client
}

// Responseの構造体
type Response struct {
	Title string `json:"title"`
}

func NewClient() *Client {
	return &Client{
		new(http.Client),
	}
}

func (c *Client) Execute(url string) *Response {
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := c.Do(req)
	defer res.Body.Close()

	return c.convertRes(res)
}

func (c *Client) convertRes(res *http.Response) *Response {
	body, _ := io.ReadAll(res.Body)

	var resbody Response
	if err := json.Unmarshal(body, &resbody); err != nil {
		log.Fatal(err)
		return nil
	}

	return &resbody
}
