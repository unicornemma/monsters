package cards

import (
	"fmt"
	"strconv"

	"github.com/elliottpolk/monsters/dice"
)

type HeroCard Card

var Heroes = []*HeroCard{
	&HeroCard{
		Name:     "Batman",
		Strength: 90,
		Health:   60,
	},
	&HeroCard{
		Name:     "SuperMan",
		Strength: 100,
		Health:   500,
	},
	&HeroCard{
		Name:     "Wonder Woman",
		Strength: 100,
		Health:   400,
	},
	&HeroCard{
		Name:     "Spider-Man",
		Strength: 90,
		Health:   99,
	},
	&HeroCard{
		Name:     "MR-MAN",
		Strength: 100,
		Health:   500,
	},
	&HeroCard{
		Name:     "PrincessRobot",
		Strength: 100,
		Health:   500,
	},
}

func ListHeroes() {
	for i, h := range Heroes {
		fmt.Printf("card %d: %+v\n", i+1, h)
	}
}

func SelectHero() (*HeroCard, error) {
	num := 0
	for {
		if num > 0 && num <= len(Heroes) {
			break
		}

		fmt.Println("")
		ListHeroes()

		fmt.Println("")
		fmt.Print("Pick your hero card: ")

		var in string
		if _, err := fmt.Scan(&in); err != nil {
			return nil, err
		}

		var err error
		num, err = strconv.Atoi(in)
		if err != nil {
			return nil, err
		}
	}

	return Heroes[num-1], nil
}

func (h *HeroCard) Attack(monster *MonsterCard) int64 {
	roll := dice.Roll()
	monster.Health -= roll

	fmt.Printf("Your hero, %s, has attacked the %s with %d damage\n", h.Name, monster.Name, roll)
	return monster.Health
}
