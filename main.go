package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const SLEEP = 400

var passLength, maxDigit, tries int
var repeatAllowed bool
var rnd *rand.Rand

func readInput(notification string) string {
	// reader := bufio.NewReader(os.Stdin)
	fmt.Print(notification)
	var text string
	// text, _ = reader.ReadString('\n')
	fmt.Scanln(&text)
	return strings.Split(text, "\n")[0]
}

func selectDifficulty() bool {
	passLength, _ = strconv.Atoi(readInput("Select password length: "))
	//fmt.Printf("%d length.\n", passLength)
	maxDigit, _ = strconv.Atoi(readInput("Select max digit for password (least is 1): "))
	//fmt.Printf("%d max.\n", maxDigit)
	if maxDigit < passLength {
		repeatAllowed = true
	} else {
		repeatAllowed = readInput("Are repeats allowed? (y/n): ") == "y"
	}
	//if repeatAllowed {
	//	fmt.Println("Repeats allowed.")
	//} else {
	//	//if maxDigit < passLength {
	//	//	fmt.Printf("ERROR: Can't generate password of %d length and digits " +
	//	//		"in 1-%d without repeats. Exiting... \n", passLength, maxDigit)
	//	//	return false
	//	//}
	//	fmt.Println("No repeats.")
	//}

	tries = passLength * maxDigit
	if !repeatAllowed {
		tries /= 2
	}
	return true
}

func doesArrHaveInt(arr []int, n int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == n {
			return true
		}
	}
	return false
}

func generateTask() string {
	task := make([]int, passLength)
	for i := 0; i < passLength; i++ {
		curr := rnd.Intn(maxDigit)+1
		if !repeatAllowed {
			for doesArrHaveInt(task, curr) {
				curr = rnd.Intn(maxDigit-1)+1
			}
		}
		task[i] = curr
	}
	taskStr := ""
	for i := 0; i < len(task); i++ {
		taskStr += strconv.Itoa(task[i])
	}
	return taskStr
}

// first is good-and-in-place, second is only matching but not in place
func checkTarget(target, input string) (int, int) {
	matches, misplaceds := 0, 0
	checkedTarget := make([]bool, len(target))
	checkedInput := make([]bool, len(input))
	for i := 0; i < len(target); i++ {
		if input[i] == target[i] {
			matches++
			checkedTarget[i] = true
			checkedInput[i] = true
		}
	}
	for i := 0; i < len(target); i++ {
		for j := 0; j < len(target); j++ {
			if !checkedTarget[j] && !checkedInput[i] && target[j] == input[i] {
				checkedTarget[j] = true
				checkedInput[i] = true
				misplaceds++
			}
		}
	}
	return matches, misplaceds
}

func main() {
	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
	if !selectDifficulty() {
		return
	}
	intro()
	target := generateTask()
	fmt.Printf("\nPassword of length %d, digits in range 1-%d, ", passLength, maxDigit)
	if repeatAllowed {
		fmt.Printf("digits can repeat.")
	} else {
		fmt.Printf("no repeating digits.")
	}
	// fmt.Print("Target is " + target)
	for currentTry := 0; currentTry < tries; currentTry++ {
		currInput := readInput(fmt.Sprintf("\nTry %d/%d > ", currentTry+1, tries))
		if currInput == target {
			fmt.Println("ACCESS GRANTED.")
			return
		}
		if len(currInput) != len(target) {
			fmt.Printf("WRONG LENGTH (%d) FOR PASSWORD OF LENGTH %d.", len(currInput), len(target))
			continue
		}
		if currentTry < tries-1 {
			match, misplaced := checkTarget(target, currInput)
			fmt.Printf("MATCH: %d, WRONG PLACE DIGITS: %d, %d tries remaining", match, misplaced, tries-currentTry-1)
		}
	}
	fmt.Println("YOU HAVE BEEN TRACED.")
	fmt.Println("CYBER SECURITY IS CONVERGING AT YOUR LOCATION.")
	fmt.Println("HAVE A NICE DAY.")
	fmt.Println("Password was " + target + ".")
}
