package puppy

import (
	"github.com/victorasongit/dog"
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