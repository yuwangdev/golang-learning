package other3

import (
	"fmt"
	"time"
	"sync"
)

// use channel to force this functon run before the main exits
func asyncHello(i int, ch chan int) {
	fmt.Printf("Async Hello %d\n", i)
	ch <- i
	time.Sleep(10 * time.Microsecond)
}

func snedToCh(ch chan int) {
	i := 0
	for ; i < 3; i++ {
		ch <- i
	}
	close(ch)
}

func snedToBufferedCh(ch chan int) {
	i := 0
	for ; i < 6; i++ {
		fmt.Printf("put %d\n", i)
		ch <- i
	}
	close(ch)
}

// the channel beauty is, for a async go method, put the result into the channel, and pop it when using it
func ExportForConcurrency() {

	fmt.Println("----------Concurrency Programming Starts Here")

	// use channel as block control tool
	var ch = make(chan int)
	for i := 1; i < 3; i++ {
		go asyncHello(i, ch)
	}
	fmt.Println(len(ch))
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// consumer and producer mode
	var ch2 = make(chan int)
	go snedToCh(ch2)

	// for v:=range(ch2) also works
	for {
		v, ok := <-ch2
		if ok {
			fmt.Println(v)
		} else {
			break
		}
	}

	// create buffered chanel
	// buffered chanel is locked when eithr full or empty, not locked for reading/sending
	var bch = make(chan int, 3)
	go snedToBufferedCh(bch)
	for v := range bch {
		fmt.Println("get ...")
		fmt.Println(v)
	}

	// mutex
	globalVar := 1
	var mutex = sync.Mutex{}
	mutex.Lock()
	globalVar = 3
	mutex.Unlock()
	fmt.Println(globalVar)

	fmt.Println("----------Concurrency Programming Ends Here")

}
