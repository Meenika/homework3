package loops

type User struct {
	Name   string
	Active bool
}

// 01. SumTo считает сумму целых чисел от 1 до n включительно.
// TODO: для n <= 0 верните 0.
func SumTo(n int) int {
	count := 0
	if n <= 0 {
		return 0
	}
	for i := 1; i <= n; i++ {
		count = count + i
	}
	return count
}

// 02. SumBetween считает сумму от start до end включительно.
// TODO: если start > end, верните 0.
func SumBetween(start, end int) int {
	count := 0
	if start > end {
		return 0
	}
	for i := start; i <= end; i++ {
		count = count + i
	}
	return count
}

// 03. CountDown возвращает числа от n до 1.
// TODO: для n <= 0 верните пустой ненулевой слайс.
func CountDown(n int) []int {
	countSlice := make([]int, 0)
	if n <= 0 {
		return countSlice
	}
	for i := n; i >= 1; i-- {
		countSlice = append(countSlice, i)
	}
	return countSlice
}

// 04. Factorial вычисляет n! циклом.
// TODO: 0! и 1! равны 1; для отрицательного n верните 0.
func Factorial(n int) int {
	count := 1
	if n < 0 {
		return 0
	}
	if n == 1 || n == 0 {
		return 1
	}
	for i := 1; i <= n; i++ {
		count = count * i
	}
	return count
}

// 05. CountEven считает чётные числа в слайсе.
// TODO: ноль считается чётным.
func CountEven(items []int) int {
	count := 0
	for i := range items {
		if items[i]%2 == 0 {
			count++
		}
	}
	return count
}

// 06. FirstNegative возвращает первое отрицательное значение и true.
// TODO: если отрицательных чисел нет, верните zero value и false.
func FirstNegative(items []int) (int, bool) {
	for i := range items {
		if items[i] < 0 {
			return items[i], true
		}
	}
	return 0, false
}

// 07. SumWithoutZeros суммирует элементы, пропуская нули.
// TODO: используйте управляющую конструкцию, которая переходит
// к следующей итерации.
func SumWithoutZeros(items []int) int {
	count := 0
	for i := range items {
		if items[i] == 0 {
			continue
		}
		count = count + items[i]
	}
	return count
}

// 08. SumUntilLimit добавляет элементы по порядку, пока следующий элемент
// не сделал бы сумму больше limit.
// TODO: в этот момент остановите цикл. Для limit < 0 верните 0.
func SumUntilLimit(items []int, limit int) int {
	if limit < 0 {
		return 0
	}
	count := 0
	for i := range items {
		if count+items[i] > limit {
			break
		}
		count = items[i] + count
	}
	return count
}

// 09. DoubleInPlace умножает каждый элемент слайса на 2.
// TODO: измените исходный слайс, а не копию value из range.
func DoubleInPlace(items []int) {
	for i := range items {
		items[i] = items[i] * 2
	}
}

// 10. ReplaceNegativeInPlace заменяет отрицательные элементы на replacement.
// TODO: ноль и положительные значения должны сохраниться.
func ReplaceNegativeInPlace(items []int, replacement int) {
	for i := range items {
		if items[i] < 0 {
			items[i] = replacement
		}
	}
}

// 11. CountActive считает активных пользователей.
// TODO: порядок пользователей не влияет на количество.
func CountActive(users []User) int {
	count := 0
	for i := range users {
		if users[i].Active {
			count++
		}
	}
	return count
}

// 12. ActiveNames возвращает имена только активных пользователей.
// TODO: сохраните исходный порядок.
func ActiveNames(users []User) []string {
	activeNames := make([]string, 0)
	for user := range users {
		if users[user].Active == true {
			activeNames = append(activeNames, users[user].Name)
		}
	}
	return activeNames
}

// 13. RuneCount считает Unicode-символы через range по строке.
// TODO: результат для кириллицы и emoji не должен равняться числу байт.
func RuneCount(text string) int {
	count := 0
	for range text {
		count++
	}
	return count
}

// 14. RuneByteIndexes возвращает байтовые индексы всех rune в строке.
// TODO: используйте индекс, который отдаёт range по string.
func RuneByteIndexes(text string) []int {
	slice := make([]int, 0)
	for index := range text {
		slice = append(slice, index)
	}

	return slice
}

// 15. RepeatEach повторяет каждый элемент times раз подряд.
// TODO: для times <= 0 верните пустой ненулевой слайс.
func RepeatEach(items []int, times int) []int {
	slice := make([]int, 0)

	if times <= 0 {
		return slice
	}

	for _, item := range items {
		for i := 0; i < times; i++ {
			slice = append(slice, item)
		}
	}

	return slice
}
