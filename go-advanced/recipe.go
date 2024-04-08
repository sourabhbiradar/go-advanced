package main

import "fmt"

type breakfast interface {
	cook()
}
type itemReqd struct {
	kitchenWare []string
	ings        map[string]float32
}

func (i *itemReqd) cook() {
	i.kitchenWare = []string{"pan,", "spatula,", "bowl,", "tablespoon"}
	i.ings = make(map[string]float32)
	i.ings["rava"] = 200
	i.ings["oil"] = 3
	i.ings["onion"] = 6
	i.ings["green chillies"] = 4
	i.ings["ginger"] = 1
	i.ings["coriander leaves"] = 1
	i.ings["curry leaves"] = 1
	i.ings["chana dal"] = 0.5
	i.ings["urad dal"] = 0.5

}

func main() {

	fmt.Println(" Upma Recipe\n", "\ningredients are in grams\n")
	var mother breakfast
	var chef itemReqd
	mother = &chef
	mother.cook()

	fmt.Println(chef.ings)
	fmt.Println("\nUtensils reqd are", chef.kitchenWare)

}
