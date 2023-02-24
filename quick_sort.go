package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	// fmt.Println(Sort([]int{5, 786, 23, 5, 342, 7, 8, 987, 9, 0, 3, 547, 567, 4566, 75, 3456, 46, 4356, 345, 76870, 56, 7, 546, 78}, true))
	// fmt.Println(Sort([]int{547, 286, 75, 3456, 512, 4356, 345, 56}, true))
	// fmt.Println(Sort([]int{12, 34, 23, 45, 547, 286}))

	rand.Seed(time.Now().UnixNano())
	co := 100000
	s := make([]int, 0, co)
	for i := 0; i < co; i++ {
		s = append(s, int(rand.Uint32()))
	}
	t := time.Now()
	Sort(s)
	fmt.Println(float64(time.Since(t)/1000) / 1000)
	good := 1
	for i := 1; i < co; i++ {
		if s[i-1] <= s[i] {
			good++
		}
	}
	fmt.Println(good)
}

const AsyncLimitEnvName = "QUICK_SORT_ASYNC_LIMIT"

func loadAsyncLimit() int {
	def := 1000
	str := os.Getenv(AsyncLimitEnvName)
	if str == "" {
		return def
	}
	if lim, err := strconv.Atoi(str); err == nil {
		return lim
	}
	return def
}

func Sort(a []int, asyncLimit ...int) []int {
	if len(a) < 2 {
		return a
	}

	if len(asyncLimit) == 0 {
		asyncLimit = append(asyncLimit, loadAsyncLimit())
	}

	i, j := 0, len(a)-1
	k := a[j>>1]
	for i <= j {
		for ; a[i] < k; i++ {
		}
		for ; a[j] > k; j-- {
		}
		if i <= j {
			a[i], a[j] = a[j], a[i]
			i++
			j--
		}
	}

	var wg *sync.WaitGroup
	lSize, rSize := j+1, len(a)-i
	if lSize > asyncLimit[0] || rSize > asyncLimit[0] {
		wg = new(sync.WaitGroup)
	}
	if j > 0 {
		if lSize > asyncLimit[0] {
			sortAsync(a[:j+1], asyncLimit, wg)
		} else {
			Sort(a[:j+1], asyncLimit...)
		}
	}
	if i < len(a)-1 {
		if rSize > asyncLimit[0] {
			sortAsync(a[i:], asyncLimit, wg)
		} else {
			Sort(a[i:], asyncLimit...)
		}
	}
	if lSize > asyncLimit[0] || rSize > asyncLimit[0] {
		wg.Wait()
	}

	return a
}

var v func()

func sortAsync(a, asyncLimit []int, wg *sync.WaitGroup) {
	wg.Add(1)
	if v != nil {
		v()
	}
	go func() {
		defer wg.Done()
		Sort(a, asyncLimit...)
	}()
}
