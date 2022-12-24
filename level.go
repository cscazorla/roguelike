package main

import (
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/norendren/go-fov/fov"
)

type TileType int

const (
	WALL TileType = iota
	FLOOR
)

// Each of the map tiles will be represented  by one of these structures
type MapTile struct {
	PixelX     int // Upper left corner of the tile
	PixelY     int
	Blocked    bool          // The tile should block the player or monster ?
	Image      *ebiten.Image // Pointer to an ebiten Image
	IsRevealed bool          // Has the tile has, at one time, been revealed to us by FoV?
	TileType   TileType
}

// Level holds the tile information for a complete dungeon level.
type Level struct {
	Tiles     []MapTile
	Rooms     []Rect
	PlayerFoV *fov.View
}

// NewLevel creates a new game level in a dungeon.
func NewLevel() Level {
	l := Level{}
	rooms := make([]Rect, 0)
	l.Rooms = rooms
	l.GenerateLevelTiles()
	l.PlayerFoV = fov.New()
	return l
}

// Tiles will be stoted in one slice. We will use GetIndexFromXY to
// determine which tile to return.
// GetIndexFromXY gets the index of the map array from a given X,Y TILE coordinate.
// This coordinate is logical tiles, not pixels.
func (level *Level) GetIndexFromXY(x int, y int) int {
	gd := NewGameData()
	return (y * gd.Cols) + x
}

// createTiles creates a map of all walls as a baseline for carving out a level.
func (level *Level) createTiles() []MapTile {
	gd := NewGameData()
	tiles := make([]MapTile, gd.Rows*gd.Cols)
	index := 0
	for x := 0; x < gd.Cols; x++ {
		for y := 0; y < gd.Rows; y++ {
			index = level.GetIndexFromXY(x, y)
			wall, _, err := ebitenutil.NewImageFromFile("assets/wall.png")
			if err != nil {
				log.Fatal(err)
			}
			tile := MapTile{
				PixelX:     x * gd.TileWidth,
				PixelY:     y * gd.TileHeight,
				Blocked:    true,
				Image:      wall,
				IsRevealed: false,
				TileType:   WALL,
			}
			tiles[index] = tile
		}
	}
	return tiles
}

func (level *Level) createRoom(room Rect) {
	for y := room.Y1 + 1; y < room.Y2; y++ {
		for x := room.X1 + 1; x < room.X2; x++ {
			index := level.GetIndexFromXY(x, y)
			level.Tiles[index].Blocked = false
			level.Tiles[index].TileType = FLOOR
			floor, _, err := ebitenutil.NewImageFromFile("assets/floor.png")
			if err != nil {
				log.Fatal(err)
			}
			level.Tiles[index].Image = floor
		}
	}
}

// GenerateLevelTiles creates a new Dungeon Level Map.
func (level *Level) GenerateLevelTiles() {
	MIN_SIZE := 6
	MAX_SIZE := 10
	MAX_ROOMS := 30

	gd := NewGameData()
	tiles := level.createTiles()
	level.Tiles = tiles
	contains_rooms := false

	for idx := 0; idx < MAX_ROOMS; idx++ {
		w := GetRandomBetween(MIN_SIZE, MAX_SIZE)
		h := GetRandomBetween(MIN_SIZE, MAX_SIZE)
		x := GetDiceRoll(gd.Cols - w - 1)
		y := GetDiceRoll(gd.Rows - h - 1)

		new_room := NewRect(x, y, w, h)
		okToAdd := true
		for _, otherRoom := range level.Rooms {
			if new_room.Intersect(otherRoom) {
				okToAdd = false
				break
			}
		}
		if okToAdd {
			level.createRoom(new_room)
			if contains_rooms {
				newX, newY := new_room.Center()
				prevX, prevY := level.Rooms[len(level.Rooms)-1].Center()

				coinflip := GetDiceRoll(2)

				if coinflip == 2 {
					level.createHorizontalTunnel(prevX, newX, prevY)
					level.createVerticalTunnel(prevY, newY, newX)
				} else {
					level.createHorizontalTunnel(prevX, newX, newY)
					level.createVerticalTunnel(prevY, newY, prevX)
				}
			}
			level.Rooms = append(level.Rooms, new_room)
			contains_rooms = true
		}
	}
}

func (level *Level) createHorizontalTunnel(x1 int, x2 int, y int) {
	gd := NewGameData()
	for x := math.Min(float64(x1), float64(x2)); x < math.Max(float64(x1), float64(x2))+1; x++ {
		index := level.GetIndexFromXY(int(x), y)
		if index > 0 && index < gd.Rows*gd.Cols {
			level.Tiles[index].Blocked = false
			level.Tiles[index].TileType = FLOOR
			floor, _, err := ebitenutil.NewImageFromFile("assets/floor.png")
			if err != nil {
				log.Fatal(err)
			}
			level.Tiles[index].Image = floor
		}
	}
}

func (level *Level) createVerticalTunnel(y1 int, y2 int, x int) {
	gd := NewGameData()
	for y := math.Min(float64(y1), float64(y2)); y < math.Max(float64(y1), float64(y2))+1; y++ {
		index := level.GetIndexFromXY(x, int(y))
		if index > 0 && index < gd.Rows*gd.Cols {
			level.Tiles[index].TileType = FLOOR
			level.Tiles[index].Blocked = false
			floor, _, err := ebitenutil.NewImageFromFile("assets/floor.png")
			if err != nil {
				log.Fatal(err)
			}
			level.Tiles[index].Image = floor
		}
	}
}

func (level Level) InBounds(x, y int) bool {
	gd := NewGameData()
	if x < 0 || x > gd.Cols || y < 0 || y > gd.Rows {
		return false
	}
	return true
}

func (level Level) IsOpaque(x, y int) bool {
	idx := level.GetIndexFromXY(x, y)
	return level.Tiles[idx].TileType == WALL
}

func (level *Level) DrawLevel(screen *ebiten.Image) {
	gd := NewGameData()
	for x := 0; x < gd.Cols; x++ {
		for y := 0; y < gd.Rows; y++ {
			idx := level.GetIndexFromXY(x, y)
			tile := level.Tiles[idx]
			isVisible := level.PlayerFoV.IsVisible(x, y)
			if isVisible {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(tile.PixelX), float64(tile.PixelY))
				screen.DrawImage(tile.Image, op)
				level.Tiles[idx].IsRevealed = true
			} else if tile.IsRevealed {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(tile.PixelX), float64(tile.PixelY))
				op.ColorM.Scale(1.0, 1.0, 1.0, 0.25)
				screen.DrawImage(tile.Image, op)
			}
		}
	}
}
