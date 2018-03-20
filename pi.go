package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const printOn = 1e7

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	in := 0
	cnt := 0

	mutex := &sync.Mutex{}

	res := make(chan float64)
	chanCnt := make(chan int)

	for i:=0;i<4;i++ {
		go func() {
			lcnt := 0
			lin := 0
			for {
				lcnt += 1
				x, y := (2 * rand.Float64()) - 1, (2 *rand.Float64()) - 1

				if x*x + y*y < 1 {
					lin += 1
				}

				if lcnt % printOn == 0 {
					mutex.Lock()
					cnt += lcnt
					in += lin
					res <- float64(in) * 4. / float64(cnt)
					chanCnt <- cnt
					mutex.Unlock()
				}
			}
		}()
	}

	for {
		fmt.Println(<- res, <- chanCnt)
	}
}