package main

import (
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

type square struct {
	occupied bool
	occupint *piece
	max      pixel.Vec
	min      pixel.Vec
	pos      pixel.Vec
	up       int
	side     int
}

type row []square

type boardStruct []row

var board boardStruct

func grid(gridSize int, tileScale, posX, posY float64) (*imdraw.IMDraw, [][]pixel.Vec) {
	var returnSquarePos [][]pixel.Vec

	grid := imdraw.New(nil)

	black := true
	posX -= (tileScale * float64(gridSize/2))
	posY += (tileScale * float64(gridSize/2))
	x := posX
	y := posY

	for i := 0; i < gridSize; i++ {

		var squarePos []pixel.Vec
		var apRow row

		for j := 0; j < gridSize; j++ {
			if black {
				grid.Color = color.RGBA{R: 79, G: 50, B: 2, A: 255}
				black = false
			} else {
				grid.Color = color.RGBA{R: 230, G: 230, B: 190, A: 255}
				black = true
			}

			grid.Push(pixel.V(x, y), pixel.V(x+tileScale, y-tileScale))
			grid.Rectangle(0)

			midX := (x + (x + tileScale)) / 2
			midY := (y + (y - tileScale)) / 2

			squarePos = append(squarePos, pixel.V(midX, midY))
			apRow = append(apRow, square{
				occupied: false,
				max:      pixel.V(midX+(tileScale/2), midY+(tileScale/2)),
				min:      pixel.V(midX-(tileScale/2), midY-(tileScale/2)),
				pos:      pixel.V(midX, midY),
				up:       i,
				side:     j,
			})

			x += tileScale
		}
		if gridSize%2 == 0 {
			if black {
				black = false
			} else {
				black = true
			}
		}

		returnSquarePos = append(returnSquarePos, squarePos)
		board = append(board, apRow)

		x = posX
		y -= tileScale
	}

	grid.Color = color.Black
	grid.EndShape = imdraw.RoundEndShape
	grid.Push(pixel.V(x, y), pixel.V(posX+(tileScale*float64(gridSize)), posY+(tileScale/float64(gridSize+1))/2))
	grid.Rectangle(10)

	return grid, returnSquarePos
}
