# Roguelike

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

This is a basic roguelike implemented in Go for learning purposes using [Roguebasin route map](http://www.roguebasin.com/index.php?title=How_to_Write_a_Roguelike_in_15_Steps) on how to write a Roguelike in 15 steps.

## Requirements

* [Go 1.19](https://go.dev/)
* [Ebitengine](github.com/hajimehoshi/ebiten/v2) for windows managing, 2D graphics, text rendering, inputs (mouse & keyboard), audio, etc.

## Assets

This game uses [Kenney's Tiny Dungeon](https://kenney.nl/assets/tiny-dungeon) tiles.

## Composition of a GameMap

The GameMap holds all the information for the entire world. The hierarchy is as follows:

* A GameMap is a collection of Dungeons
  * A Dungeon is a collection of Levels
    * A Level is a collection of MapTiles
      * A MapTile is a slice of tiles
