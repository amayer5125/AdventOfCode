package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// return a list of all elements in the group
func getGroupList(d map[string][]string, index string, tried map[string]bool) (list []string) {
	if tried[index] {
		return
	}

	list = append(list, index)
	tried[index] = true

	for _, v := range d[index] {
		list = append(list, getGroupList(d, v, tried)...)
	}

	return
}

func main() {
	var d = make(map[string][]string)

	// get the input
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		tmp := strings.Split(scanner.Text(), " <-> ")
		d[tmp[0]] = strings.Split(tmp[1], ", ")
	}

	fmt.Println(len(getGroupList(d, "0", map[string]bool{})))
}
