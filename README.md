# Roguelike

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

This is a basic roguelike implemented in Go for learning purposes. This implementation follows [Roguebasin route map](http://www.roguebasin.com/index.php?title=How_to_Write_a_Roguelike_in_15_Steps) on how to write a Roguelike in 15 steps.

![My first Roguelike in Go](images/screnshoot.png)

## Requirements

* [Go 1.19](https://go.dev/)
* [Ebitengine](https://github.com/hajimehoshi/ebiten) for windows managing, 2D graphics, text rendering, inputs (mouse & keyboard), audio, etc.
* [bytearena/ecs](https://github.com/ByteArena/ecs) for the Go implementation of the Entity/Component/System paradigm.

## Assets

This game uses [Kenney's Tiny Dungeon](https://kenney.nl/assets/tiny-dungeon) tiles.

## Roadmap

- [x] Project structure
- [x] Basic MapTiles
- [x] Adding ECS capabilities
- [x] Collisions with walls
- [x] Rooms
- [x] Corridors
- [x] Turn based
- [x] Field of View
- [x] Monsters
- [x] Monsters Pathfinding
- [x] Basic combat
- [ ] UI
- [ ] Player HUD

## Composition of a GameMap

The GameMap holds all the information for the entire world. The hierarchy is as follows:

* A GameMap is a collection of Dungeons
  * A Dungeon is a collection of Levels
    * A Level is a collection of MapTiles
      * A MapTile is a slice of tiles
