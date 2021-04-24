package main

import "fmt"
// import "time"

func TestWithoutDefault() {
      var c1, c2, c3 chan int
   var i1, i2 int

   c1 = make(chan int, 100)

   go func() {
      c1 <- 10
   }()

   // 没有default语句，阻塞住直到消费消息为止
   select {
      case i1 = <- c1:
         fmt.Printf("received %d from c1\n", i1)

      case c2 <- i2:
         fmt.Printf("sent ", i2, " to c2\n")
      
      case i3, ok := (<-c3):  // same as: i3, ok := <-c3
         if ok {
            fmt.Printf("received ", i3, " from c3\n")
         } else {
            fmt.Printf("c3 is closed\n")
         }
   }   
}

func TestWithDefault() {
   c1 := make(chan int, 100)
   var i1 int;

   go func() {
      c1 <- 10
   }()

   select {
   case i1 = <-c1:
      fmt.Printf("received %d\n", i1);
   default:
      fmt.Printf("no received\n")
   }
}

func main() {
   TestWithoutDefault()
   TestWithDefault()
}

