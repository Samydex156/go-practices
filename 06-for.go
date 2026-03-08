package main

import "fmt"

func MostrarFor() {
	fmt.Println("---- 06 For ----")
	i:=1
	/* Funciona como while: */
	for i <= 3{
		fmt.Println(i)
		i=i+1
	}

	for j:=0;j<3;j++{
		fmt.Println("range",j)
	}

	for i:= range 3{
		fmt.Println("range",i)
	}
	cont := 0

	for {
		fmt.Println("loop infinito", cont)
		if cont == 100{
			break
		}
		cont++
	}

	for n:=range 6{
		if n%2==0{
			continue
		}
		fmt.Println("continue",n)
	}


}
