package main

import "fmt"
import "math"

func MostrarConstantes() {
	fmt.Println("---- 05 Constantes ----")
	const n = 500000000
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}
