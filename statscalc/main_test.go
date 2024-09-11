package main

import (
	"strings"
	"testing"
)

func TestCalculateMean(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}

	expected := 3.0

	result := calculateMean(input)

	if result != expected {
		t.Errorf("Ошибка в calculateMean: ожидается %.2f, получено %.2f", expected, result)
	}
}

func TestCalculateMedian(t *testing.T) {
	// Тест для нечётного количества элементов
	input := []int{1, 2, 3, 4, 5}
	expected := 3.0
	result := calculateMedian(input)

	if result != expected {
		t.Errorf("Ошибка в calculateMedian: ожидается %.2f, получено %.2f", expected, result)
	}

	// Тест для чётного количества элементов
	input = []int{1, 2, 3, 4, 5, 6}
	expected = 3.5
	result = calculateMedian(input)

	if result != expected {
		t.Errorf("Ошибка в calculateMedian для чётного числа: ожидается %.2f, получено %.2f", expected, result)
	}

	// Тест для одного элемента
	input = []int{1}
	expected = 1.0
	result = calculateMedian(input)

	if result != expected {
		t.Errorf("Ошибка в calculateMedian для одного элемента: ожидается %.2f, получено %.2f", expected, result)
	}
}

func TestCalculateMode(t *testing.T) {
	// Тест для набора, где мода однозначна
	input := []int{1, 2, 2, 3, 4}
	expected := 2
	result := calculateMode(input)

	if result != expected {
		t.Errorf("Ошибка в calculateMode: ожидается %d, получено %d", expected, result)
	}

	// Тест для нескольких одинаково частых элементов
	input = []int{1, 1, 2, 2, 3}
	expected = 1 // Мода должна быть наименьшей при одинаковой частоте
	result = calculateMode(input)

	if result != expected {
		t.Errorf("Ошибка в calculateMode: ожидается %d, получено %d", expected, result)
	}

	// Тест для одного элемента
	input = []int{7}
	expected = 7
	result = calculateMode(input)

	if result != expected {
		t.Errorf("Ошибка в calculateMode для одного элемента: ожидается %d, получено %d", expected, result)
	}
}

func TestCalculateSD(t *testing.T) {
	// Тест для набора данных
	input := []int{1, 2, 3, 4, 5}
	expected := 1.41 // Ожидаемое стандартное отклонение
	result := calculateSD(input)

	if result != expected {
		t.Errorf("Ошибка в calculateSD: ожидается %.2f, получено %.2f", expected, result)
	}

	// Тест для одинаковых элементов (SD должно быть 0)
	input = []int{5, 5, 5, 5}
	expected = 0.0
	result = calculateSD(input)

	if result != expected {
		t.Errorf("Ошибка в calculateSD для одинаковых чисел: ожидается %.2f, получено %.2f", expected, result)
	}

	// Тест для одного элемента (SD должно быть 0)
	input = []int{7}
	expected = 0.0
	result = calculateSD(input)

	if result != expected {
		t.Errorf("Ошибка в calculateSD для одного элемента: ожидается %.2f, получено %.2f", expected, result)
	}
}

func TestReadInput_ValidInput(t *testing.T) {
	// Эмулируем корректный ввод данных
	input := "1\n2\n3\n4\n5\n"
	r := strings.NewReader(input)

	// Вызываем функцию
	numbers := readInput(r)

	// Ожидаемые значения
	expected := []int{1, 2, 3, 4, 5}

	// Проверка результата
	for i, num := range numbers {
		if num != expected[i] {
			t.Errorf("Ожидалось %d, получено %d", expected[i], num)
		}
	}
}

func TestReadInput_InvalidInput(t *testing.T) {
	// Эмулируем некорректный ввод данных
	input := "1\nabc\n3\n"
	r := strings.NewReader(input)

	// Вызываем функцию
	numbers := readInput(r)

	// Ожидаемый результат: корректные числа
	expected := []int{1, 3}

	if len(numbers) != len(expected) {
		t.Errorf("Ожидалась длина %d, получено %d", len(expected), len(numbers))
	}

	for i, num := range numbers {
		if num != expected[i] {
			t.Errorf("Ожидалось %d, получено %d", expected[i], num)
		}
	}
}

func TestReadInput_EmptyInput(t *testing.T) {
	// Эмулируем пустой ввод данных
	input := ""
	r := strings.NewReader(input)

	// Вызываем функцию
	numbers := readInput(r)

	// Проверяем, что список чисел пуст
	if len(numbers) != 0 {
		t.Errorf("Ожидалось 0 чисел, получено %d", len(numbers))
	}
}
