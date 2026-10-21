package arrays

// Блок закрепляет фиксированную длину массива, индексы, копирование,
// сравнение и работу с небольшими матрицами.

// 01. FirstOfThree возвращает первый элемент массива из трёх чисел.
// TODO: получите значение по правильному индексу.
func FirstOfThree(items [3]int) int {
	return items[0]
}

// 02. LastOfFour возвращает последний элемент массива из четырёх строк.
// TODO: решение не должно зависеть от конкретных значений массива.
func LastOfFour(items [4]string) string {
	return items[len(items)-1]
}

// 03. SetMiddle возвращает копию массива с заменённым средним элементом.
// TODO: исходный аргумент снаружи функции не должен измениться.
func SetMiddle(items [3]int, value int) [3]int {
	itemsCopy := items
	itemsCopy[1] = value
	return itemsCopy
}

// 04. SwapEdges меняет местами первый и последний элементы копии массива.
// TODO: два внутренних элемента должны сохранить свои позиции.
func SwapEdges(items [4]int) [4]int {
	firstItem := items[0]
	lastItem := items[len(items)-1]
	items[0] = lastItem
	items[len(items)-1] = firstItem
	return items
}

// 05. SumThree возвращает сумму трёх элементов.
// TODO: учтите все позиции массива.
func SumThree(items [3]int) int {
	count := 0
	for item := range items {
		count = count + items[item]
	}
	return count
}

// 06. AverageFour возвращает целочисленное среднее четырёх чисел.
// TODO: используйте обычные правила целочисленного деления Go.
func AverageFour(items [4]int) int {
	count := 0
	for item := range items {
		count = count + items[item]
	}
	return count / len(items)
}

// 07. ReverseFive возвращает массив в обратном порядке.
// TODO: первый элемент должен стать последним, второй — предпоследним и т.д.
func ReverseFive(items [5]int) [5]int {
	for i := 0; i < len(items)/2; i++ {
		opposite := len(items) - 1 - i
		temp := items[i]
		items[i] = items[opposite]
		items[opposite] = temp
	}
	return items
}

// 08. EqualPairs сравнивает два массива строк целиком.
// TODO: порядок элементов имеет значение.
func EqualPairs(a, b [2]string) bool {
	return a == b
}

// 09. CopyAndSet возвращает исходную копию и изменённую копию массива.
// TODO: измените только элемент с индексом index во втором результате.
// В тестах index всегда находится в допустимом диапазоне.
func CopyAndSet(items [3]int, index, value int) ([3]int, [3]int) {
	itemsCopy := items
	itemsCopy[index] = value
	return items, itemsCopy
}

// 10. CountTrue считает true в массиве из пяти bool.
// TODO: false не увеличивает счётчик.
func CountTrue(items [5]bool) int {
	count := 0
	for item := range items {
		if items[item] == true {
			count++
		}
	}
	return count
}

// 11. ContainsFour проверяет наличие target в массиве.
// TODO: верните true при первом совпадении, иначе false.
func ContainsFour(items [4]int, target int) bool {
	for item := range items {
		if items[item] == target {
			return true
		}
	}
	return false
}

// 12. MaxFour возвращает максимальное значение массива.
// TODO: корректно обработайте массив, состоящий только из отрицательных чисел.
func MaxFour(items [4]int) int {
	maxNum := items[0]
	for i := 1; i < len(items); i++ {
		if maxNum < items[i] {
			maxNum = items[i]
		}
	}
	return maxNum
}

// 13. MainDiagonalSum возвращает сумму главной диагонали матрицы 3x3.
// TODO: используйте элементы [0][0], [1][1] и [2][2].
func MainDiagonalSum(matrix [3][3]int) int {
	count := 0

	for i := 0; i < len(matrix); i++ {
		count += matrix[i][i]
	}
	return count
}

// 14. CornersSum возвращает сумму четырёх углов матрицы 3x3.
// TODO: центральный элемент и середины сторон не учитываются.
func CornersSum(matrix [3][3]int) int {
	return matrix[0][0] + matrix[0][2] + matrix[2][0] + matrix[2][2] // схитрила чутка
}

// 15. RotateRight сдвигает массив вправо на одну позицию.
// TODO: последний элемент должен перейти в начало.
func RotateRight(items [4]int) [4]int {
	last := items[3]
	for i := len(items) - 1; i > 0; i-- {
		items[i] = items[i-1]
	}
	items[0] = last
	return items
}
