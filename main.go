package main

import (
	"crypto/rand"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var modeG [2]string = [2]string{"-G", "--GENERATE"}
var modeR [2]string = [2]string{"-R", "--READ"}
var modeL [2]string = [2]string{"-L", "--Length"}

func main() {
	if len(os.Args) < 2 {
		getError("args < 2")
	}

	if strings.ToUpper(os.Args[1]) == modeG[0] || strings.ToUpper(os.Args[1]) == modeG[1] {
		checkArgs("-G", os.Args)

		max, err := strconv.Atoi(os.Args[3])
		checkError(err)
		createKey(os.Args[2], int32(max))

	}
}

func getError(err string) {
	fmt.Println("Error:", err)
	os.Exit(1)
}

func checkArgs(mode string, args []string) {
	if mode == modeG[0] || mode == modeG[1] {
		if len(args) != 4 {
			getError("mode = G, args != 4")
		}
	} else if mode == modeR[0] || mode == modeR[1] {
		if len(args) != 5 {
			getError("mode = R, args != 5")
		}
	} else if mode == modeL[0] || mode == modeL[1] {
		if len(args) != 3 {
			getError("mode = L, args != 3")
		}
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func createKey(filename string, max int32) {
	if max < 0 {
		getError("max number < 0")
	}
	saveKey(filename, generateKey(max))
}

func saveKey(filename string, key []byte) {
	file, err := os.Create(filename)
	checkError(err)
	defer file.Close()
	file.WriteString(string(key))
}

func generateKey(max int32) []byte {
	var slice []byte = make([]byte, max)
	_, err := rand.Read(slice)
	checkError(err)
	for max = max - 1; max >= 0; max-- {
		slice[max] = slice[max]%94 + 33
	}
	return slice
}
