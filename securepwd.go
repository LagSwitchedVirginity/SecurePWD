package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/TwinProduction/go-color"
)

func ascii() {
	os.Clearenv()
	println(color.Blue +
		`
	 _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _
	 |   ____                           ______        ______   |
	 |  / ___|  ___  ___ _   _ _ __ ___|  _ \ \      / /  _ \  |
	 |  \___ \ / _ \/ __| | | | '__/ _ \ |_) \ \ /\ / /| | | | |
	 |   ___) |  __/ (__| |_| | | |  __/  __/ \ V  V / | |_| | |
	 |  |____/ \___|\___|\__,_|_|  \___|_|     \_/\_/  |____/  |
	 |                                                         |
	 |       A Tool designed to create secure Passwords.       |
	 |             [Coded by github.com/zentreax]              |
	 |_ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _|
														  ` + color.Reset)
}

func main() {
	ascii()
	fmt.Println()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("~> Password length (between 8-32) : ")
	scanner.Scan()
	length, _ := strconv.ParseInt(scanner.Text(), 10, 8)

	if length >= 8 && length <= 32 {
		rand.Seed(time.Now().UnixNano())
		digits := "0123456789"
		specials := "~!@#$%^&*()_+`-={}|[]\\:\"<>?,./"
		all := "ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
			"abcdefghijklmnopqrstuvwxyz" +
			digits + specials
		buf := make([]byte, length)
		buf[0] = digits[rand.Intn(len(digits))]
		buf[1] = specials[rand.Intn(len(specials))]
		for i := 2; i < int(length); i++ {
			buf[i] = all[rand.Intn(len(all))]
		}
		rand.Shuffle(len(buf), func(i, j int) {
			buf[i], buf[j] = buf[j], buf[i]
		})
		str := string(buf) // E.g. "3i[g0|)z"
		fmt.Println("Generated Password: " + str)
	} else {
		main()
	}
}
