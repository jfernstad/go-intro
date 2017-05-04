package main

import (
	"errors"
	"fmt"
	"strconv"
)

// Player object declaration
type Player struct {
	Name         string
	FavoriteGame string
	Highscore    uint64
}

// Quaker is composed of Player
type Quaker struct {
	Player      // Member variables name and type is the same
	FavoriteMap string
}

// Construct is a method on the Player object
func (p *Player) Construct(name string, game string, score uint64) {
	p.Name = name
	p.FavoriteGame = game
	p.Highscore = score
}

// Summary does not mutate Player instance and returns a string and an error
//
// func (t T) NAME (arguments) (RET1, RET2)
//
func (p Player) Summary() (string, error) {
	if p.Highscore > 1000 {
		// Return multiple values, separated by comma
		return "", errors.New(p.Name + " is a cheater!")
	} else {
		return p.Name + " loves " + p.FavoriteGame + " and highest score " + strconv.FormatUint(p.Highscore, 10), nil
		// return fmt.Sprintf("%s loves %s and highest score is %d", p.Name, p.FavoriteGame, p.Highscore), nil
	}
}

// Summary "overloads" Players' `Summary` method
// func (p Quaker) Summary() (string, error) {
// 	return p.Name + " is crazy about " + p.FavoriteMap + " and highest score " + strconv.FormatUint(p.Highscore, 10), nil
// }

func main() {
	me := Quaker{}

	// Using Players construct
	me.Construct("GoodMan", "Quake I", 10)

	// Initialize the Quaker
	me = Quaker{Player{"BadMan", "Quake III", 10}, "q3dm17"}
	// This is equivalent, named assignment
	// me = Quaker{Player: Player{"BadMan", "Quake", 0}, FavoriteMap: "q3dm17"}

	me.Highscore = 100

	// Composition at work
	sum, _ := me.Summary()

	// When Quaker defines `Summary`
	// You can access the "overloaded" `Summary`
	// using the parent
	// sumOrig, _ := me.Player.Summary()

	fmt.Printf("%s\n", sum)
	// fmt.Printf("%s\n", sumOrig)

}
