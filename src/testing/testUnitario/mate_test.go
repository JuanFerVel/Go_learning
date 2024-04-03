package testunitario

import (
	"testing"
)

// func TestSuma(t *testing.T) {
// 	total := Suma(5, 5)
// 	if total != 10 {
// 		t.Errorf("Suma Incorrecta, tiene %d se esperaba %d", total, 10)
// 		return
// 	}

// 	fmt.Println(total)
// }

func TestSuma(t *testing.T) {
	tabla := []struct {
		a int
		b int
		c int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{25, 25, 50},
	}

	for _, item := range tabla {
		total := Suma(item.a, item.b)

		if total != item.c {
			t.Errorf("Suma Incorrecta, tiene %d se esperaba %d", total, item.c)
			return
		}
	}
}

func TestGetMax(t *testing.T) {
	tabla := []struct {
		a int
		b int
		c int
	}{
		{1, 2, 2},
		{7, 2, 7},
	}

	for _, item := range tabla {
		max := GetMax(item.a, item.b)

		if max != item.c {
			t.Errorf("GetMax Incorrecta, tiene %d se esperaba %d", max, item.c)
			return
		}
	}
}

func TestFibbonacci(t *testing.T) {
	tabla := []struct {
		n int
		r int
	}{
		{1, 1},
		{8, 21},
		{50, 12586269025},
	}

	for _, item := range tabla {
		fibo := Fibonacci(item.n)

		if fibo != item.r {
			t.Errorf("Fibonacci Incorrecta, tiene %d se esperaba %d", fibo, item.r)
			return
		}
	}
}
