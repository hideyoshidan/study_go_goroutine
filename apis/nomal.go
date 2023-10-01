package apis

import (
	"fmt"

	"goroutine.com/client"
	"goroutine.com/constant"
)

type Nomal struct {
	*client.Client
}

func NewNomal(c *client.Client) *Nomal {
	return &Nomal{c}
}

func (n *Nomal) ExecuteAPI1() {
	url := fmt.Sprintf("%s%s", constant.BASE_URL, constant.MORIOKA)

	fmt.Println("1番目", n.Execute(url).Title)
}

func (n *Nomal) ExecuteAPI2() {
	url := fmt.Sprintf("%s%s", constant.BASE_URL, constant.SENDAI)

	fmt.Println("2番目", n.Execute(url).Title)
}

func (n *Nomal) ExecuteAPI3() {
	url := fmt.Sprintf("%s%s", constant.BASE_URL, constant.KANAZAWA)

	fmt.Println("3番目", n.Execute(url).Title)
}

func (n *Nomal) ExecuteAPI4() {
	url := fmt.Sprintf("%s%s", constant.BASE_URL, constant.NAGANO)

	fmt.Println("4番目", n.Execute(url).Title)
}

func (n *Nomal) ExecuteAPI5() {
	url := fmt.Sprintf("%s%s", constant.BASE_URL, constant.GIHU)

	fmt.Println("5番目", n.Execute(url).Title)
}

func (n *Nomal) ExecuteAPI6() {
	url := fmt.Sprintf("%s%s", constant.BASE_URL, constant.SHIZUOKA)

	fmt.Println("6番目", n.Execute(url).Title)
}

func (n *Nomal) ExecuteAPI7() {
	url := fmt.Sprintf("%s%s", constant.BASE_URL, constant.KYOTO)

	fmt.Println("7番目", n.Execute(url).Title)
}

// 指定したcityCodeの天気APIを取得する
func (n *Nomal) ExecuteAPI(cityCode string) string {
	url := fmt.Sprintf("%s%s", constant.BASE_URL, cityCode)

	return n.Execute(url).Title
}
