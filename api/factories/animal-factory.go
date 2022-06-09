package factories

import "github.com/silas-ss/animalsay-go/api/animal"

type AnimalFactory interface {
	CreateAnimal(string)
}

type CowFactory struct{}

type BirdFactory struct{}

func (cf *CowFactory) CreateAnimal(name string) animal.Animal {
	cow := animal.Cow{Name: name}
	return cow
}

func (bf *BirdFactory) CreateAnimal(name string) animal.Animal {
	bird := animal.Bird{Name: name}
	return bird
}
