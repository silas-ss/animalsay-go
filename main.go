package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/silas-ss/animalsay-go/api/animal"
	"github.com/silas-ss/animalsay-go/api/factories"
)

func main() {
	if len(os.Args) == 1 {
		printHelp()
		os.Exit(0)
	}

	animalType := strings.ToLower(os.Args[1])

	toSay := strings.Join(os.Args[2:], " ")

	cf := factories.CowFactory{}
	bf := factories.BirdFactory{}

	var animal animal.Animal

	if animalType == "cow" {
		animal = cf.CreateAnimal("Leitera")
	} else if animalType == "bird" {
		animal = bf.CreateAnimal("Voador")
	} else {
		log.Fatal("Animal is required")
	}

	animal.Say(toSay)

}

func printHelp() {
	fmt.Println("animalsay-go is program for cow say anything in your terminal!")
	fmt.Println("Type: animalsay-go <animal (cow or bird)> <to say>")
	fmt.Println("For example: animalsay-go cow hello world")
}
