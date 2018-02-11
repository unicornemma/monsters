package main

import (
	"fmt"
	"strings"

	"github.com/elliottpolk/monsters/actions"
	"github.com/elliottpolk/monsters/cards"
)

func main() {
	fmt.Print("Would you like to play [y/n]? ")
	var play string
	if _, err := fmt.Scan(&play); err != nil {
		panic(err)
	}

	if strings.TrimSpace(strings.ToLower(play)) != "y" {
		fmt.Println("FINE THEN!")
		return
	}

	for {

		hero, err := cards.SelectHero()
		if err != nil {
			panic(err)
		}

		fmt.Println("")
		fmt.Printf("you have selected the hero: %s\n\n", hero.Name)

		monster := cards.RandomMonster()
		fmt.Printf("UH OH! A %s has appeared!\n", monster.Name)
		fmt.Printf("monster stats: %+v\n", monster)

		action, err := actions.SelectAction()
		if err != nil {
			panic(err)
		}

		fmt.Println("")
		fmt.Printf("you have selected to: %s\n", action)

		if action == "RUN" {
			fmt.Println("FINE, you chicken!")
			return
		}

		for {
			monsterRemaining := hero.Attack(monster)
			if monsterRemaining < 1 {
				fmt.Println("You have vanquished the monster!")
				break
			}
			fmt.Printf("The %s has %d health left\n", monster.Name, monsterRemaining)

			fmt.Println("")

			heroRemaining := monster.Attack(hero)
			if heroRemaining < 1 {
				fmt.Println("The monster has defeated you!")
				break
			}
			fmt.Printf("You have %d health left\n", heroRemaining)
		}

		fmt.Print("\nWould you like to play again [y/n]? ")

		var play string
		if _, err := fmt.Scan(&play); err != nil {
			panic(err)
		}

		if strings.TrimSpace(strings.ToLower(play)) != "y" {
			fmt.Println("FINE THEN!")
			return
		}
	}

}
