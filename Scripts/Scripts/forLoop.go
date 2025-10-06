package main

import "fmt"

func forLoop() {
	int1 := 25
	for int1 <= 125 {
		fmt.Println(int1 * 1)
		int1 += 5
	}
}

func forLoopBreak() {
	for int1 := 1; int1 <= 20; int1++ {
		if int1 == 20 {
			break
		}
		fmt.Println(int1)
	}
}

func forLoopContinue() {
	for int1 := 1; int1 <= 18; int1++ {
		if int1 == 3 || int1 == 6 || int1 == 9 || int1 == 12 || int1 == 15 || int1 == 18 {
			continue
		}
		fmt.Println(int1)
	}
}

func main() {
	forLoop()
	forLoopBreak()
	forLoopContinue()

}
