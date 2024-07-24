package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var sum = 0
var arrayInt = []int{}

// var mutex sync.Mutex
var mutexrw = sync.RWMutex{}
var mutex = sync.WaitGroup{}

func main() {
	//run()
	//doOnce()
	//cond()
	contextSync()
}

func contextSync() {
	mutex.Add(1)
	//stopCh := make(chan bool)
	ctx, stop := context.WithCancel(context.Background())
	str := "1123123"
	go func() {
		defer mutex.Done()
		fmt.Println("监控狗1启动")
		fmt.Println(str)
		//watchDog("监控狗1....", stopCh)
		watchDog("监控狗1....", ctx)
	}()
	time.Sleep(5 * time.Second)
	fmt.Println(ctx.Err())
	stop()
	//stopCh <- true
	mutex.Wait()
}

func watchDog(name string, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "指令已停止")
			return
		default:
			fmt.Println(name, "正在监控")
		}
		time.Sleep(1 * time.Second)
	}
}

func cond() {
	cond := sync.NewCond(&mutexrw)
	mutex.Add(10)
	for i := 1; i <= 10; i++ {
		go func(num int) {
			defer mutex.Done()
			fmt.Println(num, "号已就位")
			cond.L.Lock()
			cond.Wait()
			fmt.Println(num, "号开始跑~~~")
			cond.L.Unlock()
		}(i)
	}

	time.Sleep(2 * time.Second)
	go func() {
		fmt.Println("准备开始")
		//cond.Broadcast()
		for i := 1; i <= 10; i++ {
			cond.Signal()
		}
	}()
	mutex.Wait()
}

func doOnce() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("only once")
	}
	done := make(chan bool)
	defer close(done)
	for i := 1; i <= 10; i++ {
		go func() {
			once.Do(onceBody)
			done <- true
			//fmt.Println("i1:", i)
		}()
	}
	for i := 1; i <= 10; i++ {
		<-done
		//fmt.Println("i2:", i)
	}
}

func run() {
	mutex.Add(3000)
	for i := 1; i <= 1000; i++ {
		go add(10)
	}

	for i := 1; i <= 2000; i++ {
		go read()
	}

	mutex.Wait()
	fmt.Println("sum:", max())
	fmt.Println("sum:", sum)
}

func add(i int) {
	//mutex.Lock()
	mutexrw.Lock()
	sum += i
	defer mutex.Done()
	mutexrw.Unlock()
	//mutex.Unlock()
}

func read() {
	mutexrw.RLock()
	arrayInt = append(arrayInt, sum)
	defer mutex.Done()
	mutexrw.RUnlock()
}

func max() int {
	if len(arrayInt) == 0 {
		return 0
	}

	max := arrayInt[0]
	for _, value := range arrayInt[1:] {
		if value > max {
			max = value
		}
	}
	return max
}
