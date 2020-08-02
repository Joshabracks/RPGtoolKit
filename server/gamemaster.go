package main

// NewGame constructs a new game object and assigns it to a "gm" player.  
// Will have pointers to players, world chunks, world maps, characters and all other game objects.
// 
func NewGame (gm User) {
	var g Game
	g.gm = gm
}