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

	return values
}

func playNote(num int) {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}

	if num == -1 || num == -2 {
		time.Sleep(70 * time.Millisecond)
		return
	}

	kb.SetKeys(keybd_event.VK_J)
	err = kb.Launching()
	if err != nil {
		panic(err)
	}

	//Augsta frekvence
	if num < 10 {
		if num == 0 || num == 1 || num == 2 || num == 3 {
			kb.SetKeys(keybd_event.VK_E, keybd_event.VK_Q, keybd_event.VK_TAB, keybd_event.VK_SPACE, keybd_event.VK_2, keybd_event.VK_4, keybd_event.VK_ENTER)
		} else if num == 4 || num == 5 || num == 6 {
			kb.SetKeys(keybd_event.VK_E, keybd_event.VK_Q, keybd_event.VK_TAB, keybd_event.VK_SPACE, keybd_event.VK_2, keybd_event.VK_3, keybd_event.VK_ENTER)
		} else if num == 7 || num == 8 || num == 9 {
			kb.SetKeys(keybd_event.VK_E, keybd_event.VK_Q, keybd_event.VK_TAB, keybd_event.VK_SPACE, keybd_event.VK_2, keybd_event.VK_2, keybd_event.VK_ENTER)
		}

	} else if num < 20 {
		if num == 10 || num == 11 || num == 12 || num == 13 {
			kb.SetKeys(keybd_event.VK_E, keybd_event.VK_Q, keybd_event.VK_TAB, keybd_event.VK_SPACE, keybd_event.VK_2, keybd_event.VK_1, keybd_event.VK_ENTER)
		} else if num == 14 || num == 15 || num == 16 {
			kb.SetKeys(keybd_event.VK_E, keybd_event.VK_Q, keybd_event.VK_TAB, keybd_event.VK_SPACE, keybd_event.VK_2, keybd_event.VK_0, keybd_event.VK_ENTER)
		} else if num == 17 || num == 18 || num == 19 {
			kb.SetKeys(keybd_event.VK_E, keybd_event.VK_Q, keybd_event.VK_TAB, keybd_event.VK_SPACE, keybd_event.VK_1, keybd_event.VK_9, keybd_event.VK_ENTER)
		}

	} else if num < 30 {
		if num == 20 || num == 21 || num == 22 || num == 23 {
			kb.SetKeys(keybd_event.VK_E, keybd_event.VK_Q, keybd_event.VK_TAB, keybd_event.VK_SPACE, keybd_event.VK_1, keybd_event.VK_8, keybd_event.VK_ENTER)
		} else if num == 24 || num == 25 || num == 26 {
			kb.SetKeys(keybd_event.VK_E, keybd_event.VK_Q, keybd_event.VK_TAB, keybd_event.VK_SPACE, keybd_event.VK_1, keybd_event.VK_7, keybd_event.VK_ENTER)
		} else if num == 27 || num == 28 || num == 29 {
			kb.SetKeys(keybd_event.VK_E, keybd_event.VK_Q, keybd_event.VK_TAB, keybd_event.VK_SPACE, keybd_event.VK_1, keybd_event.VK_6, keybd_event.VK_ENTER)
		}

	} else if num < 40 {
		if num == 30 || num == 31 || num == 32 || num == 33 {
			kb.SetKeys(keybd_event.VK_E, keybd_event.VK_Q, keybd_event.VK_TAB, keybd_event.VK_SPACE, keybd_event.VK_1, keybd_event.VK_5, keybd_event.VK_ENTER)
		} else if num == 34 || num == 35 || num == 36 {
			kb.SetKeys(keybd_event.VK_E, keybd_event.VK_Q, keybd_event.VK_TAB, keybd_event.VK_SPACE, keybd_event.VK_1, keybd_event.VK_4, keybd_event.VK_ENTER)
		} else if num == 37 || num == 38 || num == 39 {
			kb.SetKeys(keybd_event.VK_E, keybd_event.VK_Q, keybd_event.VK_TAB, keybd_event.VK_SPACE, keybd_event.VK_1, keybd_event.VK_3, keybd_event.VK_ENTER)
		}

	} else if num < 50 {
		if num == 40 || num == 41 || num == 42 || num == 43 {
			kb.SetKeys(keybd_event.VK_E, keybd_event.VK_Q, keybd_event.VK_TAB, keybd_event.VK_SPACE, keybd_event.VK_1, keybd_event.VK_2, keybd_event.VK_ENTER)
		} else if num == 44 || num == 45 || num == 46 {
			kb.SetKeys(keybd_event.VK_E, keybd_event.VK_Q, keybd_event.VK_TAB, keybd_event.VK_SPACE, keybd_event.VK_1, keybd_event.VK_1, keybd_event.VK_ENTER)
		} else if num == 47 || num == 48 || num == 49 {
			kb.SetKeys(keybd_event.VK_E, keybd_event.VK_Q, keybd_event.VK_TAB, keybd_event.VK_SPACE, keybd_event.VK_1, keybd_event.VK_0, keybd_event.VK_ENTER)
		}

	} else if num < 60 {
		if num == 50 || num == 51 || num == 52 || num == 53 {
			kb.SetKeys(keybd_event.VK_E, keybd_event.VK_Q, keybd_event.VK_TAB, keybd_event.VK_SPACE, keybd_event.VK_SP2, keybd_event.VK_5, keybd_event.VK_ENTER)
		} else if num == 54 || num == 55 || num == 56 {
			kb.SetKeys(keybd_event.VK_E, keybd_event.VK_Q, keybd_event.VK_TAB, keybd_event.VK_SPACE, keybd_event.VK_SP2, keybd_event.VK_6, keybd_event.VK_ENTER)
		} else if num == 57 || num == 58 || num == 59 {
			kb.SetKeys(keybd_event.VK_E, keybd_event.VK_Q, keybd_event.VK_TAB, keybd_event.VK_SPACE, keybd_event.VK_SP2, keybd_event.VK_4, keybd_event.VK_ENTER)
		}

	}

	time.Sleep(200 * time.Millisecond)
	err = kb.Launching()
	if err != nil {
		panic(err)
	}
}

func main() {
	values := arrays()
	//var play string
	var plays []int

	for i := 1; i < 80; i++ {
		if values[0][i] == 45 && values[1][i] == 45 && values[2][i] == 45 && values[3][i] == 45 && values[4][i] == 45 && values[5][i] == 45 {
			//play = play + `-`
			plays = append(plays, -1)
		} else { //ir  kāda nots
			switch {
			case values[0][i] != 45:
				plays = append(plays, int(values[0][i]-48))
			case values[1][i] != 45:
				plays = append(plays, int(values[1][i]-48)+10)
			case values[2][i] != 45:
				plays = append(plays, int(values[2][i]-48)+20)
			case values[3][i] != 45:
				plays = append(plays, int(values[3][i]-48)+30)
			case values[4][i] != 45:
				plays = append(plays, int(values[4][i]-48)+40)
			case values[5][i] != 45:
				plays = append(plays, int(values[5][i]-48)+50)
			default:
				plays = append(plays, -2)
				//play = play + `@`
			}
		}
	}

	time.Sleep(3 * time.Second)

	for _, number := range plays {
		fmt.Println(number)
		playNote(number)
	}

	fmt.Println(`PLAY: `, plays)
	time.Sleep(5 * time.Second)
	// for _, r := range play {
	// 	kb, err := keybd_event.NewKeyBonding()
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	c := string(r)
	// 	switch c {
	// 	case `-`:
	// 		time.Sleep(350 * time.Millisecond)
	// 	case `0`:
	// 		kb.SetKeys(keybd_event.VK_J)
	// 		err = kb.Launching()
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		kb.SetKeys(keybd_event.VK_E, keybd_event.VK_Q, keybd_event.VK_TAB, keybd_event.VK_SPACE, keybd_event.VK_1, keybd_event.VK_0, keybd_event.VK_ENTER)

	// 		time.Sleep(150 * time.Millisecond)
	// 		err = kb.Launching()
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		time.Sleep(50 * time.Millisecond)
	// 	case `1`:
	// 		kb.SetKeys(keybd_event.VK_J)
	// 		err = kb.Launching()
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		kb.SetKeys(keybd_event.VK_E, keybd_event.VK_Q, keybd_event.VK_TAB, keybd_event.VK_SPACE, keybd_event.VK_1, keybd_event.VK_0, keybd_event.VK_ENTER)

	// 		time.Sleep(150 * time.Millisecond)
	// 		err = kb.Launching()
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		time.Sleep(50 * time.Millisecond)
	// 	case `2`:
	// 		kb.SetKeys(keybd_event.VK_J)
	// 		err = kb.Launching()
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		kb.SetKeys(keybd_event.VK_E, keybd_event.VK_Q, keybd_event.VK_TAB, keybd_event.VK_SPACE, keybd_event.VK_1, keybd_event.VK_3, keybd_event.VK_ENTER)

	// 		time.Sleep(150 * time.Millisecond)
	// 		err = kb.Launching()
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		time.Sleep(50 * time.Millisecond)
	// 	case `3`:
	// 		kb.SetKeys(keybd_event.VK_J)
	// 		err = kb.Launching()
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		kb.SetKeys(keybd_event.VK_E, keybd_event.VK_Q, keybd_event.VK_TAB, keybd_event.VK_SPACE, keybd_event.VK_1, keybd_event.VK_6, keybd_event.VK_ENTER)

	// 		time.Sleep(150 * time.Millisecond)
	// 		err = kb.Launching()
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		time.Sleep(50 * time.Millisecond)
	// 	case `4`:
	// 		kb.SetKeys(keybd_event.VK_4)
	// 		time.Sleep(150 * time.Millisecond)
	// 		err = kb.Launching()
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 	case `5`:
	// 		kb.SetKeys(keybd_event.VK_5)
	// 		time.Sleep(150 * time.Millisecond)
	// 		err = kb.Launching()
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 	case `6`:
	// 		kb.SetKeys(keybd_event.VK_6)
	// 		time.Sleep(150 * time.Millisecond)
	// 		err = kb.Launching()
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 	default:
	// 		fmt.Println(`ERRORS`)
	// 		kb.SetKeys(keybd_event.VK_W)
	// 		time.Sleep(150 * time.Millisecond)
	// 		err = kb.Launching()
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 	}
	// }
}
