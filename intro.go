package main

import (
	"time"
	"fmt"
)

func intro() {
	fmt.Println("")
	fmt.Println("======================================")
	fmt.Println("=== WELCOME TO TRI-OPTIMUM NETWORK ===")
	fmt.Println("======================================")
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
	time.Sleep(2*SLEEP * time.Millisecond)
	fmt.Println(" failed.")
	time.Sleep(3*SLEEP * time.Millisecond)
	fmt.Println("Access violation detected.")
	time.Sleep(2*SLEEP * time.Millisecond)
	fmt.Println("You are being traced.")
	time.Sleep(2*SLEEP * time.Millisecond)
	fmt.Printf("You have %d tries before they find you.\n", tries)
	time.Sleep(2*SLEEP * time.Millisecond)
}

func prettyPrint(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Print(string(s[i]))
		time.Sleep(time.Duration(rnd.Intn(100)+50) * time.Millisecond)
	}
}
