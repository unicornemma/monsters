package main

import (
	"image"
	_ "image/png"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

// "github.com/unicornemma/monsters/actions"
// "github.com/unicornemma/monsters/cards"
func loadImg(path string) (pixel.Picture, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	return pixel.PictureDataFromImage(img), nil
}

func main() {
	pixelgl.Run(func() {
		cfg := pixelgl.WindowConfig{
			Title:  "Pixel Rocks!",
			Bounds: pixel.R(0, 0, 1024, 768),
			VSync:  true,
		}
		win, err := pixelgl.NewWindow(cfg)
		if err != nil {
			panic(err)
		}

		win.Clear(colornames.Hotpink)

		imgs := []string{
			"assets/textures/gopher_pink.png",
			"assets/textures/gopher_blue.png",
			"assets/textures/gopher_purple.png",
			"assets/textures/gopher_yellow.png",
		}

		for i, color := range imgs {
			img, err := loadImg(color)
			if err != nil {
				panic(err)
			}

			center := win.Bounds().Center()
			center.X -= img.Bounds().W() * 1.5

			sprite := pixel.NewSprite(img, img.Bounds())
			sprite.Draw(win, pixel.IM.Moved(center.Add(pixel.V(img.Bounds().W()*float64(i), 0))))
		}

		for !win.Closed() {
			win.Update()
		}

	})
}

// func main() {
// 	fmt.Print("Would you like to play [y/n]? ")
// 	var play string
// 	if _, err := fmt.Scan(&play); err != nil {
// 		panic(err)
// 	}

// 	if strings.TrimSpace(strings.ToLower(play)) != "y" {
// 		fmt.Println("FINE THEN!")
// 		return
// 	}

// 	for {

// 		hero, err := cards.SelectHero()
// 		if err != nil {
// 			panic(err)
// 		}

// 		fmt.Println("")
// 		fmt.Printf("you have selected the hero: %s\n\n", hero.Name)

// 		monster := cards.RandomMonster()
// 		fmt.Printf("UH OH! A %s has appeared!\n", monster.Name)
// 		fmt.Printf("monster stats: %+v\n", monster)

// 		action, err := actions.SelectAction()
// 		if err != nil {
// 			panic(err)
// 		}

// 		fmt.Println("")
// 		fmt.Printf("you have selected to: %s\n", action)

// 		if action == "RUN" {
// 			fmt.Println("FINE, you chicken!")
// 			return
// 		}

// 		for {
// 			monsterRemaining := hero.Attack(monster)
// 			if monsterRemaining < 1 {
// 				fmt.Println("You have vanquished the monster!")
// 				break
// 			}
// 			fmt.Printf("The %s has %d health left\n", monster.Name, monsterRemaining)

// 			fmt.Println("")

// 			heroRemaining := monster.Attack(hero)
// 			if heroRemaining < 1 {
// 				fmt.Println("The monster has defeated you!")
// 				break
// 			}
// 			fmt.Printf("You have %d health left\n", heroRemaining)
// 		}

// 		fmt.Print("\nWould you like to play again [y/n]? ")

// 		var play string
// 		if _, err := fmt.Scan(&play); err != nil {
// 			panic(err)
// 		}

// 		if strings.TrimSpace(strings.ToLower(play)) != "y" {
// 			fmt.Println("FINE THEN!")
// 			return
// 		}
// 	}

// }
