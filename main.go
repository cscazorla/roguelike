package main

import (
	"fmt"
	_ "image/png"
	"log"

	"github.com/bytearena/ecs"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Game holds all data the entire game will need.
type Game struct {
	Debug       bool
	Map         GameMap
	World       *ecs.Manager
	WorldTags   map[string]ecs.Tag
	Turn        TurnState
	TurnCounter int
}

// NewGame creates a new Game Object and initializes the data
func NewGame() *Game {
	g := &Game{}
	g.Debug = true
	g.Map = NewGameMap()
	world, tags := InitializeWorld(g.Map.CurrentLevel)
	g.World = world
	g.WorldTags = tags
	g.Turn = PlayerTurn
	g.TurnCounter = 0
	return g
}

// Update is called on each frame loop
// The default value is 1/60 [s]
func (g *Game) Update() error {
	g.TurnCounter++
	// Update is called 60 times a second, so this
	// just adds a small delay which is
	// good enough for a turn-based game.
	if g.Turn == PlayerTurn && g.TurnCounter > 20 {
		TryMovePlayer(g)
	}
	// TODO: just for now
	g.Turn = PlayerTurn
	return nil
}

// Draw is called each on each frame loop
func (g *Game) Draw(screen *ebiten.Image) {
	//Draw the Map
	level := g.Map.CurrentLevel
	level.DrawLevel(screen)

	// Draw other renderables
	ProcessRenderables(g, level, screen)

	if g.Debug {
		gd := NewGameData()
		debug := fmt.Sprintf(
			"FPS: %.0f\nSize: %d rows x %d cols\nDimensions: %dx%dpx\nTurnCounter: %d",
			ebiten.ActualFPS(),
			gd.Cols,
			gd.Rows,
			gd.GameWidth(),
			gd.GameHeight(),
			g.TurnCounter)
		ebitenutil.DebugPrint(screen, debug)
	}

}

// Layout accepts an outside size, which is a window size on desktop,
// and returns the game's logical screen size.
func (g *Game) Layout(w, h int) (int, int) {
	gd := NewGameData()
	return gd.GameWidth(), gd.GameHeight()
}

func main() {
	g := NewGame()
	ebiten.SetWindowSize(1280, 800)
	ebiten.SetWindowTitle("Roguelike")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
