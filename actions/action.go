package actions

import (
	"fmt"
	"strconv"
)

var Actions = []string{
	"FIGHT",
	"RUN",
}

func ListActions() {
	for i, a := range Actions {
		fmt.Printf("action %d: %s\n", i+1, a)
	}
}

func SelectAction() (string, error) {
	num := 0
	for {
		if num > 0 && num <= len(Actions) {
			break
		}

		fmt.Println("")
		ListActions()

		fmt.Println("")
		fmt.Print("What would you like to do: ")

		var in string
		if _, err := fmt.Scan(&in); err != nil {
			return "", err
		}

		var err error
		num, err = strconv.Atoi(in)
		if err != nil {
			return "", err
		}
	}

	return Actions[num-1], nil
}
