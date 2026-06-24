package concurrency

import (
	"fmt"
	"sync"
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
