package snippets

import (
	"fmt"
)

/*
func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			worker(i)
		}(i)
	}

		for j := 1; j <= 10; j++ {
			go func(j int) {
				defer wg.Done()
				worker(j)
			}(j)
		}

	wg.Wait()
}

*/

func Concurrency() {
	num_jobs := 10
	jobs := make(chan int, num_jobs)
	results := make(chan int, num_jobs)
	go worker(jobs, results)
	for i := 0; i < num_jobs; i++ {
		jobs <- i
	}
	close(jobs)

	for j := 0; j < num_jobs; j++ {
		fmt.Println(<-results)
	}
	close(results)

}

func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fibo(n)
	}

}

func fibo(n int) int {
	if n <= 1 {
		return n
	}

	return fibo(n-1) + fibo(n-2)

}
