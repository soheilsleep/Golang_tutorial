package main

import "fmt"

type Player struct {
	Name     string
	Age      int
	Weight   int
	Height   int
	Position string
}
type Runner interface {
	Run()
}
type Shooter interface {
	Shoot()
}
type Passer interface {
	Pass()
}

func main() {
	player1 := &Player{
		Name:     "John",
		Age:      20,
		Weight:   10,
		Height:   20,
		Position: "Forward",
	}
	var runner Runner = player1
	var shooter Shooter = player1
	var passer Passer = player1

	runner.Run()
	shooter.Shoot()
	passer.Pass()

}
func (player *Player) Run() {
	fmt.Printf("name : %s, position: %s, player is running\n", player.Name, player.Position)
}
func (player *Player) Shoot() {
	fmt.Printf("name : %s, position: %s, player is shooting\n", player.Name, player.Position)
}
func (player *Player) Pass() {
	fmt.Printf("name : %s,  position: %s , player is passing\n", player.Name, player.Position)

}
