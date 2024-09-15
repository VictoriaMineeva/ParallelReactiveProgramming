package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	fmt.Println("=== Реактивное программирование ===")

	// Создаем канал, который будет имитировать поток событий
	eventStream := make(chan int)

	// Запускаем горутину, которая будет генерировать события
	go generateEvents(eventStream)

	// Обрабатываем события в реактивном стиле
	processEvents(eventStream)
	asyncExample()
	syncExample()
}

func asyncExample() {
	fmt.Println("=== Асинхронное программирование ===")

	// Используем WaitGroup для синхронизации
	wg := sync.WaitGroup{}
	wg.Add(3)

	// Создаем канал для передачи данных
	ch := make(chan int)

	// Запускаем горутину-производителя
	go producer(ch, &wg)

	// Запускаем две горутины-потребителя
	go consumer(ch, &wg)
	go consumer(ch, &wg)

	// Ждем завершения всех горутин
	wg.Wait()
	fmt.Println("Асинхронные задачи завершены")
}

func producer(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		num := rand.Intn(100)
		fmt.Printf("Произведено число: %d\n", num)
		ch <- num
		time.Sleep(time.Millisecond * 500)
	}
	close(ch)
}

func consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range ch {
		fmt.Printf("Обработано число: %d\n", num)
		time.Sleep(time.Millisecond * 200)
	}
}

func syncExample() {
	fmt.Println("\n=== Синхронное программирование ===")

	nums := make([]int, 10)
	for i := 0; i < 10; i++ {
		nums[i] = rand.Intn(100)
		fmt.Printf("Произведено число: %d\n", nums[i])
	}

	fmt.Println("Обработка чисел:")
	for _, num := range nums {
		if num%2 == 0 {
			fmt.Printf("Обработано число: %d\n", num)
			time.Sleep(time.Millisecond * 200)
		}
	}

	fmt.Println("Синхронные задачи завершены")
}

func generateEvents(eventStream chan<- int) {
	defer close(eventStream)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 10; i++ {
		num := rand.Intn(100)
		fmt.Printf("Произведено число: %d\n", num)
		eventStream <- num
		time.Sleep(time.Millisecond * 500)
	}
}

func processEvents(eventStream <-chan int) {
	for num := range eventStream {
		if num%2 == 0 {
			fmt.Printf("Обработано число: %d\n", num)
			time.Sleep(time.Millisecond * 200)
		}
	}
	fmt.Println("Реактивные задачи завершены")
}
