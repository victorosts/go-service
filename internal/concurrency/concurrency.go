package concurrency

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func FireAndForget() {
	fmt.Println("[FireAndForget] Provavelmente não vai imprimir o resultado da função")
	go func() {
		fmt.Println("Executando em algum momento. Pode ser até mesmo quando estiver rodando outras funções")
	}()
	fmt.Println("...")
}

func WgGroup() {
	fmt.Println("[WgGroup] Aqui já vai imprimir")

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		fmt.Println("Executando com wait group")
	}()
	fmt.Println("...")

	wg.Wait()
	fmt.Println("Impresso depois da espera")
}

func WgGroupLoop() {
	fmt.Println("[WgGroupLoop] ...")

	var wg sync.WaitGroup

	for i := range 10 {
		wg.Add(1)

		go func(n int) {
			defer wg.Done()
			fmt.Println(n)
		}(i)
	}
	fmt.Println("Impresso antes da espera")

	wg.Wait()
	fmt.Println("Impresso depois da espera")
}

func Channels() {
	result := make(chan string)

	go func() {
		result <- "ok"
	}()

	message := <-result
	fmt.Println(message)
}

func BuffChannels() {
	ch := make(chan int, 3)

	ch <- 1
	ch <- 2
	ch <- 3

	close(ch)

	for v := range ch {
		fmt.Printf("Channel %d\n", v)
	}
}

func BuffChannels2() {
	size := 10
	result := make(chan string, size)

	for i := range size {
		go func(i int) {
			result <- fmt.Sprintf("Channel %d", i)
		}(i)
	}

	for range size {
		fmt.Println(<-result)
	}
}

func BuffChannels3() {
	ch := make(chan int, 4)

	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4

	close(ch)

	for {
		v, ok := <-ch

		if !ok {
			break
		}

		fmt.Printf("Channel %d\n", v)
	}
}

func worker(jobs <-chan string, workerId int, outputs chan<- string) {
	for job := range jobs {
		fmt.Printf("Executando o job %s no worker %d\n", job, workerId)

		time.Sleep(2 * time.Second)

		result := fmt.Sprintf("Job %s finalizado no worker %d", job, workerId)

		if outputs != nil {
			outputs <- result
		} else {
			fmt.Println(result)
		}
	}
}

func FanOut() {
	jobs := make(chan string)

	for i := range 4 {
		// worker
		go worker(jobs, i, nil)
	}

	// Apenas gerando os jobs
	for j := range 12 {
		jobs <- fmt.Sprintf("job-%d", j)
	}
}

func FanIn() {
	var wg sync.WaitGroup

	jobs := make(chan string)
	outputs := make(chan string)

	for i := range 4 {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			worker(jobs, n, outputs)
		}(i)
	}

	go func() {
		for j := range 12 {
			time.Sleep(1 * time.Second)
			jobs <- fmt.Sprintf("job-%d", j)
		}

		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(outputs)
	}()

	for output := range outputs {
		fmt.Println(output)
	}
}

func ContextCancel() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				fmt.Println("Executando")
			}
		}
	}()

	time.Sleep(10 * time.Microsecond)

	cancel()
}

func SelectWithTimeout() {
	ctx, cancel := context.WithCancel(context.Background())

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		for {
			select {
			case msg := <-ch1:
				fmt.Printf("Mensagem do canal 1: %s\n", msg)
			case msg := <-ch2:
				fmt.Printf("Mensagem do canal 2: %s\n", msg)
			case <-time.After(1 * time.Second):
				cancel()
			}
		}
	}()

	go func() {
		for i := range 5 {
			ch1 <- fmt.Sprintf("channel 1 -> item %d", i)
		}
	}()

	go func() {
		for i := range 5 {
			ch2 <- fmt.Sprintf("channel 2 -> item %d", i)
		}
	}()

	<-ctx.Done()
	fmt.Println("Tarefa finalizada")
}
