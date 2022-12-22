package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	_ "image/png"
	"log"
)

// Game holds all data the entire game will need.
type Game struct {
	Tiles []MapTile
}

// NewGame creates a new Game Object and initializes the data
func NewGame() *Game {
	g := &Game{}
	g.Tiles = CreateTiles()
	return g
}

// Update is called on each fram
func (g *Game) Update() error {
	return nil
}

// Draw is called each draw cycle and is where we will blit.
func (g *Game) Draw(screen *ebiten.Image) {
	gd := NewGameData()

	//Draw the Map
	for x := 0; x < gd.ScreenWidth; x++ {
		for y := 0; y < gd.ScreenHeight; y++ {
			tile := g.Tiles[GetIndexFromXY(x, y)]
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(tile.PixelX), float64(tile.PixelY))
			screen.DrawImage(tile.Image, op)
		}
	}
}

// Layout will return the screen dimensions.
func (g *Game) Layout(w, h int) (int, int) { return 1280, 800 }

func main() {
	g := NewGame()
	ebiten.SetWindowSize(1280, 800)
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowTitle("Roguelike")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
