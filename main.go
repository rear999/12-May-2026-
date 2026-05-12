package main

import (
	"fmt"
	"math/rand"
	"time"
)

// TODO 1: Функция добавления случайных чисел в массив
func addNum(count, rangeMax int, nums *[]int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < count; i++ {
		randomNum := rand.Intn(rangeMax) + 1 // числа от 1 до rangeMax
		*nums = append(*nums, randomNum)
	}
}

// TODO 2: Рекурсивная функция поиска дубликатов
func searchDubl(whatWeLookFor, index int, nums *[]int) int {
	if index >= len(*nums) {
		return 0
	}
	
	count := 0
	if (*nums)[index] == whatWeLookFor {
		count = 1
	}
	
	return count + searchDubl(whatWeLookFor, index+1, nums)
}

func main() {
	nums := []int{}
	
	//TODO 1.1:  100 диапазоне 1-1000
	fmt.Println("Заполнение массива случайными числами...")
	addNum(100, 1000, &nums)
	
	fmt.Printf("Длина массива: %d\n", len(nums))
	fmt.Printf("Первые 10 элементов: %v\n", nums[:10])
	
	//TODO 2.1: Ищем колво 15
	targetNumber := 15
	duplicatesCount := searchDubl(targetNumber, 0, &nums)
	
	//TODO: Результат
	fmt.Printf("\nЧисло %d встречается %d раз(а)\n", targetNumber, duplicatesCount)
	
	//Покажем все индексы где встречается число 15
	fmt.Printf("Индексы числа %d: ", targetNumber)
	for i, num := range nums {
		if num == targetNumber {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}