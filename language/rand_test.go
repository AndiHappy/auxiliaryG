package gotutorial

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
)

func TestRand(t *testing.T) {
	rand.NewSource(1000)
	var res, step string
	for i := 0; i < 100; i++ {
		res += step + strconv.Itoa(rand.Intn(100))
		step = " "
	}
	fmt.Println(res)
}
