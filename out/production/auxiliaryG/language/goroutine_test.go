package gotutorial

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func responseSize(url string) {
	fmt.Println("Step1: ", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Step2: ", url)
	defer response.Body.Close()
	fmt.Println("Step3: ", url)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Step4: ", len(body))
}

func TestGoRoutines(t *testing.T) {
	go responseSize("https://www.golangprograms.com")
	go responseSize("https://coderwall.com")
	go responseSize("https://stackoverflow.com")
	time.Sleep(10 * time.Second)
	fmt.Println("test main goroutines")
}

// WaitGroup is used to wait for the program to finish goroutines.
var wg sync.WaitGroup

func responseSizeInSync(url string) {
	// Schedule the call to WaitGroup's Done to tell goroutine is completed.
	defer wg.Done()
	fmt.Println("Step1: ", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Step2: ", url)
	defer response.Body.Close()
	fmt.Println("Step3: ", url)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Step4: ", len(body))
}
func TestSync(t *testing.T) {
	// Add a count of three, one for each goroutine.
	wg.Add(3)
	fmt.Println("Start Goroutines")
	go responseSizeInSync("https://www.golangprograms.com")
	go responseSizeInSync("https://stackoverflow.com")
	go responseSizeInSync("https://coderwall.com")
	// Wait for the goroutines to finish.
	wg.Wait()
	fmt.Println("Terminating Program")
}

func responseSizeSyncChannel(url string, nums chan int) {
	// Schedule the call to WaitGroup's Done to tell goroutine is completed.
	defer wg.Done()
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	// Send value to the unbuffered channel
	nums <- len(body)
}

func TestSyncDataFromGoroutine(t *testing.T) {
	nums := make(chan int) // Declare a unbuffered channel
	wg.Add(1)
	go responseSizeSyncChannel("https://www.golangprograms.com", nums)
	fmt.Println(<-nums) // Read the value from unbuffered channel
	wg.Wait()
	close(nums) // Closes the channel
}

var i int

func work() {
	time.Sleep(250 * time.Millisecond)
	i++
	fmt.Println(i)
}
func routine(command <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	var status = "Play"
	for {
		select {
		case cmd := <-command:
			fmt.Println(cmd)
			switch cmd {
			case "Stop":
				return
			case "Pause":
				status = "Pause"
			default:
				status = "Play"
			}
		default:
			if status == "Play" {
				work()
			}
		}
	}
}

func TestPlayAndPauseExecutionOfGoroutine(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	command := make(chan string)
	go routine(command, &wg)
	time.Sleep(1 * time.Second)
	command <- "Pause"
	time.Sleep(1 * time.Second)
	command <- "Play"
	time.Sleep(1 * time.Second)
	command <- "Stop"
	wg.Wait()
}

var (
	counter int32          // counter is a variable incremented by all goroutines.
	wg1     sync.WaitGroup // wg is used to wait for the program to finish.
)

func increment(name string) {
	defer wg1.Done() // Schedule the call to Done to tell main we are done.
	for range name {
		atomic.AddInt32(&counter, 1)
		runtime.Gosched() // Yield the thread and be placed back in queue.
	}
}

func TestAtomic(t *testing.T) {
	wg.Add(3) // Add a count of two, one for each goroutine.
	go increment("Python")
	go increment("Java")
	go increment("Golang")
	wg.Wait() // Wait for the goroutines to finish.
	fmt.Println("Counter:", counter)
}

var (
	counter2 int32          // counter is a variable incremented by all goroutines.
	wg2      sync.WaitGroup // wg is used to wait for the program to finish.
	mutex    sync.Mutex     // mutex is used to define a critical section of code.
)

func increment2(lang string) {
	defer wg2.Done() // Schedule the call to Done to tell main we are done.
	for i := 0; i < 3; i++ {
		mutex.Lock()
		{
			fmt.Println(lang)
			counter2++
		}
		mutex.Unlock()
	}
}

func TestMutex(t *testing.T) {
	wg2.Add(3) // Add a count of two, one for each goroutine.
	go increment2("Python")
	go increment2("Go Programming Language")
	go increment2("Java")
	wg2.Wait() // Wait for the goroutines to finish.
	fmt.Println("Counter:", counter2)
}

var (
	mu    = sync.Mutex{}
	cond  = sync.NewCond(&mu)
	cond2 = sync.NewCond(&mu)
	index int
	state = true
)

func TestAlternatelyPrint(t *testing.T) {
	maxV := 2000
	go print1(maxV, cond)
	go print2(maxV, cond)
	time.Sleep(time.Second * 10000)
}

func print1(v int, cond *sync.Cond) {
	for index < v {
		cond.L.Lock()
		if state {
			fmt.Println("p1:", index)
			index++
			state = false
		}
		cond.L.Unlock()
	}
}

func print2(v int, cond *sync.Cond) {
	for index < v {
		cond.L.Lock()
		if !state {
			fmt.Println("p2:", index)
			index++
			state = true
		}
		cond.L.Unlock()
	}
}
