package main

import (
	"fmt"
	"hash/fnv"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

/**
哲学家吃饭的问题
*/

type PH struct {
	name  string
	left  Fork
	right Fork
}

type Fork struct {
	fork *sync.Mutex
	name string
}

// Number of philosophers is simply the length of this list.
var ph = []PH{{name: "Mark"}, {name: "Russell"}, {name: "Rocky"}, {name: "Haris"}, {name: "Root"}}

const hunger = 3                // Number of times each philosopher eats
const think = time.Second / 100 // Mean think time
const eat = time.Second / 100   // Mean eat time
var dining sync.WaitGroup

func diningProblem(p PH) {
	dominantHand := p.left.fork
	otherHand := p.right.fork
	phName := p.name
	//大家都能够坐下
	fmt.Println(phName, "Seated")
	h := fnv.New64a()
	h.Write([]byte(phName))
	rg := rand.New(rand.NewSource(int64(h.Sum64())))
	rSleep := func(t time.Duration) {
		time.Sleep(t/2 + time.Duration(rg.Int63n(int64(t))))
	}
	for h := hunger; h > 0; h-- {
		fmt.Println(phName, "Hungry")
		dominantHand.Lock() // pick up forks
		otherHand.Lock()
		fmt.Println(phName, "Eating")
		rSleep(eat)
		dominantHand.Unlock() // put down forks
		otherHand.Unlock()
		fmt.Println(phName, "Thinking")
		rSleep(think)
	}
	fmt.Println(phName, "Satisfied")
	dining.Done()
	fmt.Println(phName, "Left the table")
}
func main() {
	fmt.Println("begin")
	dining.Add(5) // 等待模拟5个人吃饭，思考的协程结束

	forks := [5]Fork{}
	for i := 1; i <= len(forks); i++ {
		forks[i-1] = Fork{
			fork: &sync.Mutex{},
			name: "" + strconv.Itoa(i),
		}
	}

	for i := 0; i < len(ph); i++ {
		ph[i].left = forks[i%len(ph)]
		ph[i].right = forks[(i+1)%len(ph)]
	}

	for i := 0; i < len(ph); i++ {
		go diningProblem(ph[i])
	}

	dining.Wait() // 等待模拟5个人吃饭，思考的协程结束
	fmt.Println("Over")
}
