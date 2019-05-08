package main

//TODO get length of first line and use it as measurement minus 1
//TODO

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/micmonay/keybd_event"
)

func arrays() [][]rune {
	values := [][]rune{}
	file, err := os.Open("song.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() { //reads single line
		str := scanner.Text()
		a := []rune(str)
		values = append(values, a)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// // Display first row.
	// fmt.Println("Row 1")
	// fmt.Println(values[0])

	// // Display second row.
	// fmt.Println("Row 2")
	// fmt.Println(values[1])

	// // Access an element.
	// fmt.Println("First element")
	// fmt.Println(values[0][0])

	// //Pirmā vērtība ir rinda, Otrā vērtība ir elements rindā
	// fmt.Println("Second element")
	// fmt.Println(values[0][1])

	return values
}

func main() {

	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}
	values := arrays()
	var play string

	for i := 1; i < 29; i++ {
		if values[0][i] == 45 && values[1][i] == 45 && values[2][i] == 45 && values[3][i] == 45 && values[4][i] == 45 && values[5][i] == 45 {
			play = play + `-`
		} else { //ir jāspēlē kāda nots
			switch {
			case values[0][i] != 45:
				play = play + string(values[0][i])
			case values[1][i] != 45:
				play = play + string(values[1][i])
			case values[2][i] != 45:
				play = play + string(values[2][i])
			case values[3][i] != 45:
				play = play + string(values[3][i])
			case values[4][i] != 45:
				play = play + string(values[4][i])
			case values[5][i] != 45:
				play = play + string(values[5][i])
			default:
				play = play + `@`
			}
		}
		fmt.Println(`Curr: `, play)
	}
	fmt.Println(`PLAY: `, play)

	time.Sleep(3 * time.Second)
	for _, r := range play {
		c := string(r)
		switch c {
		case `-`:
			kb.SetKeys(keybd_event.VK_J)
			err = kb.Launching()
			if err != nil {
				panic(err)
			}

			kb.SetKeys(keybd_event.VK_SP2)
			time.Sleep(150 * time.Millisecond)
			err = kb.Launching()
			if err != nil {
				panic(err)
			}
		case `0`:
			kb.SetKeys(keybd_event.VK_0)
			time.Sleep(150 * time.Millisecond)
			err = kb.Launching()
			if err != nil {
				panic(err)
			}
		case `1`:
			kb.SetKeys(keybd_event.VK_1)
			time.Sleep(150 * time.Millisecond)
			err = kb.Launching()
			if err != nil {
				panic(err)
			}
		case `2`:
			kb.SetKeys(keybd_event.VK_2)
			time.Sleep(150 * time.Millisecond)
			err = kb.Launching()
			if err != nil {
				panic(err)
			}
		case `3`:
			kb.SetKeys(keybd_event.VK_3)
			time.Sleep(150 * time.Millisecond)
			err = kb.Launching()
			if err != nil {
				panic(err)
			}
		case `4`:
			kb.SetKeys(keybd_event.VK_4)
			time.Sleep(150 * time.Millisecond)
			err = kb.Launching()
			if err != nil {
				panic(err)
			}
		case `5`:
			kb.SetKeys(keybd_event.VK_5)
			time.Sleep(150 * time.Millisecond)
			err = kb.Launching()
			if err != nil {
				panic(err)
			}
		case `6`:
			kb.SetKeys(keybd_event.VK_6)
			time.Sleep(150 * time.Millisecond)
			err = kb.Launching()
			if err != nil {
				panic(err)
			}
		default:
			kb.SetKeys(keybd_event.VK_W)
			time.Sleep(150 * time.Millisecond)
			err = kb.Launching()
			if err != nil {
				panic(err)
			}
		}
	}
}
