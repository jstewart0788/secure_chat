package main

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func generateAesKey() uint64 {
	key := make([]byte, 64)
	rand.Read(key)
	aesKey := binary.BigEndian.Uint64(key)
	return aesKey
}

func generateIV() uint64 {
	key := make([]byte, 16)
	rand.Read(key)
	Iv := binary.BigEndian.Uint64(key)
	return Iv
}

func saveAESKeyToFile(k uint64) error {
	return ioutil.WriteFile("./keys/aes/private", []byte(strconv.FormatUint(k, 10)), 0666)
}
func saveIvToFile(k uint64) error {
	return ioutil.WriteFile("./keys/iv/public", []byte(strconv.FormatUint(k, 10)), 0666)
}

func getAesKeyFromFile() string {
	key, err := ioutil.ReadFile("./keys/aes/private")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	return string(key)
}

func getIvFromFile() string {
	key, err := ioutil.ReadFile("./keys/iv/private")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	return string(key)
}
