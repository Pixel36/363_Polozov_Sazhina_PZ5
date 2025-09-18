package main

import (
 "fmt"
 "time"
)

func longRequest(result chan<- string) {
 time.Sleep(3 * time.Second)
 result <- "Ответ от сервера"
}

func main() {
 result := make(chan string)

 go longRequest(result)

 select {
 case r := <-result:
  fmt.Println("Получено:", r)
 case <-time.After(2 * time.Second):
  fmt.Println("Таймаут: сервер не ответил за 2 секунды")
 }
}
