package main

import "fmt"

type player struct {
	plr1, plr2, plr3 string
}
type score struct {
	plr1, plr2, plr3 int
	ors              float64
}

func (p *player) bl(p1 string, p2 string, p3 string) {
	p.plr1 = p1
	p.plr2 = p2
	p.plr3 = p3
	fmt.Println("Batting lineup =", *p)
}
func (s *score) rs(r1 int, r2 int, r3 int, ov float64) {
	s.plr1 = r1
	s.plr2 = r2
	s.plr3 = r3
	s.ors = ov
	var ts int = (r1 + r2 + r3) //(s.plr1 + s.plr2 + s.plr3)
	fmt.Println("total score by INDIA =", ts)
	fmt.Println("in overs", ov)
	fmt.Println("runs scorred by players =", *s)
}
func main() {
	var lineup player
	lineup.bl("Rohit ", "Virat ", "Pant ")
	var runs score
	runs.rs(34, 78, 103, 48.4)

}
