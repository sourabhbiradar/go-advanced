package main

import "fmt"

type cards struct {
	card1, card2, card3 string
}

func (p *cards) inti(c1 string, c2 string, c3 string) {
	p.card1 = c1
	p.card2 = c2
	p.card3 = c3
	fmt.Println(*p, &p)
}
func (p *cards) game() {

	var hand cards = *p
	if hand == *p {

		fmt.Println("your odds of winning game are good")
	}

}
func main() {
	var deck cards
	deck.inti("ace of hearts,", "two of clubs,", "three of diamonds")
	deck.game()

}
