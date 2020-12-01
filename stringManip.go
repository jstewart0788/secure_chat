package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type cypherText [][]string

func removeCarriageWindows(s string) string {
	re := regexp.MustCompile(`\r?\n`)
	newString := re.ReplaceAllString(s, "")
	return newString
}

func stringToBinaryArray(s string) []string {
	binaryString := ""
	for _, c := range s {
		binaryString = fmt.Sprintf("%s%.8b", binaryString, c)
	}
	return strings.Split(binaryString, "")
}

func binaryArrayToByteSlice(s []string) []byte {

	var newArray []byte
	var tempString string = ""
	for i, bit := range s {
		tempString = tempString + bit
		if i != 0 && len(tempString)%8 == 0 {
			newByte, _ := strconv.ParseUint(tempString, 2, 8)
			newArray = append(newArray, byte(newByte))
			tempString = ""
		}
	}
	return newArray
}

func stringToEightByteArray(s string) []string {

	if len(s) == 0 {
		fmt.Println("No message to encyrpt!")
		os.Exit(1)
	}

	if len(s)%8 != 0 {
		for ok := true; ok; ok = len(s)%8 != 0 {
			s = s + " "
		}
	}

	var eightByteArray []string
	var tempString string = ""
	for i, char := range s {
		tempString = tempString + string(char)
		if i != 0 && len(tempString)%8 == 0 {
			eightByteArray = append(eightByteArray, tempString)
			tempString = ""
		}
	}

	return eightByteArray
}

func stringToSixBitArray(s []string) []string {

	if len(s) != 48 {
		fmt.Println("Invalid length string passed into sbox!")
		os.Exit(1)
	}
	var eightBitArray []string
	var tempString string = ""
	for i, bit := range s {
		tempString = tempString + string(bit)
		if i != 0 && len(tempString)%6 == 0 {
			eightBitArray = append(eightBitArray, tempString)
			tempString = ""
		}
	}

	return eightBitArray
}

func (c cypherText) print() {
	var byteArray []byte
	for _, bits := range c {
		byteArray = append(byteArray, binaryArrayToByteSlice(bits)...)
	}

	fmt.Println("Cyphertext:", string(byteArray))
}

func formatIV(s string) string {

	IVLength := len(s)
	if IVLength <= 8 {
		zerosNeeded := 8 - IVLength
		for i := 0; i < zerosNeeded; i++ {
			s = strconv.Itoa(0) + s
		}
	} else {
		s = s[0:8]
	}
	return s
}

func splitKeyintoKeys(s string) []string {

	randomKeyLength := len(s)
	zerosNeeded := 32 - randomKeyLength
	for i := 0; i < zerosNeeded; i++ {
		s = strconv.Itoa(0) + s
	}
	var eightByteArray []string
	var tempString string = ""
	for i, char := range s {
		tempString = tempString + string(char)
		if i != 0 && len(tempString)%8 == 0 {
			eightByteArray = append(eightByteArray, tempString)
			tempString = ""
		}
	}

	return eightByteArray[0:4]
}

func (c cypherText) formatToString() string {
	var byteArray []byte
	for _, bits := range c {
		byteArray = append(byteArray, binaryArrayToByteSlice(bits)...)
	}

	return string(byteArray)
}
