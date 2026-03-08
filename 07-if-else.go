package main
import "fmt"

func MostrarIfElse(){
	fmt.Println("--- If / Else ---")

	if 7%2 == 0{
		fmt.Println("7 es Par")
	}else{
		fmt.Println("7 es Impar")
	}

	if 8%4 ==0{
		fmt.Println("8 es divisible por 4")
	}

	if 8%2 == 0 || 7%2 == 0{
		fmt.Println("8 es par o 7 es impar")
	}

	if 8%2 ==0 && 7<10{
		fmt.Println("8 es par y 7 es menor a 10")
	}

	if num:= 9 ; num <0{
		fmt.Println(num, "es negativo")
	}else if num < 10{
		fmt.Println(num, "tiene un digito")
	}else{
		fmt.Println(num, "tiene multiplos digitos")
	}

}