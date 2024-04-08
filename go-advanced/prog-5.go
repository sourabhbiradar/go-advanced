package main

import "fmt"

type fighting struct {
	son, daughter string
	beatings      int
}

func (p *fighting) inti(s string, d string, b int) {
	p.son = s
	p.daughter = d
	p.beatings = b
	fmt.Println(*p, &p)
}
func (p *fighting) peace() {
	if p.son == "Abhi" {
		fmt.Println("both got beatings")
	}

}
func main() {

	var ramu fighting
	ramu.inti("Abhi", "Rashu", 4)
	ramu.peace()

}
