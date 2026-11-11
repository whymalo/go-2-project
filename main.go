package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

func generateRandomElements(size int) []int {
	if size <= 0 {
		return nil
	}
	data := make([]int, size)
	for i := range data {
		data[i] = rand.Intn(1_000_000_000) + 1
	}
	return data
}

func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}
	maxValue := data[0]
	for _, value := range data[1:] {
		if value > maxValue {
			maxValue = value
		}
	}
	return maxValue
}

func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}
	maxima := make([]int, CHUNKS)
	chunkSize := len(data) / CHUNKS
	var wg sync.WaitGroup
	wg.Add(CHUNKS)
	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := start + chunkSize
		// Последняя часть включает остаток
		if i == CHUNKS-1 {
			end = len(data)
		}
		go func(index int, chunk []int) {
			defer wg.Done()
			maxima[index] = data[0]
			if len(chunk) > 0 {
				maxima[index] = maximum(chunk)
			}
		}(i, data[start:end])
	}
	wg.Wait()
	return maximum(maxima)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)
	fmt.Println("Ищем максимальное значение последовательно")
	started := time.Now()
	maxValue := maximum(data)
	elapsed := time.Since(started).Microseconds()
	fmt.Printf("Максимальное значение: %d\nВремя поиска: %d мкс\n", maxValue, elapsed)
	fmt.Printf("Ищем максимальное значение в %d горутинах\n", CHUNKS)
	started = time.Now()
	maxValue = maxChunks(data)
	elapsed = time.Since(started).Microseconds()
	fmt.Printf("Максимальное значение: %d\nВремя поиска: %d мкс\n", maxValue, elapsed)
}
