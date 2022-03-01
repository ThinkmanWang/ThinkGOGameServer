package main

import (
	"ThinkGOGameServer/thinkutils/logger"
	"fmt"
	"runtime"
)

var (
	log *logger.LocalLogger = logger.DefaultLogger()
)

type Animal struct {
	Name string
}

func (this *Animal) Eat() {
	fmt.Printf("%v is eating", this.Name)
	fmt.Println()
}

type Cat struct {
	*Animal
}

func (this *Cat) Eat() {
	//this.Animal.Eat()
	fmt.Println("FXXK")
}


func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())
    log.Info("Hello World")

	cat := &Cat{
		Animal: &Animal{
			Name: "cat",
		},
	}

	cat.Eat()
}
