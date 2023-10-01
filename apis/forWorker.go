package apis

import (
	"fmt"

	"goroutine.com/client"
	"goroutine.com/constant"
)

type ForWorker struct {
	*client.Client
}

func NewForWorker(c *client.Client) *ForWorker {
	return &ForWorker{c}
}

func (f *ForWorker) ExecuteAPI() *client.Response {
	return f.Execute(fmt.Sprintf("%s%s", constant.BASE_URL, constant.MORIOKA))
}
