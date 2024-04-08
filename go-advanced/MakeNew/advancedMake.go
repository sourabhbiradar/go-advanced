package main

import "fmt"

type year struct {
	month
	months []string
	count  []int
}
type month struct {
	days []string
}
type solarsys struct {
	num     int
	planets string
}

/*func (p *solarsys) sys(n int, pla string) {  //no need  of method
	p.num = n
	p.planets = pla
}*/

func (y *year) adMake(d []string, m []string, c []int) {
	y.days = d
	y.months = m
	y.count = c
	fmt.Println(y)

}
func main() {
	fmt.Println("Advanced make()")
	var t year
	t.adMake([]string{"sunday", "monday", "tuesday", "wednesday", "thrusday", "friday", "saturday"}, []string{"January", "February", "March", "April", "May", "June", "July"}, []int{1, 2, 3, 4, 5, 6, 7})
	var b []year = make([]year, 7)

	b[0] = year{month{[]string{"sunday"}}, []string{"january"}, []int{1}} //data type id needed for slice
	b[1] = year{month{[]string{"monday"}}, []string{"february"}, []int{2}}
	b[2] = year{month{[]string{"tuesday"}}, []string{"march"}, []int{3}}
	b[3] = year{month{[]string{"wednesday"}}, []string{"april"}, []int{4}}
	b[4] = year{month{[]string{"thrusday"}}, []string{"may"}, []int{5}}
	b[5] = year{month{[]string{"friday"}}, []string{"june"}, []int{6}}
	b[6] = year{month{[]string{"saturday"}}, []string{"july"}, []int{7}}

	fmt.Println(b[4])
	var s []solarsys = make([]solarsys, 8)
	s[0] = solarsys{1, "mercury"} // no need of data type
	s[1] = solarsys{2, "venus"}
	s[2] = solarsys{3, "earth"}
	s[3] = solarsys{4, "mars"}
	s[4] = solarsys{5, "jupiter"}
	s[5] = solarsys{6, "saturn"}
	s[6] = solarsys{7, "uranus"}
	s[7] = solarsys{8, "neptune"}
	fmt.Println("Planets in solar system :\n", s)

}
