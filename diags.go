package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"math/rand"
	//"time"
)

func main() {
	// Get terminal height and width
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	w, h := termbox.Size()
	defer termbox.Close()
	
	colors := []string{"\033[30m", "\033[31m", "\033[32m", "\033[33m", "\033[34m", "\033[35m", "\033[36m", "\033[37m"}
	endcolor := "\033[0m"
	chars := []string{"╱", "╲"}

	for y := 1; y <= h; y++ {
		for x := 1; x <= w; x++ {
			fmt.Printf("%s%s%s", choice(colors), choice(chars), endcolor)
		}
		//fmt.Printf("\n")
	}
	
	//termbox.SetInputMode(termbox.InputEsc)
	/*
	go func() {
		time.Sleep(30 * time.Second)
		termbox.Interrupt()
	}()
	*/

loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
			case termbox.EventKey:
				break loop

			//case termbox.EventInterrupt:
				//break loop
		}
	}
}

func choice(choices []string) string {
	return choices[rand.Intn(len(choices))]
}
