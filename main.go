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
	} else if strings.ToUpper(os.Args[1]) == modeR[0] || strings.ToUpper(os.Args[1]) == modeR[1] {
		checkArgs("-R", os.Args)

		pos, err := strconv.Atoi(os.Args[3])
		checkError(err)
		quan, err := strconv.Atoi(os.Args[4])
		checkError(err)

		readKey(os.Args[2], int32(pos), int32(quan), getLength(os.Args[2]))
	} else if strings.ToUpper(os.Args[1]) == modeL[0] || strings.ToUpper(os.Args[1]) == modeL[1] {
		checkArgs("-L", os.Args)
		fmt.Println("Length of file:", getLength(os.Args[2]))
	} else {
		getError("mode not found")
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

func getLength(filename string) int32 {
	file, err := os.Open(filename)
	checkError(err)
	defer file.Close()

	info, err := file.Stat()
	checkError(err)
	return int32(info.Size())
}

func readKey(filename string, position, quantity, length int32) {
	if length < position+quantity {
		getError("length of file < position + quantity")
	} else {
		fmt.Println(getChars(filename, length, position, quantity))
	}
}

func getChars(filename string, length, position, quantity int32) string {
	file, err := os.Open(filename)
	checkError(err)
	defer file.Close()

	var bytes []byte = make([]byte, length)
	_, err = file.Read(bytes)
	checkError(err)

	var read string
	var max int32 = position + quantity

	for i := position; i < max; i++ {
		read += string(bytes[i])
	}

	return read
}
