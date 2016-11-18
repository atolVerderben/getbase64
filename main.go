package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func loop() {
	fmt.Println("Enter a filename to be convereted to Base64\nType 'exit' to quit")
	reader := bufio.NewReader(os.Stdin)
	rawtext, _ := reader.ReadString('\n')
	text := strings.TrimSpace(rawtext)
	switch text {
	case "exit":
		os.Exit(0)
		break
	default:
		SaveBase64(text)
		break
	}
}

//SaveBase64 creates a file output of the Base64 encoding of the given file
func SaveBase64(convertFile string) {
	file, err := os.Open(convertFile)

	if err != nil {
		fmt.Printf("There was a problem with that file: %s\n", err)
		return
	}

	defer file.Close()

	fileInfo, _ := file.Stat()
	var size = fileInfo.Size()
	bytes := make([]byte, size)

	// read file into bytes
	buffer := bufio.NewReader(file)
	n, err := buffer.Read(bytes)
	check(err)

	//save to string
	s := string(bytes[:n])
	base64Text := EncodeBase64(s)

	f, err := os.Create("output/" + convertFile + "_rawB64.txt")
	check(err)
	defer f.Close()

	f2, _ := os.Create("output/reEncoded_" + convertFile)
	defer f2.Close()
	newImg := DecodeBase64(base64Text)
	f2.Write(newImg)
	_, err = f.WriteString(base64Text)

	fmt.Printf("%s Successfully Converted to Base64!\n", convertFile)
}

func main() {
	for {
		loop()
	}
}

//EncodeBase64 returns a string representation of base64
func EncodeBase64(message string) string {
	base64Text := make([]byte, base64.StdEncoding.EncodedLen(len(message)))
	base64.StdEncoding.Encode(base64Text, []byte(message))
	return string(base64Text)
}

//DecodeBase64 decodes the given string into a byte slice
func DecodeBase64(message string) []byte {
	base64Text := make([]byte, base64.StdEncoding.DecodedLen(len(message)))
	l, _ := base64.StdEncoding.Decode(base64Text, []byte(message))
	return base64Text[:l]
}
