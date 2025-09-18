package main
import (
	"fmt"
	"sync"
)
func test(wg *sync.WaitGroup, done <-chan string, id int) {
	defer wg.Done()
		if len(done)==0{
			return
		}
		fmt.Println(<-done,id)
}

func main() {
	var wg sync.WaitGroup
	done := make(chan string,1)
	for i := 1; i < 10; i++ {
		wg.Add(1)
		done <- "Горутинка"
		go test(&wg, done, i)
	}
	wg.Wait()
	fmt.Println("Все горутинки завершены")
}
