package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const SLEEP = 400

var passLength, maxDigit, tries int
var repeatAllowed bool

func readInput(notification string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(notification)
	text, _ := reader.ReadString('\n')
	return strings.Split(text, "\n")[0]
}

func selectDifficulty() bool {
	passLength, _ = strconv.Atoi(readInput("Select password length: "))
	fmt.Printf("\n%d length.\n", passLength)
	maxDigit, _ = strconv.Atoi(readInput("Select max digit for password (least is 1): "))
	fmt.Printf("\n%d max.\n", maxDigit)
	repeatAllowed = readInput("Are repeats allowed? (y/n): ") == "y"
	if repeatAllowed {
		fmt.Println("Repeats allowed.")
	} else {
		if maxDigit-1 < passLength {
			fmt.Printf("ERROR: Can't generate password of %d length and digits " +
				"in 1-%d without repeats. Exiting... \n", passLength, maxDigit)
			return false
		}
		fmt.Println("No repeats.")
	}
	tries = passLength * 5
	if !repeatAllowed {
		tries /= 2
	}
	return true
}

func intro() {
	fmt.Println("Connecting to database...")
	time.Sleep(SLEEP * time.Millisecond)
	fmt.Println("Retrieving encrypted hash list...")
	time.Sleep(SLEEP * time.Millisecond)
	fmt.Println("List found. Parsing...")
	time.Sleep(SLEEP * time.Millisecond)
	fmt.Print("Username: ")
	time.Sleep(SLEEP * time.Millisecond)
	prettyPrint("admin")
	fmt.Println()
	fmt.Print("Password: ")
	time.Sleep(SLEEP * time.Millisecond)
	prettyPrint("************")
	fmt.Println()
	fmt.Print("Bruteforcing the code")
	time.Sleep(SLEEP * time.Millisecond)
	fmt.Print(".")
	time.Sleep(SLEEP * time.Millisecond)
	fmt.Print(".")
	time.Sleep(SLEEP * time.Millisecond)
	fmt.Print(".")
	time.Sleep(SLEEP * time.Millisecond)
	fmt.Println(" failed.")
	time.Sleep(SLEEP * time.Millisecond)
	fmt.Println("Access violation detected.")
	time.Sleep(SLEEP * time.Millisecond)
	fmt.Println("You are being traced.")
	time.Sleep(SLEEP * time.Millisecond)
	fmt.Printf("You have %d tries before they find you.", tries)
	time.Sleep(SLEEP * time.Millisecond)
}

func prettyPrint(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Print(string(s[i]))
		time.Sleep(70 * time.Millisecond)
	}
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
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	task := make([]int, passLength)
	for i := 0; i < passLength; i++ {
		curr := rnd.Intn(maxDigit-1)+1
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
	if !selectDifficulty() {
		return
	}
	intro()
	target := generateTask()
	// fmt.Print("Target is " + target)
	for currentTry := 0; currentTry < tries; currentTry++ {
		currInput := readInput(fmt.Sprintf("\nTry %d/%d >", currentTry+1, tries))
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
			fmt.Printf("MATCH: %d, WRONG PLACE DIGITS: %d, %d tries remaining", match, misplaced, tries-currentTry)
		}
	}
	fmt.Println("YOU HAVE BEEN TRACED.")
	fmt.Println("CYBER SECURITY IS CONVERGING AT YOUR LOCATION.")
	fmt.Println("Password was " + target)
}
