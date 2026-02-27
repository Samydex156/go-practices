package main
import "time"
import "fmt"

func MostrarTiempo(){
	fmt.Println("---000-tiempo.go---")
	fmt.Println(time.Now())
	ahora:=time.Now()
	fecha:=ahora.Format("02/01/2006")
	fmt.Println(fecha)
	hora24:=ahora.Format("15:04:05")
	fmt.Println(hora24)
	hora12:= ahora.Format("03:04 PM")
	fmt.Println(hora12)
	fechaT:=ahora.Format("Monday, January 02, 2006")
	fmt.Println(fechaT)
	fechaT2:=ahora.Format("Viernes, Enero 02, 2006")
	fmt.Println(fechaT2)

	hoy:=ahora.Format("Hoy es Monday 02 de January de 2006")
	fmt.Println(hoy)
	
}