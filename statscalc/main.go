// statscalc is a program for calculating basic statistical metrics
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
)

const (
	maxValue = 100000
	minValue = -100000
)

func readInput(r io.Reader) []int {
	var inputNum []int
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		txt := scanner.Text()

		if txt == "" {
			// Пустой ввод
			fmt.Println("Пустая строка. Введите число")
			continue
		}

		num, err := strconv.Atoi(txt)
		if err != nil {
			// Некорректный ввод
			fmt.Println("Ошибка ввода. Пожалуйста, введите число")
			continue
		}

		if num < minValue || num > maxValue {
			// Число вне допустимого диапазона
			fmt.Println("Число вне допустимого диапазона")
			continue
		}
		inputNum = append(inputNum, num)
	}

	return inputNum
}

// Функция для расчёта среднего
func calculateMean(inputNum []int) float64 {
	var sum float64
	for _, v := range inputNum {
		sum += float64(v)
	}
	return sum / float64(len(inputNum))
}

// Функция для расчёта медианы
func calculateMedian(inputNum []int) float64 {
	sort.Ints(inputNum)

	n := len(inputNum)

	if n%2 != 0 {
		return float64(inputNum[n/2])
	}

	num1 := inputNum[n/2-1]
	num2 := inputNum[n/2]
	return float64(num1+num2) / 2
}

// Функция для расчёта моды
func calculateMode(inputNum []int) int {
	counts := make(map[int]int)
	for _, num := range inputNum {
		counts[num]++
	}
	mode := inputNum[0]
	maxCount := counts[mode]
	for number, count := range counts {
		if count > maxCount || (count == maxCount && number < mode) {
			mode = number
			maxCount = count
		}
	}

	return mode
}

// Функция для расчёта стандартного отклонения
func calculateSD(inputNum []int) float64 {
	mean := calculateMean(inputNum)

	var sumPow float64
	for _, num := range inputNum {
		deviation := float64(num) - mean
		sumPow += math.Pow(deviation, 2)
	}

	sd := math.Sqrt(sumPow / float64(len(inputNum)))
	return math.Round(sd*100) / 100
}

func parseFlag() (bool, bool, bool, bool) {
	meanFlag := flag.Bool("mean", false, "Вывести среднее значение")
	medianFlag := flag.Bool("median", false, "Вывести медиану")
	modeFlag := flag.Bool("mode", false, "Вывести моду")
	sdFlag := flag.Bool("sd", false, "Вывести стандартное отклонение")

	flag.Parse()

	return *meanFlag, *medianFlag, *modeFlag, *sdFlag
}

func main() {
	// Чтение данных из стандартного ввода
	fmt.Println("Введите числа в диапазоне от -100000 до 100000")
	fmt.Println("Чтобы завершить ввод: Ctrl+D")

	inputNum := readInput(os.Stdin)

	if len(inputNum) == 0 {
		fmt.Println("Нет введенных чисел. Пожалуйста, введите хотя бы одно число.")
		return
	}

	// Подсчёт метрик
	mean := calculateMean(inputNum)
	median := calculateMedian(inputNum)
	mode := calculateMode(inputNum)
	sd := calculateSD(inputNum)

	// Парсинг флагов
	meanFlag, medianFlag, modeFlag, sdFlag := parseFlag()

	// Если не было флагов, показываем все
	if !meanFlag && !medianFlag && !modeFlag && !sdFlag {
		meanFlag = true
		medianFlag = true
		modeFlag = true
		sdFlag = true
	}

	fmt.Println("______________")

	// Вывод метрик
	if meanFlag {
		fmt.Printf("Mean: %.2f\n", mean)
	}

	if medianFlag {
		fmt.Printf("Median: %.2f\n", median)
	}

	if modeFlag {
		fmt.Printf("Mode: %.2f\n", float64(mode))
	}

	if sdFlag {
		fmt.Printf("SD: %.2f\n", sd)
	}

}
