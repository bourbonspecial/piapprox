package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const printOn = 1e8

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	in := 0
	cnt := 0

	mutex := &sync.Mutex{}

	for i:=0;i<4;i++ {
		go func() {
			for {
				cnt += 1
				x, y := (2 * rand.Float64()) - 1, (2 *rand.Float64()) - 1

				if x*x + y*y < 1 {
					in += 1
				}

				if cnt % printOn == 0 {
					mutex.Lock()
					fmt.Println(float64(in) * 4. / float64(cnt), cnt)
					mutex.Unlock()
				}
			}
		}()
	}
}