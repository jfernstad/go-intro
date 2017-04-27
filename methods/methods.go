package main

import (
	"errors"
	"fmt"
)

// Player object declaration
type Player struct {
	Name         string
	FavoriteGame string
	Highscore    uint64
}

// Construct is a method on the Player object
//
// func (var T) NAME (arguments)
// func (var *T) NAME (arguments)  // `var *T` if the method will mutate the object, equivalent to `this`.
//
func (p *Player) Construct(name string, game string, _ uint64) {
	// Ignore arguments with underscore `_`

	p.Name = name
	p.FavoriteGame = game
	p.Highscore = 0
}

// Summary does not mutate Player instance and returns a string and an error
//
// func (var T) NAME (arguments) (RET1, RET2)
//
func (p Player) Summary() (string, error) {
	if p.Highscore > 1000 {
		// Return multiple values, separated by comma
		return "", errors.New("Cheater!")
	} else {
		return p.Name + " loves " + p.FavoriteGame + " and highest score " + string(p.Highscore), nil
		// return fmt.Sprintf("%s loves %s and highest score is %d", p.Name, p.FavoriteGame, p.Highscore), nil
	}
}

func main() {
	me := Player{} // Null state for object

	// These are equivalent
	me = Player{"BadMan", "Quake", 0}
	me.Construct("BadMan", "Quake", 999) // 999 is ignored
	// me.Construct("BadMan", "Quake 3")   // This is an error

	me.Highscore = 1337

	// sum, _ := me.Summary() // Ignore error
	// _, _ := me.Summary()   // Ignore all return values
	sum, err := me.Summary()

	// Standard error handling in golang
	if err != nil {
		fmt.Print(err)
	}

	fmt.Printf("%s\n", sum)
	// Prints: Cheater!

}
