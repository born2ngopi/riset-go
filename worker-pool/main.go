package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

type Worker struct {
	TotalWorker int
	Wg          *sync.WaitGroup
	TaskC       chan func()
}

type result struct {
	ID int
}

func (w *Worker) AddTask(task func()) {
	w.TaskC <- task
}

func (w *Worker) Run() {
	for i := 0; i < w.TotalWorker; i++ {
		go func(index int) {
			for task := range w.TaskC {
				log.Printf("stated worker : %d\n", index)
				task()
				log.Printf("finished worker : %d\n", index)
				w.Wg.Done()
			}
		}(i)
	}
}

func main() {

	var totalWorker, totalTask int = 3, 5

	wg := &sync.WaitGroup{}
	wg.Add(totalTask)

	worker := &Worker{
		Wg:          wg,
		TotalWorker: totalWorker,
		TaskC:       make(chan func()),
	}
	worker.Run()

	var resC = make(chan result, totalTask)

	for i := 0; i < totalTask; i++ {
		worker.AddTask(func() {
			id := rand.Int()

			log.Printf("starting task %d\n", id)
			time.Sleep(time.Second)
			log.Printf("finished task %d\n", id)

			resC <- result{ID: id}
		})
	}

	wg.Wait()

	var respA = []result{}

	for i := 0; i < totalTask; i++ {
		resp := <-resC
		respA = append(respA, resp)
	}

	fmt.Println(respA)

}
