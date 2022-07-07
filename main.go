package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	animalNumber := 10
	animals := farmGenerator(animalNumber)
	farmInfo(animals)
	fmt.Printf("\nCount of feed u need for this animals - %d", allFeedCount(animals))
}

const dogFeed int = 2
const catFeed int = 7
const cowFeed int = 25
const catCoefficient float64 = 0.4
const cowCoefficient float64 = 2.45

type Animal struct {
	name      string
	weight    int
	feedPerKg int
}

type Dog struct {
	Animal
}
type Cat struct {
	Animal
}
type Cow struct {
	Animal
}

type createAnimal interface {
	feedWeight() int
	infoAnimal() string
}

func (animal Animal) feedWeight() int {
	return animal.weight * animal.feedPerKg
}

func (animal Animal) infoAnimal() string {
	return fmt.Sprintf("Name - %s\tWeight - %dkg\t Feed - %d\n", animal.name, animal.weight, animal.feedWeight())
}

func farmGenerator(animalCount int) (animals []createAnimal) {
	var animal createAnimal

	for i := 0; i < animalCount; i++ {
		animalNumber := rand.Intn(3) + 1
		animalWeightTemplate := rand.Intn(20) + 40

		switch animalNumber {
		case 1:
			animal = Dog{
				Animal{
					name:      "Dog",
					weight:    animalWeightTemplate,
					feedPerKg: dogFeed,
				},
			}
		case 2:
			animal = Cat{
				Animal{
					name:      "Cat",
					weight:    int(float64(animalWeightTemplate) * catCoefficient),
					feedPerKg: catFeed,
				},
			}
		case 3:
			animal = Cow{
				Animal{
					name:      "Cow",
					weight:    int(float64(animalWeightTemplate) * cowCoefficient),
					feedPerKg: cowFeed,
				},
			}
		}
		animals = append(animals, animal)
	}

	return animals
}

func farmInfo(animals []createAnimal) {
	fmt.Println("\nSo here is animals list:")
	for i, a := range animals {
		fmt.Printf("#%d\t%s", i+1, a.infoAnimal())
	}
}

func allFeedCount(animals []createAnimal) (count int) {
	for _, a := range animals {
		count += a.feedWeight()
	}
	return count
}
