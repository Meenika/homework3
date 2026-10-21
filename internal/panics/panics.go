package panics

import (
	"fmt"
	"strconv"
)

// 01. SafeRun запускает fn и сообщает, произошла ли panic.
// TODO: panic не должна выйти за пределы SafeRun.
func SafeRun(fn func()) (panicked bool) {
	defer func() {
		if recover() != nil {
			panicked = true
		}
	}()
	fn()
	return false
}

// 02. RecoverMessage возвращает текст значения, переданного в panic.
// TODO: если panic не было, верните пустую строку.
func RecoverMessage(fn func()) (message string) {
	defer func() {
		r := recover()
		if r != nil {
			message = fmt.Sprint(r)
		}
	}()
	fn()
	return ""
}

// 03. PanicIfEmpty вызывает panic("empty text") для пустой строки.
// TODO: непустой текст не должен приводить к panic.
func PanicIfEmpty(text string) {
	if text == "" {
		panic("empty text")
	}

}

// 04. MustPositive возвращает n, если n > 0.
// TODO: для нуля и отрицательных значений вызовите panic("not positive").
func MustPositive(n int) int {
	if n <= 0 {
		panic("not positive")
	}
	return n
}

// 05. ParseOrPanic преобразует строку в int.
// TODO: при ошибке преобразования вызовите panic с исходной ошибкой.
func ParseOrPanic(raw string) int {
	num, err := strconv.Atoi(raw)
	if err != nil {
		panic(err)
	}
	return num
}

// 06. TryParse вызывает ParseOrPanic и преобразует panic в (0, false).
// TODO: корректное число возвращается вместе с true.
func TryParse(raw string) (value int, ok bool) {
	defer func() {
		if recover() != nil {
			value = 0
			ok = false
		}
	}()
	return ParseOrPanic(raw), true
}

// 07. SafeIndex возвращает элемент по индексу и true.
// TODO: выход за границы не должен покидать функцию как panic;
// в таком случае верните zero value и false.
func SafeIndex(items []int, index int) (value int, ok bool) {
	defer func() {
		if recover() != nil {
			value = 0
			ok = false
		}
	}()
	return items[index], true
}

// 08. MustGet возвращает строку по индексу.
// TODO: не перехватывайте panic при выходе за границы.
func MustGet(items []string, index int) string {
	return items[index]
}

// 09. SafeMustGet вызывает MustGet и возвращает false вместо panic.
// TODO: корректный результат возвращается с true.
func SafeMustGet(items []string, index int) (value string, ok bool) {
	defer func() {
		if recover() != nil {
			value = ""
			ok = false
		}
	}()
	value = MustGet(items, index)
	return value, true
}

// 10. PanicToError преобразует panic в error с текстом "panic: <значение>".
// TODO: если fn завершилась нормально, верните nil.
func PanicToError(fn func()) (err error) {
	defer func() {
		r := recover()
		if r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
	}()
	fn()
	return nil
}

// 11. DeferBeforePanic показывает, что defer выполняется при panic.
// TODO: добавьте "body", затем вызовите panic и в deferred-функции
// добавьте "defer" и восстановите выполнение. Верните оба события.
func DeferBeforePanic() (events []string) {
	defer func() {
		events = append(events, "defer")
		recover()
	}()
	events = append(events, "body")
	panic("test panic")
}

// 12. RecoverOutsideDefer вызывает recover в обычном коде.
// TODO: верните true, только если recover вернул ненулевое значение.
func RecoverOutsideDefer() bool {
	return recover() != nil
}

// 13. RunSequence выполняет функции по порядку без recover.
// TODO: если panic нет, верните число выполненных функций.
// При panic выполнение должно остановиться естественным образом.
func RunSequence(functions []func()) int {
	completed := 0
	for _, fn := range functions {
		fn()
		completed++
	}
	return completed
}

// 14. RunSequenceSafe выполняет все функции, даже если отдельные функции паникуют.
// TODO: верните количество нормально завершившихся и количество panic.
func RunSequenceSafe(functions []func()) (completed, panicked int) {
	for _, fn := range functions {
		func() {
			didPanic := false
			defer func() {
				if recover() != nil {
					didPanic = true
				}
				if didPanic {
					panicked++
				}
			}()
			fn()
			completed++
		}()
	}
	return completed, panicked
}

// 15. MustNonNil возвращает значение по указателю.
// TODO: для nil вызовите panic("nil pointer").
func MustNonNil(value *int) int {
	if value == nil {
		panic("nil pointer")
	}
	return *value
}
