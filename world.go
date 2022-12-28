package main

import (
	"log"

	"github.com/bytearena/ecs"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var player *ecs.Component
var position *ecs.Component
var renderable *ecs.Component
var movable *ecs.Component
var monster *ecs.Component
var health *ecs.Component
var meleeWeapon *ecs.Component
var armor *ecs.Component
var name *ecs.Component

func InitializeWorld(startingLevel Level) (*ecs.Manager, map[string]ecs.Tag) {
	tags := make(map[string]ecs.Tag)
	manager := ecs.NewManager()

	// Register some components
	player = manager.NewComponent()
	position = manager.NewComponent()
	renderable = manager.NewComponent()
	movable = manager.NewComponent()
	monster = manager.NewComponent()
	health = manager.NewComponent()
	meleeWeapon = manager.NewComponent()
	armor = manager.NewComponent()
	name = manager.NewComponent()

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
		}).
		AddComponent(health, &Health{
			MaxHealth:     30,
			CurrentHealth: 30,
		}).
		AddComponent(meleeWeapon, &MeleeWeapon{
			Name:          "Fist",
			MinimumDamage: 1,
			MaximumDamage: 3,
			ToHitBonus:    2,
		}).
		AddComponent(armor, &Armor{
			Name:       "Wooden armor",
			Defense:    1,
			ArmorClass: 1,
		}).
		AddComponent(name, &Name{
			Label: "Player",
		})

	// Add a Monster in each room except the player's room
	for _, room := range startingLevel.Rooms {
		if room.X1 != startingRoom.X1 {
			mX, mY := room.Center()
			manager.NewEntity().
				AddComponent(monster, Monster{}).
				AddComponent(renderable, &Renderable{
					Image: monsterImg,
				}).
				AddComponent(position, &Position{
					X: mX,
					Y: mY,
				}).
				AddComponent(health, &Health{
					MaxHealth:     10,
					CurrentHealth: 10,
				}).
				AddComponent(meleeWeapon, &MeleeWeapon{
					Name:          "Short Sword",
					MinimumDamage: 2,
					MaximumDamage: 6,
					ToHitBonus:    0,
				}).
				AddComponent(armor, &Armor{
					Name:       "Fabric",
					Defense:    3,
					ArmorClass: 4,
				}).
				AddComponent(name, &Name{
					Label: "Ghost",
				})
		}
	}

	// Views
	players := ecs.BuildTag(player, position, health, meleeWeapon, armor, name)
	tags["players"] = players
	renderables := ecs.BuildTag(renderable, position)
	tags["renderables"] = renderables
	monsters := ecs.BuildTag(monster, position, health, meleeWeapon, armor, name)
	tags["monsters"] = monsters

	return manager, tags
}
