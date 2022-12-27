package main

import (
	"github.com/norendren/go-fov/fov"
)

func UpdateMonster(game *Game) {
	l := game.Map.CurrentLevel
	playerPosition := Position{}

	for _, plr := range game.World.Query(game.WorldTags["players"]) {
		pos := plr.Components[position].(*Position)
		playerPosition.X = pos.X
		playerPosition.Y = pos.Y
	}

	for _, result := range game.World.Query(game.WorldTags["monsters"]) {
		monsterPos := result.Components[position].(*Position)
		monsterFoV := fov.New()
		monsterFoV.Compute(l, monsterPos.X, monsterPos.Y, 8)
		if monsterFoV.IsVisible(playerPosition.X, playerPosition.Y) {
			astar := AStar{}
			path := astar.GetPath(l, monsterPos, &playerPosition)
			if len(path) > 1 {
				nextTile := l.Tiles[l.GetIndexFromXY(path[1].X, path[1].Y)]
				if !nextTile.Blocked {
					l.Tiles[l.GetIndexFromXY(monsterPos.X, monsterPos.Y)].Blocked = false
					monsterPos.X = path[1].X
					monsterPos.Y = path[1].Y
					nextTile.Blocked = true
				}
			}
		}
	}
	game.Turn = PlayerTurn
}
