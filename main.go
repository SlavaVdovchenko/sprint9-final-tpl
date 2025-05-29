package main

import (
	"fmt"
	"math"
	"math/rand/v2"
	"slices"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func GenerateRandomElements(size int) []int {
	// ваш код здесь
	slice := rand.Perm(size)
	return slice
}

// maximum returns the maximum number of elements.
func Maximum(data []int) int {
	// ваш код здесь
	if len(data) == 0 {
		return math.MinInt
	}
	return slices.Max(data)
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) == 0 {
		panic("empty slice")
	}

	var wg sync.WaitGroup
	chunkSize := len(data) / CHUNKS
	maxValues := make([]int, CHUNKS)

	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > len(data) {
			end = len(data)
		}

		if start >= len(data) {
			continue
		}

		wg.Add(1)
		go func(index int, chunk []int) {
			defer wg.Done()
			max := chunk[0]
			for _, val := range chunk {
				if val > max {
					max = val
				}
			}
			maxValues[index] = max
		}(i, data[start:end])
	}

	wg.Wait()

	globalMax := maxValues[0]
	for _, val := range maxValues {
		if val > globalMax {
			globalMax = val
		}
	}

	return globalMax
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	// ваш код здесь
	slice := GenerateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь
	start := time.Now()
	max := Maximum(slice)
	elapsed := time.Since(start)
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	// ваш код здесь
	start = time.Now()
	max = maxChunks(slice)
	elapsed = time.Since(start)
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

}
