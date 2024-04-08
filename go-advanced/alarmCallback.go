package main

import (
	"fmt"
)

var globaltime int

func setalarm(time int) {

	globaltime = time
	
	fmt.Println("alarm set time :", time)
}
func ringalarm(time int) {
	fmt.Println("alarm ring", time)

}
func alarmapp(trigger func(int)) {
	fmt.Println("alarmapp")
	trigger(globaltime)
}

func main() {
	fmt.Println("callback function")
	setalarm(7)
	alarmapp(ringalarm)

}
