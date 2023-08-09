package puppy

import (
	"fmt"

	"github.com/Shobhit130/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}

func From1() {
	fmt.Println("I am from version 1.1.0")
}
