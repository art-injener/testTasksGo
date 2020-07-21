package main

import (
	"fmt"
	"sort"
)

const (
	A_ARRAY = iota
	B_ARRAY
)

type UnionArr struct {
	cntA uint32
	cntB uint32
}

//посик пересечения множеств
func ArrayProcessor(A []int, B []int) ([]int, []int, []int) {

	set := make(map[int]*UnionArr)

	//определение длины для прохода цикла
	N := len(A)
	lenghtA, lenghtB := N, len(B)

	if lenghtB > lenghtA {
		N = lenghtB
	}
	//добавление элементов в множество
	for i := 0; i < N; i++ {

		if i < lenghtA {
			addValueToSet(set, A[i], A_ARRAY)
		}

		if i < lenghtB {
			addValueToSet(set, B[i], B_ARRAY)
		}
	}
	//преобразование множества в слайсы
	return setToSlice(set)
}

//добавление элемента в множество
func addValueToSet(set map[int]*UnionArr, value int, index int) {
	//проверка на наличие элемента
	if _, ok := set[value]; !ok {
		set[value] = &UnionArr{0, 0}
	}
	//увеличение счётчика вхождения в зависимости от источника данных
	switch index {
	case A_ARRAY:
		set[value].cntA++
	case B_ARRAY:
		set[value].cntB++
	}
}

//преобразование множества в слайс
func setToSlice(set map[int]*UnionArr) ([]int, []int, []int) {

	var AA []int
	var AB []int
	var BB []int

	for key, value := range set {
		//если оба счётчика > 0, то элемент входит в оба массива
		if value.cntA > 0 && value.cntB > 0 {
			AB = append(AB, key)
		} else if value.cntA > 0 {
			//счётчик А > 0, то элемент входит в массив А
			AA = append(AA, key)
		} else {
			// в остальных случаях элемент принадлежит массиву В
			BB = append(BB, key)
		}
	}

	//сортировка для упорядоченного вывода
	sort.Ints(AA)
	sort.Ints(AB)
	sort.Ints(BB)

	return AA, AB, BB
}

// функция для вывод данных
func arrayPrinter(AA []int, AB []int, BB []int) {

	fmt.Println("After process :")
	fmt.Printf("only array A: %v \n", AA)
	fmt.Printf("only array B: %v \n", BB)
	fmt.Printf("union arrays A and B: %v \n", AB)
}

func main() {
	A := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	B := []int{6, 7, 8, 9, 10, 11, 12, 13, 14}

	fmt.Printf("A: %v \n", A)
	fmt.Printf("B: %v \n", B)

	arrayPrinter(ArrayProcessor(A, B))
}
