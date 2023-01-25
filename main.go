package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

// Структура для показателей самого большого участка
type target struct {
	count     int
	color     int
	x, y      int
	addresses []*int
}

func main() {

	// Ширина, высота матрицы и количество цветов
	var width, height, countColors int

	// Входные данные
	fmt.Println("Введите ширину и высоту матрицы через пробел")
	fmt.Scanf("%d %d", &width, &height)
	if width == 0 || height == 0 {
		fmt.Println("Ширина или высота не может быть равна нулю")
		os.Exit(0)
	}
	fmt.Println("Введите количество цветов")
	fmt.Scanf("%d", &countColors)
	if countColors == 0 {
		fmt.Println("Количество цветом не может быть равно нулю")
		os.Exit(0)
	}

	// Создаем матрицу
	matrix := createMatrix(height, width, countColors)

	// Выводим матрицу в консоль
	printMatrix(matrix)

	// Проводим основную логику
	target := find(matrix)

	// Проходимся по каждому указателю и меняем его цвет
	for _, address := range target.addresses {
		*address = target.color
	}

	// Снова принтуем матрицу
	printMatrix(matrix)
	fmt.Printf("Координаты верхней левой точки участка (с левого верхнего угла): %d, %d\nЦвет: %d\nКоличество элементов: %d", target.x+1, target.y+1, target.color, target.count)
}

// Функция создания матрицы
func createMatrix(height, width, countColors int) [][]int {

	// Конфигурируем сид
	rand.Seed(time.Now().UnixNano())

	// Создаем массив массивов длиной, равной высоте матрицы
	matrix := make([][]int, height)

	for i := 0; i < height; i++ {

		// Создаем строку, наполняем рандомными данными
		line := make([]int, width)
		for j := 0; j < width; j++ {
			line[j] = rand.Intn(countColors) + 1
		}

		// Заполняем матрицу
		matrix[i] = line
	}

	return matrix
}

func find(matrix [][]int) target {

	target := target{}

	// Создаем массив указателей на число объемом, равным площади матрицы
	addresses := make([]*int, 0, len(matrix)*len(matrix[0]))

	// Проходимся по каждой клетке матрицы построчно
	for y, line := range matrix {
		for x := 0; x < len(line); x++ {

			// Если цвет очередной клетки ненулевой
			if matrix[y][x] != 0 {

				// Заранее запоминаем цвет
				color := matrix[y][x]

				// То сканируем область
				count, addresses := findAndReplace(matrix, x, y, color, 1, addresses)

				// Сравниваем количество элементов с последним наибольшим
				if target.count < count {

					// Копируем адреса элементов для дальнейшего закрашивания
					target.addresses = make([]*int, count)
					copy(target.addresses, addresses)

					// Передаем остальные параметры
					target.color = color
					target.x, target.y = x, y
					target.count = count
				}
			}
		}
	}

	return target
}

// Функция поиска наибольшего участка
func findAndReplace(matrix [][]int, x, y, targetColor, count int, addresses []*int) (int, []*int) {

	// Получаем ограничения матрицы
	maxY, maxX := len(matrix)-1, len(matrix[0])-1

	// Обнуляем цвет текущего элемента
	matrix[y][x] = 0

	// Добавляем адрес текущего элемента в массив
	addresses = append(addresses, &matrix[y][x])

	/* Смотрим на ограничения массивов, чтобы не выйти из диапазона
	и смотрим клетки, находящиеся вокруг - если клетка интересующего нас цвета,
	то рекурсивно вызываем эту же функцию, пока не обнулим все соседние клетки */
	if x < maxX && matrix[y][x+1] == targetColor {
		count, addresses = findAndReplace(matrix, x+1, y, targetColor, count+1, addresses)
	}
	if y < maxY && matrix[y+1][x] == targetColor {
		count, addresses = findAndReplace(matrix, x, y+1, targetColor, count+1, addresses)
	}
	if y > 0 && matrix[y-1][x] == targetColor {
		count, addresses = findAndReplace(matrix, x, y-1, targetColor, count+1, addresses)
	}
	if x > 0 && matrix[y][x-1] == targetColor {
		count, addresses = findAndReplace(matrix, x-1, y, targetColor, count+1, addresses)
	}
	return count, addresses
}

func printMatrix(matrix [][]int) {
	for _, line := range matrix {
		fmt.Println(line)
	}
	fmt.Println()
}
