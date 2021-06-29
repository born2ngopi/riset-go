# Worker Pool

----

### Intro

Pada bahasa pemrograman golang, kita tau bahwa kita dapat membuat concurrency dengan goroutine. Pada tulisan kali ini kita akan membahas dan membuat **worker pool** sederhana, lalu apa itu worker pool ?. Worker pool atau thread pool itu beberapa goroutine yang siap gantian ngambil antrian tugas dari channel task dan ketika selesai akan dikirim ke channel result.

![](/home/rizky/development/go/src/github.com/needkopi/riset-go/worker-pool/Group 1.png)

Lalu kenapa kita butuh worker pool?, tentunya kita tidak memliki resource yang tidak terbatas pada komputer/server kita, kalau kita membuat banyak goroutine itu akan menghabiskan resource memory dan CPU. Dengan memakai limited worker pool  kita dapat meminimalisir hal tersebut

### Implement

Pertama kita akan membuat struct `Worker` 

``` go
type Worker struct {
    TotalWorker int
    Wg 			*sync.WaitGroup
    TaskC 		chan func()
}
```

`TotalWorker` digunakan untuk mendefinisikan berapa jumlah worker yang akan kita buat, `TaskC` adalah channel task yang akan kita eksekusi.

Selanjutnya kita akan membuat method untuk menambakan task ke channel

``` go
func (w *Worker) AddTask(task func()) {
	w.TaskC <- task
}
```

Lalu kita akan membuat method `Run()` yang di gunakan untuk membuat worker sebanyak jumlah yang kita tentukan pada `TotalWorker`

``` go
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
```

Selanjutnya pada fungsi `main` kita akan membuat instance dari struct `Worker`, sebelum kita define dulu jumlah worker dan jumlah task yang akan kita buat. Lalu kita panggil method `Run()` untuk menjalankan worker sebanyak yang kita define sebelumnya.

``` go
var totalWorker, totalTask int = 3, 5

wg := &sync.WaitGroup{}
wg.Add(totalTask)

worker := &Worker{
    Wg:          wg,
    TotalWorker: totalWorker,
    TaskC:       make(chan func()),
}
worker.Run()
```

Setelah kita buat channel result, dan kita tambah task ke dalam worker

``` go
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
```

Kalau kita running akan akan seperti ini :

``` bash
$ time go run main.go
2021/06/29 17:19:59 stated worker : 2
2021/06/29 17:19:59 starting task 5577006791947779410
2021/06/29 17:19:59 stated worker : 1
2021/06/29 17:19:59 starting task 8674665223082153551
2021/06/29 17:19:59 stated worker : 0
2021/06/29 17:19:59 starting task 6129484611666145821
2021/06/29 17:20:00 finished task 5577006791947779410
2021/06/29 17:20:00 finished worker : 2
2021/06/29 17:20:00 stated worker : 2
2021/06/29 17:20:00 starting task 4037200794235010051
2021/06/29 17:20:00 finished task 8674665223082153551
2021/06/29 17:20:00 finished worker : 1
2021/06/29 17:20:00 stated worker : 1
2021/06/29 17:20:00 starting task 3916589616287113937
2021/06/29 17:20:00 finished task 6129484611666145821
2021/06/29 17:20:00 finished worker : 0
2021/06/29 17:20:01 finished task 3916589616287113937
2021/06/29 17:20:01 finished worker : 1
2021/06/29 17:20:01 finished task 4037200794235010051
2021/06/29 17:20:01 finished worker : 2

[{5577006791947779410} {8674665223082153551} {6129484611666145821} {3916589616287113937} {4037200794235010051}]

real    0m2,148s
user    0m0,201s
sys     0m0,094s

```

Lama waktu aplikasi jalan sekitar 2 detik yang seharusnya jika tanpa worker pool applikasi akan jalan selama 5 detik karna kita menggunakan `time.Sleep` selama satu detik. Worker pool ini juga sangan berguna jika kita implementasikan ke cron job.

