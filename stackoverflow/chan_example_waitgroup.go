package main

import (
	"fmt"
	"sync"
	"time"
)

func UseWaitGroupControlFlow() {
	// 声明等待组
	var wg sync.WaitGroup
	// 设置，需要等待3个协程执行完成
	wg.Add(3)
	// 创建通道
	intChan := make(chan int)
	// 计算1-50的和
	go func(intChan chan int, wg *sync.WaitGroup) {
		sum := 0
		for i := 1; i <= 50; i++ {
			sum += i
		}
		intChan <- sum
		// 计数器减一
		wg.Done()
	}(intChan, &wg)
	// 计算51-100的和
	go func(intChan chan int, wg *sync.WaitGroup) {
		sum := 0
		for i := 51; i <= 100; i++ {
			sum += i
		}
		intChan <- sum
		// 计数器减一
		wg.Done()
	}(intChan, &wg)
	// 另外创建个channle聚合结果
	go func(intChan chan int, wg *sync.WaitGroup) {
		sum1 := <-intChan
		sum2 := <-intChan
		fmt.Printf("sum1 = %d sum2 = %d  \nsum1 + sum2 = %d \n", sum1, sum2, sum1+sum2)
		// 计数器减一
		wg.Done()
	}(intChan, &wg)
	// 阻塞，直到等待组的计数器等于0
	wg.Wait()
	fmt.Println("运行结束")
}

// UseChanReplaceWaitGroupControlFlow 使用chan来替换WaitGroup
func UseChanReplaceWaitGroupControlFlow() {
	//声明等待组
	goroutine := 4
	waitGroup := make(chan string, goroutine)
	// 创建通道
	intChan := make(chan int)
	// 计算1-50的和
	go func(intChan chan int, wg chan<- string) {
		sum := 0
		for i := 1; i <= 50; i++ {
			sum += i
		}
		intChan <- sum
		// 计数器减一
		wg <- "1-50"
	}(intChan, waitGroup)
	// 计算51-100的和
	go func(intChan chan int, wg chan<- string) {
		sum := 0
		for i := 51; i <= 100; i++ {
			sum += i
		}
		intChan <- sum
		// 计数器减一
		wg <- "51-100"
	}(intChan, waitGroup)
	go func(intChan chan int, wg chan<- string) {
		sum := 0
		for i := 101; i <= 200; i++ {
			sum += i
		}
		time.Sleep(time.Second)
		intChan <- sum
		// 计数器减一
		wg <- "101-200"
	}(intChan, waitGroup)
	// 另外创建个channle聚合结果
	go func(intChan chan int, wg chan<- string) {
		sum1 := <-intChan
		sum2 := <-intChan
		sum3 := <-intChan
		fmt.Printf("sum1 = %d sum2 = %d sum3 = %d  \nsum1 + sum2 +sum3= %d \n", sum1, sum2, sum3, sum1+sum2+sum3)
		// 计数器减一
		wg <- "sumValue"
	}(intChan, waitGroup)

	//等待组
	for i := 0; i < goroutine; i++ {
		result := <-waitGroup
		fmt.Println(result, "运行结束")
	}
	fmt.Println("运行结束")
}

func main() {
	//UseWaitGroupControlFlow()
	UseChanReplaceWaitGroupControlFlow()
}

/**输出
  sum1 = 3775 sum2 = 1275
  sum1 + sum2 = 5050
  运行结束
*/
