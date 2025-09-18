package main

import (
 "fmt"
 "time"
)

type Job struct {
 Name string
}

func rabota(high <-chan Job, low <-chan Job) {
 for {
  select {
  case j := <-high:
   fmt.Println("важный джоб:", j.Name)
  case j := <-low:
   fmt.Println("джоб поменьбше:", j.Name)
  default:
   time.Sleep(100 * time.Millisecond)
  }
 }
}

func main() {
 high := make(chan Job, 5)
 low := make(chan Job, 5)

 go rabota(high, low)

 low <- Job{"п-1"}
 low <- Job{"п-2"}
 high <- Job{"в-1"}
 low <- Job{"п-3"}
 high <- Job{"в-2"}

 time.Sleep(1 * time.Second)
}
