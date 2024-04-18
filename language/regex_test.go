package gotutorial

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func TestRegex(t *testing.T) {
	str1 := "this is a [sample] [[string]] with [SOME] special words"
	re := regexp.MustCompile(`\[([^\[\]]*)\]`)
	fmt.Printf("Pattern: %v\n", re.String())      // print pattern
	fmt.Println("Matched:", re.MatchString(str1)) // true
	fmt.Println("\nText between square brackets:")
	submatchall := re.FindAllString(str1, -1)
	for _, element := range submatchall {
		element = strings.Trim(element, "[")
		element = strings.Trim(element, "]")
		fmt.Println(element)
	}
}

func TestRegex2(t *testing.T) {
	str1 := "We @@@Love@@@@ #Go!$! ****Programming****Language^^^"
	re := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	fmt.Printf("Pattern: %v\n", re.String()) // print pattern
	fmt.Println(re.MatchString(str1))        // true

	t1, _ := re.MarshalText()
	fmt.Println(string(t1))

	submatchall := re.FindAllString(str1, -1)
	for _, element := range submatchall {
		fmt.Println(element)
	}
}
