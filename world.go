package main

import (
	"fmt"
	"log"

	"github.com/bytearena/ecs"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var position *ecs.Component
var renderable *ecs.Component
var monster *ecs.Component

func InitializeWorld(startingLevel Level) (*ecs.Manager, map[string]ecs.Tag) {
	tags := make(map[string]ecs.Tag)
	manager := ecs.NewManager()

	// Register some components
	player := manager.NewComponent()
	position = manager.NewComponent()
	renderable = manager.NewComponent()
	movable := manager.NewComponent()
	monster = manager.NewComponent()

	// Load images for player and monster
	playerImg, _, err := ebitenutil.NewImageFromFile("assets/player.png")
	if err != nil {
		log.Fatal(err)
	}
	monsterImg, _, err := ebitenutil.NewImageFromFile("assets/monster.png")
	if err != nil {
		log.Fatal(err)
	}

	// Get First Room
	startingRoom := startingLevel.Rooms[0]
	x, y := startingRoom.Center()

	// Create entity for our player
	manager.NewEntity().
		AddComponent(player, Player{}).
		AddComponent(movable, Movable{}).
		AddComponent(renderable, &Renderable{
			Image: playerImg,
		}).
		AddComponent(position, &Position{
			X: x,
			Y: y,
		})

	// Add a Monster in each room except the player's room
	for idx, room := range startingLevel.Rooms {
		if room.X1 != startingRoom.X1 {
			mX, mY := room.Center()
			manager.NewEntity().
				AddComponent(monster, &Monster{
					Name: "Ghost " + fmt.Sprint(idx),
				}).
				AddComponent(renderable, &Renderable{
					Image: monsterImg,
				}).
				AddComponent(position, &Position{
					X: mX,
					Y: mY,
				})
		}
	}

	// Views
	players := ecs.BuildTag(player, position)
	tags["players"] = players
	renderables := ecs.BuildTag(renderable, position)
	tags["renderables"] = renderables
	monsters := ecs.BuildTag(monster, position)
	tags["monsters"] = monsters

	return manager, tags
}
