package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Each of the map tiles will be represented  by one of these structures
type MapTile struct {
	PixelX  int // Upper left corner of the tile
	PixelY  int
	Blocked bool          // The tile should block the player or monster ?
	Image   *ebiten.Image // Pointer to an ebiten Image
}

// Level holds the tile information for a complete dungeon level.
type Level struct {
	Tiles []MapTile
}

// NewLevel creates a new game level in a dungeon.
func NewLevel() Level {
	l := Level{}
	tiles := l.CreateTiles()
	l.Tiles = tiles
	return l
}

// Tiles will be stoted in one slice. We will use GetIndexFromXY to
// determine which tile to return.
// GetIndexFromXY gets the index of the map array from a given X,Y TILE coordinate.
// This coordinate is logical tiles, not pixels.
func (level *Level) GetIndexFromXY(x int, y int) int {
	gd := NewGameData()
	return (y * gd.ScreenWidth) + x
}

// CreateTiles creates the initial map. It consists of all floor tiles
// except the outer bounds of the window, which will be walls.
func (level *Level) CreateTiles() []MapTile {
	gd := NewGameData()
	tiles := make([]MapTile, gd.ScreenHeight*gd.ScreenWidth)
	index := 0

	for x := 0; x < gd.ScreenWidth; x++ {
		for y := 0; y < gd.ScreenHeight; y++ {
			index = level.GetIndexFromXY(x, y)
			if x == 0 || x == gd.ScreenWidth-1 || y == 0 || y == gd.ScreenHeight-1 {
				wall, _, err := ebitenutil.NewImageFromFile("assets/wall.png")
				if err != nil {
					log.Fatal(err)
				}
				tile := MapTile{
					PixelX:  x * gd.TileWidth,
					PixelY:  y * gd.TileHeight,
					Blocked: true,
					Image:   wall,
				}
				tiles[index] = tile
			} else {
				floor, _, err := ebitenutil.NewImageFromFile("assets/floor.png")
				if err != nil {
					log.Fatal(err)
				}
				tile := MapTile{
					PixelX:  x * gd.TileWidth,
					PixelY:  y * gd.TileHeight,
					Blocked: false,
					Image:   floor,
				}
				tiles[index] = tile
			}
		}
	}

	return tiles
}

func (level *Level) DrawLevel(screen *ebiten.Image) {
	gd := NewGameData()
	for x := 0; x < gd.ScreenWidth; x++ {
		for y := 0; y < gd.ScreenHeight; y++ {
			tile := level.Tiles[level.GetIndexFromXY(x, y)]
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(tile.PixelX), float64(tile.PixelY))
			screen.DrawImage(tile.Image, op)
		}
	}
}
