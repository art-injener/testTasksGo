package main

import (
	"fmt"
	"strings"
)

// Проверка строки на сбалансированность
func checkBrackets(s string) bool {
	// создадим слайс где будет размещаться стек
	stack := make([]int, 0)
	// добавление в стек
	push := func(v int) {
		stack = append(stack, v)
	}
	// изъятие из стека
	pop := func() int {
		if len(stack) > 0 {
			v := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			return v
		}
		return -1
	}
	// группа открывающихся скобочек
	bracketsIn := "{[("
	// группа закрывающихся скобочек
	bracketsOut := "}])"
	// основной цикл, один проход по строке, проходим по всем буквам
	for _, c := range s {
		inIdx := strings.Index(bracketsIn, string(c))
		if inIdx >= 0 { // если очередная буква это открывающаяся скобочка - добавляем её в стек
			push(inIdx)
		} else {
			outIdx := strings.Index(bracketsOut, string(c))
			// если очередная буква это закрывающаяся скобочка - вычитываем стек и проверяем соответстует ли ей предыдущая скобочка
			if outIdx >= 0 {
				inIdx = pop()
				if inIdx < 0 || inIdx != outIdx {
					return false // стек пуст или скобочка не соответствует, не сбалансированная строка
				}
			}
		}
	}
	if len(stack) != 0 {
		return false // что-то осталось в стеке - точно не сбалансированная строка
	}
	return true // всё ок, строка сбалансированная
}

func main() {
	s := "((b)"
	fmt.Println(s, checkBrackets(s))
}
