package main

import (
	"fmt"
	"time"
)

func MostrarSwitch() {
	fmt.Println("--- 08-switch ---")
	fmt.Println(time.Now())

	i := 2
	fmt.Print("Escribe ", i, " como ")
	switch i {
	case 1:
		fmt.Println("Uno")
	case 2:
		fmt.Println("Dos")
	case 3:
		fmt.Println("Tres")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Es fin de semana")
	default:
		fmt.Println("Es un dia de la semana")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println(" Es antes del mediodia")
	default:
		fmt.Println("Es despues del mediodia")
	}

	QueSoy := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("Soy booleano")
		case int: fmt.Println("Soy numero Entero")
		default: fmt.Printf("Aún no lo sé pero por ahora tipo %T\n",t)
		}

	}
	QueSoy(true)
	QueSoy(123)
	QueSoy("samydex156")
}
