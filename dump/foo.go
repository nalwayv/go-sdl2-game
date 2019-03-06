package main

/*
  CHANNELS
  --------
  channel's block

  range 'blocks' :: loops over values sent to a channel

  select :: pull values from given channel when a value is sent

  ok idiom :: check if a channel is closed
              open = true
              closed = false

              v,ok := <- c
              if !ok { closed }
              if  ok { open }

  close() :: closes a given channel.
             if not closed could cause 'deadlock' errors
             with range as it sits waiting for values.

  chan <- :: send only
  <- chan :: receive only
*/

import (
	"fmt"
)

// receive channel only
func bar(c <-chan int) {
	// if 'c is not closed range keeps waiting for values
	// that will never come so turns into a deadlock error
	for v := range c {
		fmt.Println(v)
	}
}

func fooChan() {
	c := make(chan int)

	go func(c chan<- int) {
		for i := 0; i < 100; i++ {
			c <- i * 2
		}
		close(c)
	}(c)

	bar(c) // blocking till foo finishes

	fmt.Println("Exit")

}

func selectChan() {
	/*
	   Using send to catch values sent to different channel's
	*/
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	// send
	go func(even, odd chan<- int, quit chan<- bool) {

		for num := 0; num < 10; num++ {
			if num&1 == 0 {
				// EVEN
				even <- num
			} else {
				// ODD
				odd <- num
			}
		}

		// stop sending data
		close(even)
		close(odd)
		close(quit)

	}(even, odd, quit)

	// receive
	func(even, odd <-chan int, quit <-chan bool) {
		for {
			select {
			case v, ok := <-even:
				// if open
				if ok {
					fmt.Println("EVEN\t", v)
				}
			case v, ok := <-odd:
				// if open
				if ok {
					fmt.Println("ODD\t", v)
				}
			case v, ok := <-quit:
				// is closed
				if !ok {
					fmt.Println("A QUIT\t", v, ok)
					return
				}
			}
		}
	}(even, odd, quit)
}

func main() {
	//fooChan()
	selectChan()
}
