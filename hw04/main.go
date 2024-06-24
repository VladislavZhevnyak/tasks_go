package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"unicode/utf8"
)

func RemoveSpecialCharacters(input string) string {
	reg := regexp.MustCompile("[^a-zA-Z]+")
	newInput := reg.ReplaceAllString(input, "")
	return newInput
}

func inputString() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	input = RemoveSpecialCharacters(input)
	if err != nil {
		return "", err
	}
	return input, nil
}

func countPerLine(s string) {
	lengthOfString := utf8.RuneCountInString(s)
	count := 1
	for i := 1; i < lengthOfString; i++ {
		if s[i] == s[i-1] {
			count++
			if i == lengthOfString-1 {
				fmt.Printf("%s%d", string(s[i-1]), count)
			}
		} else {
			fmt.Printf("%s%d", string(s[i-1]), count)
			count = 1
		}
	}

}

func main() {
	str, err := inputString()
	if err != nil {
		log.Fatal(err)
	}
	countPerLine(str)
}
