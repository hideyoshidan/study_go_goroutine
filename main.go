package main

import (
	"fmt"
	"sync"
	"time"

	"goroutine.com/apis"
	"goroutine.com/client"
	"goroutine.com/constant"
	"goroutine.com/mutex"
	"goroutine.com/worker"
)

var wg sync.WaitGroup

func main() {
	// c := client.NewClient()

	// Nomal API without goroutine
	// executeNomal(c)
	// executeGoroutine(c)
	// executeRoutineWithChan(c)
	// executeRoutineWithChanAndSymbol(c)
	// executeRoutineWithSelect(c)
	// executeRoutineWithChan2(c)
	// executeWithTicker(c)
	// executeWorker(c)
	// executeWithWaitGroup(c, true)
	executeWithMutex(true)

}

// 普通に実行した場合
func executeNomal(c *client.Client) {
	fmt.Printf("==============%s\n", "Nomal")
	n := apis.NewNomal(c)
	n.ExecuteAPI1()
	n.ExecuteAPI2()
	n.ExecuteAPI3()
	n.ExecuteAPI4()
	n.ExecuteAPI5()
	n.ExecuteAPI6()
	n.ExecuteAPI7()
}

// Go routineを使った場合
func executeGoroutine(c *client.Client) {
	fmt.Printf("==============%s\n", "Go routine")
	n := apis.NewNomal(c)
	go n.ExecuteAPI1()
	go n.ExecuteAPI2()
	go n.ExecuteAPI3()
	go n.ExecuteAPI4()
	go n.ExecuteAPI5()
	go n.ExecuteAPI6()
	go n.ExecuteAPI7()
}

// channelを使った場合
func executeRoutineWithChan(c *client.Client) {
	fmt.Printf("==============%s\n", "Channel")
	w := apis.NewWithChannel(c)

	ga := make(chan string, 2) // Group A
	gb := make(chan string, 2) // Group B
	gc := make(chan string, 2) // Group C
	gd := make(chan string, 2) // Group D

	defer close(ga)
	defer close(gb)
	defer close(gc)
	defer close(gd)

	go w.ExecuteAPI1(ga)
	go w.ExecuteAPI2(ga)
	go w.ExecuteAPI3(gb)
	go w.ExecuteAPI4(gb)
	go w.ExecuteAPI5(gc)
	go w.ExecuteAPI6(gc)
	go w.ExecuteAPI7(gd)

	w.GetAPIGroup("A", ga)
	w.GetAPIGroup("A", ga)

	w.GetAPIGroup("B", gb)
	w.GetAPIGroup("B", gb)

	w.GetAPIGroup("C", gc)
	w.GetAPIGroup("C", gc)

	w.GetAPIGroup("D", gd)
}

// 関数の引数に<-を使った場合
func executeRoutineWithChanAndSymbol(c *client.Client) {
	fmt.Printf("==============%s\n", "Channle With Symbol")
	w := apis.NewWithChannel(c)

	ga := make(chan string, 2) // Group A
	gb := make(chan string, 2) // Group B
	gc := make(chan string, 2) // Group C
	gd := make(chan string, 2) // Group D

	defer close(ga)
	defer close(gb)
	defer close(gc)
	defer close(gd)

	go w.ExecuteAPI1WithSend(ga)
	go w.ExecuteAPI2WithSend(ga)
	go w.ExecuteAPI3WithSend(gb)
	go w.ExecuteAPI4WithSend(gb)
	go w.ExecuteAPI5WithSend(gc)
	go w.ExecuteAPI6WithSend(gc)
	go w.ExecuteAPI7WithSend(gd)

	w.GetAPIGroupWithRecive("A", ga)
	w.GetAPIGroupWithRecive("A", ga)

	w.GetAPIGroupWithRecive("B", gb)
	w.GetAPIGroupWithRecive("B", gb)

	w.GetAPIGroupWithRecive("C", gc)
	w.GetAPIGroupWithRecive("C", gc)

	w.GetAPIGroupWithRecive("D", gd)
}

// Selectを試す
func executeRoutineWithSelect(c *client.Client) {
	fmt.Printf("==============%s\n", "Channle With Select")
	w := apis.NewWithChannel(c)

	ga := make(chan string) // Group A
	gb := make(chan string) // Group B
	gc := make(chan string) // Group C
	gd := make(chan string) // Group D

	defer close(ga)
	defer close(gb)
	defer close(gc)
	defer close(gd)

	go func() {
		w.ExecuteAPI1WithSend(ga)
		w.ExecuteAPI2WithSend(ga)

	}()
	go func() {
		w.ExecuteAPI3WithSend(gb)
		w.ExecuteAPI4WithSend(gb)
	}()
	go func() {
		w.ExecuteAPI5WithSend(gc)
		w.ExecuteAPI6WithSend(gd)
	}()

	go w.ExecuteAPI7WithSend(gc)

	for i := 0; i < 7; i++ {
		select {
		case result := <-ga:
			fmt.Printf("GroupA : %v\n", result)
		case result := <-gb:
			fmt.Printf("GroupB : %v\n", result)
		case result := <-gc:
			fmt.Printf("GroupC : %v\n", result)
		case result := <-gd:
			fmt.Printf("GroupD : %v\n", result)
		}
	}
}

// Groupごとにgo routineを設定してみる
func executeRoutineWithChan2(c *client.Client) {
	fmt.Printf("==============%s\n", "Channel")
	w := apis.NewWithChannel(c)

	ga := make(chan string, 2) // Group A
	gb := make(chan string, 2) // Group B
	gc := make(chan string, 2) // Group C
	gd := make(chan string, 2) // Group D

	// go routineを分けると2回同じチャンネルを使いたいのに1回目でchannelがcloseされてしまう可能性がある。
	// なので1度のgo routineで同じgourpのものを実行
	go func() {
		w.ExecuteAPI1(ga)
		w.ExecuteAPI2(ga)
	}()

	go func() {
		w.ExecuteAPI3(gb)
		w.ExecuteAPI4(gb)
	}()

	go func() {
		w.ExecuteAPI5(gc)
		w.ExecuteAPI6(gc)
	}()

	go w.ExecuteAPI7(gd)

	for i := 0; i < 2; i++ {
		w.GetAPIGroup2("A", ga)
	}
	close(ga)
	_, ok := <-ga
	fmt.Printf("After G%s is not closed? : %v\n------------------\n", "A", ok)

	for i := 0; i < 2; i++ {
		w.GetAPIGroup2("B", gb)
	}
	close(gb)
	_, ok = <-gb
	fmt.Printf("After G%s is not closed? : %v\n------------------\n", "B", ok)

	for i := 0; i < 2; i++ {
		w.GetAPIGroup2("C", gc)
	}
	close(gc)
	_, ok = <-gc
	fmt.Printf("After G%s is not closed? : %v\n------------------\n", "C", ok)

	for i := 0; i < 1; i++ {
		w.GetAPIGroup2("D", gd)
	}
	close(gd)
	_, ok = <-gd
	fmt.Printf("After G%s is not closed? : %v\n------------------\n", "D", ok)
}

// 10秒間に1秒毎に実行
func executeWithTicker(c *client.Client) {
	ticker := time.NewTicker(1 * time.Second)
	n := apis.NewNomal(c)
	go func() {
		for t := range ticker.C {
			fmt.Println(t)
			n.ExecuteAPI1()
		}
	}()

	time.Sleep(10 * time.Second)
	ticker.Stop()
}

// Workerを使ってみる
func executeWorker(c *client.Client) {
	api := apis.NewForWorker(c)
	worker := worker.NewWorker(api)

	in := make(chan int)
	out := make(chan string)

	defer close(in)
	defer close(out)

	// maxWorkerNum分のwokerを起動
	maxWorkerNum := 5
	for i := 0; i < maxWorkerNum; i++ {
		go worker.Run(in, out)
	}

	// 数をchannel inに送る
	go func() {
		for i := 1; i < 11; i++ {
			in <- i
		}
	}()

	// 結果を取得し、標準出力
	for result := range out {
		fmt.Println(result)
	}
}

// WaitGroupのお試し
func executeWithWaitGroup(c *client.Client, useWait bool) {
	api := apis.NewNomal(c)
	cityCodes := []string{
		constant.MORIOKA,
		constant.SENDAI,
		constant.KANAZAWA,
		constant.NAGANO,
		constant.GIHU,
		constant.SHIZUOKA,
		constant.KYOTO,
	}

	if useWait {
		withWaitGroup(api, cityCodes)
	} else {
		noWaitGroup(api, cityCodes)
	}
}

func noWaitGroup(api *apis.Nomal, cityCodes []string) {
	var result []string
	res := make(chan string)

	fmt.Println("-----", "Start Routine without WaitGroup", "-----")
	for _, city := range cityCodes {
		go func(code string, res chan string) {
			res <- api.ExecuteAPI(code)
		}(city, res)
	}

	for i := 0; i < len(cityCodes); i++ {
		var format string
		if i == len(cityCodes)-1 {
			format = <-res
		} else {
			format = fmt.Sprintf("%s,", <-res)
		}
		result = append(result, format)
	}
	fmt.Println("-----", "Finish Routine without WaitGroup", "-----")
}

func withWaitGroup(api *apis.Nomal, cityCodes []string) {
	var wg sync.WaitGroup
	var result []string

	res := make(chan string)

	fmt.Println("-----", "Start Routine with WaitGroup", "-----")

	for _, city := range cityCodes {
		wg.Add(1)
		go func(code string, res chan string) {
			res <- api.ExecuteAPI(code)
			wg.Done()
		}(city, res)
	}

	for i := 0; i < len(cityCodes); i++ {
		var format string
		if i == len(cityCodes)-1 {
			format = <-res
		} else {
			format = fmt.Sprintf("%s,", <-res)
		}
		result = append(result, format)
	}

	wg.Wait()
	fmt.Println(result)
	fmt.Println("-----", "Finish Routine with WaitGroup", "-----")
}

func executeWithMutex(useMutex bool) {
	if useMutex {
		withMutex()
	} else {
		withoutMutex()
	}
}

// fatal error: concurrent map writes
// map型の変数に対して、書き込みが競合したことによるエラー
// go routineによって複数threadから同じkeyのvalueに対して書き込みをしようとして
// このエラーが発生
func withoutMutex() {
	woutM := mutex.NewWioutMutex()
	for i := 0; i < 1000; i++ {
		go woutM.Set("TEST")
	}
	fmt.Println(woutM.Get("TEST"))
}

// Mutexを使うと、排他的Lockをかけてくれる。
func withMutex() {
	wiM := mutex.NewWiMutex()
	for i := 0; i < 1000; i++ {
		go wiM.Set("TEST")
	}
	fmt.Println(wiM.Get("TEST"))
}
