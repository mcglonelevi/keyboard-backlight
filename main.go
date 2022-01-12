package main

import (
	"flag"
	"log"
    "os"
	"strconv"
)


func isFlagPassed(name string) bool {
    found := false
    flag.Visit(func(f *flag.Flag) {
        if f.Name == name {
            found = true
        }
    })
    return found
}

func check(e error) {
    if e != nil {
        log.Println("Make sure you have root permissions when running this script")
		panic(e)
    }
}

func main() {
	var colorFlag = flag.String("c", "", "color")
	var brightnessFlag = flag.Int("b", 255, "brightness")

	flag.Parse()

	if (*colorFlag != "") {
		color := []byte(*colorFlag)
		err := os.WriteFile("/sys/class/leds/system76_acpi::kbd_backlight/color", color, 0644)
		check(err)
	}
    
	brightness := []byte(strconv.Itoa(*brightnessFlag))
	brightnessErr := os.WriteFile("/sys/class/leds/system76_acpi::kbd_backlight/brightness", brightness, 0644)
    check(brightnessErr)
}