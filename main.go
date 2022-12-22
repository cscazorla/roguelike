package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

// Game holds all data the entire game will need.
type Game struct {
	Map GameMap
}

// NewGame creates a new Game Object and initializes the data
func NewGame() *Game {
	g := &Game{}
	g.Map = NewGameMap()
	return g
}

// Update is called on each frame loop
func (g *Game) Update() error {
	return nil
}

// Draw is called each on each frame loop
func (g *Game) Draw(screen *ebiten.Image) {
	gd := NewGameData()

	//Draw the Map
	level := g.Map.Dungeons[0].Levels[0]
	for x := 0; x < gd.ScreenWidth; x++ {
		for y := 0; y < gd.ScreenHeight; y++ {
			tile := level.Tiles[level.GetIndexFromXY(x, y)]
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(tile.PixelX), float64(tile.PixelY))
			screen.DrawImage(tile.Image, op)
		}
	}
}

// Layout will return the screen dimensions.
func (g *Game) Layout(w, h int) (int, int) {
	gd := NewGameData()
	return gd.TileWidth * gd.ScreenWidth, gd.TileHeight * gd.ScreenHeight
}

func main() {
	g := NewGame()
	ebiten.SetWindowSize(1280, 800)
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowTitle("Roguelike")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
