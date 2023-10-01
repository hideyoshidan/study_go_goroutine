package worker

import (
	"fmt"
	"time"

	"goroutine.com/apis"
)

type Worker struct {
	*apis.ForWorker
}

func NewWorker(a *apis.ForWorker) *Worker {
	return &Worker{a}
}

func (w *Worker) Run(in <-chan int, out chan<- string) {
	for {
		number := <-in
		time.Sleep(1 * time.Second)
		result := fmt.Sprintf("%d回目の結果 : %s", number, w.ExecuteAPI())
		out <- result
	}
}
