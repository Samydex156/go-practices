package main

import (
	"time"

	"github.com/gen2brain/beeep"
) 

func main(){
	nombre :="Samuel Duran"
	hora:= time.Now()
	
	err := beeep.Notify(nombre,hora.GoString(),"")
	if err !=nil{
		panic(err)
	}
}