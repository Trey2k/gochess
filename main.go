package main

import (
	"fmt"
	"image/color"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:     "Most boring chess ever!",
		Bounds:    pixel.R(0, 0, 1024, 800),
		VSync:     false,
		Resizable: false,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	board, tilePos := grid(8, 90, win.Bounds().W()/2, win.Bounds().H()/2)
	initPieces(tilePos)

	batch := pixel.NewBatch(&pixel.TrianglesData{}, spritesheet)

	var (
		frames = 0
		second = time.Tick(time.Second)
	)

	for !win.Closed() {
		win.Clear(color.RGBA{R: 54, G: 57, B: 62, A: 255})

		updateBoard(batch)
		updateMovement(win, batch)

		board.Draw(win)
		batch.Draw(win)

		win.Update()

		frames++
		select {
		case <-second:
			win.SetTitle(fmt.Sprintf("%s | FPS: %d", cfg.Title, frames))
			frames = 0
		default:
		}
	}
}

func main() {
	pixelgl.Run(run)
}
