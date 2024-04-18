package main

var mMap map[int]int

func TestMyMap() {
	mMap = make(map[int]int)

	for i := 0; i < 5000; i++ {
		go func() {
			mMap[i] = i
		}()
		go readMap(i)
	}
}
func readMap(i int) int {
	return mMap[i]
}
