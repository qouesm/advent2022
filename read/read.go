package read

import (
	"io/ioutil"
	"log"
	"strings"
)

// returns content of fileName as one string
func ReadFile(fileName string) string {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Panicf("Could not read input\n%v", err)
	}
	return string(content)
}

// returns the content of fileName as a slice of strings split by \n
func ReadLines(fileName string) []string {
	v := strings.Split(ReadFile(fileName), "\n")
	return v[:len(v)-1] // trim ending blank line
}
