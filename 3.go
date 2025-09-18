package main

import "fmt"

func main() {
 ch := make(chan int, 1)
 ch <- 2
 select {
 case v := <-ch:
  fmt.Println("Данные:", v)
 default:
  fmt.Println("Данных нет")
 }
}
