package cards

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/elliottpolk/monsters/dice"
)

type MonsterCard Card

var Monsters = []*MonsterCard{
	&MonsterCard{
		Name:     "Dragon",
		Strength: 99,
		Health:   400,
	},
	&MonsterCard{
		Name:     "Zombie",
		Strength: 80,
		Health:   10,
	},
	&MonsterCard{
		Name:     "Spider",
		Strength: 20,
		Health:   340,
	},
	&MonsterCard{
		Name:     "Skeleton",
		Strength: 60,
		Health:   90,
	},
	&MonsterCard{
		Name:     "Creeper",
		Strength: 80,
		Health:   300,
	},
}

func RandomMonster() *MonsterCard {
	rando := rand.New(rand.NewSource(time.Now().UnixNano()))
	return Monsters[rando.Intn(len(Monsters))]
}
func (m *MonsterCard) Attack(hero *HeroCard) int64 {
	roll := dice.Roll()
	hero.Health -= roll

	fmt.Printf("The monster, %s, has attacked your hero with %d damage\n", m.Name, roll)
	return hero.Health
}
