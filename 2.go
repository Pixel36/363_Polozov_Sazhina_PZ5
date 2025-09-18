package main

import (
 "fmt"
 "time"
)

func rabota(done <-chan struct{}) {
 ticker := time.NewTicker(1 * time.Second)
 defer ticker.Stop()
 for {
  select {
  case <-done:
   fmt.Println("rabota: получил сигнал остановки, завершаюсь")
   return
  case t := <-ticker.C:
   fmt.Println("rabota:", t.Format("15:04:05"))
  }
 }
}

func main() {
 done := make(chan struct{})

 go rabota(done)

 time.Sleep(5 * time.Second)
 close(done)

 time.Sleep(500 * time.Millisecond)
 fmt.Println("main: программа завершена")
}
