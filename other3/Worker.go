package other3

import (
	"fmt"
	"sync"
	"time"
)

/*
as long as concurrent concept, it uses GO as must
 */

var globalResource int = 1

func processWaited(i int, wg *sync.WaitGroup) {
	fmt.Printf("work on %d\n", i)
	mutex := sync.Mutex{}
	mutex.Lock()
	globalResource += i
	mutex.Unlock()
	wg.Done()
}

func server1(ch chan int) {
	time.Sleep(3 * time.Second)
	ch <- 1
}

func server2(ch chan int) {
	time.Sleep(6 * time.Second)
	ch <- 2
}

// the channel beauty is, for a async go method, put the result into the channel, and pop it when using it
func ExportForConcurrencyWorker() {

	// waiting for execution worker
	var wg sync.WaitGroup = sync.WaitGroup{}
	for i := 0; i < 4; i++ {
		wg.Add(i)
		go processWaited(i, &wg) // need to use the same wg for global
	}

	wg.Wait()

	// for select
	//
	chan1 := make(chan int)
	chan2 := make(chan int)
	go server1(chan1)
	go server2(chan2)
	select {
	case data1 := <-chan1:
		{
			fmt.Printf("from chan 1 with the data of %d\n", data1)
		}
	case data2 := <-chan2:
		{
			fmt.Printf("from chan 2 with the data of %d\n", data2)
		}
	default: // ensure no deadlock
		fmt.Println("no data received")
	}

}
