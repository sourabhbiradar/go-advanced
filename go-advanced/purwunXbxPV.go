package main

import "fmt"

type channel struct {
	ch1, ch2, ch3, ch4 string
	prices             [4]int
}

func (c *channel) pass(m string, p string, w string, h string, prc [4]int) {
	c.ch1 = m
	c.ch2 = p
	c.ch3 = w
	c.ch4 = h
	c.prices = prc
	fmt.Println(*c, &c)

}

func (c *channel) tv() {
	if c.ch1 == "jj" {
		fmt.Println("cahnnel switched to ", c.ch1)
	} else {
		fmt.Println("channel not found")

	}

}

func main() {
	var watch channel
	watch.pass("mtv", "pogo", "wion", "hbo", [4]int{121, 32, 54, 23})
	watch.tv()
}
