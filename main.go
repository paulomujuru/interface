package main

import "fmt"

type schumuck interface {
	makeSituationAwkward() string
	partOfSociety() bool
	lieButBelieveItTrue() bool
}

type costanza struct {
	sex      string
	fullName string
	strong   bool
}

func (c costanza) makeSituationAwkward() string {
	return "I was in the pool!"
}

func (c costanza) partOfSociety() bool {
	return true
}

func (c costanza) lieButBelieveItTrue() bool {
	return false
}

type swanson struct {
	sex      string
	fullName string
	strong   string
}

func (s swanson) partOfSociety() bool {
	return true

}

func main() {
	var winston = costanza{}
	var george = costanza{
		sex:      "male",
		fullName: "George Costanza",
		strong:   true,
	}
	schumuck := []schumuck{george, winston}

	for _, s := range schumuck {
		fmt.Println(s.makeSituationAwkward())
	}
}
