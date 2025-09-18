package main

import (
 "fmt"
 "time"
)

type Job struct {
 Name string
}

func rabota(high <-chan Job, low <-chan Job) {
 for (len(high)>0) {
  fmt.Println(<-high)
 }
 for (len(low)>0) {
 fmt.Println(<-low)
 }
}

func main() {
 high := make(chan Job, 5)
 low := make(chan Job, 5)

 go rabota(high, low)

 low <- Job{"джоб поменьбше-1"}
 low <- Job{"джоб поменьбше-2"}
 high <- Job{"важный джоб-1"}
 low <- Job{"джоб поменьбше-3"}
 high <- Job{"важный джоб-2"}

 time.Sleep(1 * time.Second)
}
