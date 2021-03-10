package main

import (
	"fmt"
)

type Cutter interface {
	Cut()
}

type Cooker interface {
	Cook()
}

//The chef has the ability to cut and cook
type chef struct {
	Cutter
	Cooker
}

type knife struct {
}

type chainsaw struct {
}

type oven struct {
}

type pan struct {
}

//Knife becomes an implementer of Cutter
func (k knife) Cut() {
	fmt.Println("Chop Chop Chop...!")
}

//Chainsaw becomes an implementer of Cutter
func (c chainsaw) Cut() {
	fmt.Println("ZRRRRRRRRRRR!!!!")
}

//Pan becomes an implementer of Cooker
func (p pan) Cook() {
	fmt.Println("sweeeeesh!")
}

//Oven becomes an implementer of Cooker
func (o oven) Cook() {
	fmt.Println("BHHHHHHHMMMMMMMMM!")
}

//Chef does his work that is cut and cook irrespective of the implementers of the interfaces
func (c chef) Do() {
	c.Cut()
	c.Cook()
}

func main() {
	//Shannon Bennett a renowned Australian chef loves to use a knife and pan to work
	shannon := chef{
		Cutter: knife{},
		Cooker: pan{},
	}

	//Swish swissh and a swissh he makes an amazing dish!
	shannon.Do()

	//Fat Dude loves to chainsaw things and throw them into the oven to roast them.
	fatDude := chef{
		Cutter: chainsaw{},
		Cooker: oven{},
	}

	//ZRRRRRRRRRRR... setting to 220C and ... Nom nom nom!
	fatDude.Do()
}
