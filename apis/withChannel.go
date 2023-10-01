package apis

import (
	"fmt"

	"goroutine.com/client"
	"goroutine.com/constant"
)

type WithChannel struct {
	*client.Client
}

func NewWithChannel(c *client.Client) *WithChannel {
	return &WithChannel{
		c,
	}
}

// -------------------------------------------------
// Channel Sender
// Group A
func (w *WithChannel) ExecuteAPI1(g chan string) {
	url := fmt.Sprintf("%s%s", constant.BASE_URL, constant.MORIOKA)
	g <- w.Execute(url).Title
}

// Group A
func (w *WithChannel) ExecuteAPI2(g chan string) {
	url := fmt.Sprintf("%s%s", constant.BASE_URL, constant.SENDAI)
	g <- w.Execute(url).Title
}

// Group B
func (w *WithChannel) ExecuteAPI3(g chan string) {
	url := fmt.Sprintf("%s%s", constant.BASE_URL, constant.KANAZAWA)
	g <- w.Execute(url).Title
}

// Group B
func (w *WithChannel) ExecuteAPI4(g chan string) {
	url := fmt.Sprintf("%s%s", constant.BASE_URL, constant.NAGANO)
	g <- w.Execute(url).Title
}

// Group C
func (w *WithChannel) ExecuteAPI5(g chan string) {
	url := fmt.Sprintf("%s%s", constant.BASE_URL, constant.GIHU)
	g <- w.Execute(url).Title
}

// Group C
func (w *WithChannel) ExecuteAPI6(g chan string) {
	url := fmt.Sprintf("%s%s", constant.BASE_URL, constant.SHIZUOKA)
	g <- w.Execute(url).Title
}

// Group D
func (w *WithChannel) ExecuteAPI7(gd chan string) {
	url := fmt.Sprintf("%s%s", constant.BASE_URL, constant.KYOTO)
	gd <- w.Execute(url).Title
}

// -------------------------------------------------
// Channel Getter
func (w *WithChannel) GetAPIGroup(gType string, g chan string) {
	fmt.Println(fmt.Sprintf("Group %s : ", gType), <-g)
}

func (w *WithChannel) GetAPIGroup2(gType string, g chan string) {
	result, ok := <-g
	fmt.Printf("Before G%s is not closed? : %v\n", gType, ok)
	fmt.Println(fmt.Sprintf("Group %s : ", gType), result)
}

// ================================================
// Channelの引数に送受信を指定

func (w *WithChannel) ExecuteAPI1WithSend(g chan<- string) {
	url := fmt.Sprintf("%s%s", constant.BASE_URL, constant.MORIOKA)
	g <- w.Execute(url).Title
}

// Group A
func (w *WithChannel) ExecuteAPI2WithSend(g chan<- string) {
	url := fmt.Sprintf("%s%s", constant.BASE_URL, constant.SENDAI)
	g <- w.Execute(url).Title
}

// Group B
func (w *WithChannel) ExecuteAPI3WithSend(g chan<- string) {
	url := fmt.Sprintf("%s%s", constant.BASE_URL, constant.KANAZAWA)
	g <- w.Execute(url).Title
}

// Group B
func (w *WithChannel) ExecuteAPI4WithSend(g chan<- string) {
	url := fmt.Sprintf("%s%s", constant.BASE_URL, constant.NAGANO)
	g <- w.Execute(url).Title
}

// Group C
func (w *WithChannel) ExecuteAPI5WithSend(g chan<- string) {
	url := fmt.Sprintf("%s%s", constant.BASE_URL, constant.GIHU)
	g <- w.Execute(url).Title
}

// Group C
func (w *WithChannel) ExecuteAPI6WithSend(g chan<- string) {
	url := fmt.Sprintf("%s%s", constant.BASE_URL, constant.SHIZUOKA)
	g <- w.Execute(url).Title
}

// Group D
func (w *WithChannel) ExecuteAPI7WithSend(gd chan<- string) {
	url := fmt.Sprintf("%s%s", constant.BASE_URL, constant.KYOTO)
	gd <- w.Execute(url).Title
}

func (w *WithChannel) GetAPIGroupWithRecive(gType string, g <-chan string) {
	fmt.Println(fmt.Sprintf("Group %s : ", gType), <-g)
}
