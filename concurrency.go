package main
import (
  "fmt"
  "time"
  "math/rand"
)

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func pinger(c chan string) {
		for i := 0; ; i++ {
			c <- "ping"
		}
}

func printer(c chan string) {
	for {
		msg := <- c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func ponger(c chan string)  {
  for i := 0; ; i++ {
    c <- "pong"
  }
}

func main()  {
  // buffered channels
  // its also posible to pass a second parameter to the make function when creating a channel
  // a buffered channel is asyncronous: sending or receiving a message will not wait unless the channel is already full
  d := make(chan int, 1) // this create a buffered  channel with a capaccity of 1.

  /*
  // goroutine
  for i := 0; i < 10; i++ {
		go f(i)
	}
  go f(0)
  */
  /*
  // cahnnels
  // normally channels are syncronous, both side of the channel will wait until the other side is ready
  var c chan string = make(chan string)

	go pinger(c)
  go ponger(c) //will print pong forever
	go printer(c) //will print ping forever
  */
  //select for channels
  c1 := make(chan string)
  c2 := make(chan string)
  go func() {
    for {
      c1 <- "from 1"
      time.Sleep(time.Second * 2)
    }
  }()

  go func() {
    for {
      c2 <- "from 2"
      time.Sleep(time.Second * 3)
    }
  }()

  go func() {
    for {
      select { // select statement is often used to implement a timeout, work like switch but for channels
      case msg1 := <- c1:
        fmt.Println(msg1)
      case msg2 := <- c2:
        fmt.Println(msg2)
      case <- time.After(time.Second):
        //time.After create a channel and after the given duration will send the current time on it
        fmt.Println("timeout")
      /*
      default:
        // we can also specify a default case:
        // if none of the channels are eady
        fmt.Println("nothing ready")
      */
      }
    }
  }()

  var input string
	fmt.Scanln(&input)
  /*
          output goroutine
          0 : 0
          0 : 1
          0 : 2
          0 : 3
          0 : 4
          0 : 5
          0 : 6
          0 : 7
          0 : 8
          0 : 9
          ping
          pong
          ping
          pong
          ping //hit enter to stop
  */
  /*
  select result
        from 1
        timeout
        from 2
        from 1
        timeout
        timeout
        from 1
        from 2
        timeout
        from 1
        timeou
  */
}
