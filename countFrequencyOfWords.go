// WAP to find frequency occurence of words from a string.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Enter value")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	words := strings.Split(line, " ")
	myMap := map[string]int{}
	for _, oneWord := range words {
		value, flag := myMap[oneWord]
		if flag == true {
			myMap[oneWord] = value + 1
		} else {
			myMap[oneWord] = 1
		}

	}

	fmt.Println(myMap)
}
