package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
	"unicode/utf8"
)

func RemoveSpecialCharacters(input string) string {
	//Delete special characters from string
	reg := regexp.MustCompile("[^a-zA-Z]+")
	newInput := reg.ReplaceAllString(input, "")
	return newInput
}

func inputString() (string, error) {
	//Entering a string with special characters removed
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	input = strings.ToLower(input)
	input = RemoveSpecialCharacters(input)
	if err != nil {
		return "", err
	}
	return input, nil
}

/*
func countPerLine(s string) {
	//Count characters in string without using collections
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
*/

func printOrderedCollection(str string) {
	//Count characters in string with using collections
	lengthOfString := utf8.RuneCountInString(str)
	countMap := make(map[string]int)
	for i := 0; i < lengthOfString; i++ {
		countMap[string(str[i])]++
	}
	//Add keys to slice and sort it
	sortedKeys := make([]string, 0, len(countMap))
	for key := range countMap {
		sortedKeys = append(sortedKeys, key)
	}
	sort.Strings(sortedKeys)
	//printing in alphabetical order
	for i := 0; i < len(sortedKeys); i++ {
		fmt.Printf("%s%d", sortedKeys[i], countMap[sortedKeys[i]])
	}
}

func main() {
	str, err := inputString()
	if err != nil {
		log.Fatal(err)
	}
	printOrderedCollection(str)
}
