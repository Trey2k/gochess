package main

import (
	"image"
	_ "image/png"
	"os"

	"github.com/faiface/pixel"
)

type piece struct {
	pos         pixel.Vec
	frame       pixel.Rect
	class       int
	tileNumUp   int
	tileNumSide int
	moves       int
	black       bool
	dead        bool
}

type pieceSet struct {
	pieces []piece
}

var whiteSet pieceSet
var blackSet pieceSet

var spritesheet pixel.Picture
var pieceFrames [2][6]pixel.Rect
var updatePieces bool
var initBoard bool

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func initPieces(tilePos [][]pixel.Vec) {
	var err error
	spritesheet, err = loadPicture("./assets/chessSprite.png")
	if err != nil {
		panic(err)
	}

	i := 0
	for y := spritesheet.Bounds().Min.Y; y < spritesheet.Bounds().Max.Y; y += 333.5 {
		j := 0
		for x := spritesheet.Bounds().Min.X; x < spritesheet.Bounds().Max.X; x += 333.5 {
			pieceFrames[i][j] = pixel.R(x, y, x+333.5, y+333.5)
			j++
		}
		i++
	}

	whiteSet = initSet(pieceFrames, tilePos, false)
	blackSet = initSet(pieceFrames, tilePos, true)

	initBoard = true
	moving = false
	movingPiece = nil
}

func (set *pieceSet) init(batch *pixel.Batch) {
	for i := 0; i < len(set.pieces); i++ {
		if !set.pieces[i].dead {
			cehssPiece := pixel.NewSprite(spritesheet, set.pieces[i].frame)
			cehssPiece.Draw(batch, pixel.IM.Moved(set.pieces[i].pos).Scaled(set.pieces[i].pos, 0.3))
			board[set.pieces[i].tileNumUp][set.pieces[i].tileNumSide].occupied = true
			board[set.pieces[i].tileNumUp][set.pieces[i].tileNumSide].occupint = &set.pieces[i]
		}
	}
}

func (set *pieceSet) update(batch *pixel.Batch) {
	for i := 0; i < len(set.pieces); i++ {
		if !set.pieces[i].dead {
			cehssPiece := pixel.NewSprite(spritesheet, set.pieces[i].frame)
			cehssPiece.Draw(batch, pixel.IM.Moved(set.pieces[i].pos).Scaled(set.pieces[i].pos, 0.3))
		}
	}
}

func (p *piece) moveUpdate(batch *pixel.Batch) {
	batch.Clear()
	whiteSet.update(batch)
	blackSet.update(batch)
	cehssPiece := pixel.NewSprite(spritesheet, p.frame)
	cehssPiece.Draw(batch, pixel.IM.Moved(p.pos).Scaled(p.pos, 0.4))
}

func updateBoard(batch *pixel.Batch) {
	if initBoard {
		batch.Clear()
		whiteSet.init(batch)
		blackSet.init(batch)
		initBoard = false
	}

	if updatePieces {
		batch.Clear()
		whiteSet.update(batch)
		blackSet.update(batch)
		updatePieces = false
	}
}

func initSet(pieceFrames [2][6]pixel.Rect, tilePos [][]pixel.Vec, black bool) pieceSet {
	var f int
	var s int
	var ps int
	var set pieceSet

	if black {
		f = 0
		s = 0
		ps = 1
	} else {
		f = 1
		s = 7
		ps = 6
	}
	pawnFrme := pieceFrames[f][5]
	rookFrme := pieceFrames[f][4]
	knightFrme := pieceFrames[f][3]
	bishopFrme := pieceFrames[f][2]
	queenFrme := pieceFrames[f][1]
	kingFrme := pieceFrames[f][0]

	set = pieceSet{
		pieces: []piece{
			{
				pos:         tilePos[ps][0],
				frame:       pawnFrme,
				class:       0,
				tileNumUp:   ps,
				tileNumSide: 0,
				black:       black,
			},
			{
				pos:         tilePos[ps][1],
				frame:       pawnFrme,
				class:       0,
				tileNumUp:   ps,
				tileNumSide: 1,
				black:       black,
			},
			{
				pos:         tilePos[ps][2],
				frame:       pawnFrme,
				class:       0,
				tileNumUp:   ps,
				tileNumSide: 2,
				black:       black,
			},
			{
				pos:         tilePos[ps][3],
				frame:       pawnFrme,
				class:       0,
				tileNumUp:   ps,
				tileNumSide: 3,
				black:       black,
			},
			{
				pos:         tilePos[ps][4],
				frame:       pawnFrme,
				class:       0,
				tileNumUp:   ps,
				tileNumSide: 4,
				black:       black,
			},
			{
				pos:         tilePos[ps][5],
				frame:       pawnFrme,
				class:       0,
				tileNumUp:   ps,
				tileNumSide: 5,
				black:       black,
			},
			{
				pos:         tilePos[ps][6],
				frame:       pawnFrme,
				class:       0,
				tileNumUp:   ps,
				tileNumSide: 6,
				black:       black,
			},
			{
				pos:         tilePos[ps][7],
				frame:       pawnFrme,
				class:       0,
				tileNumUp:   ps,
				tileNumSide: 7,
				black:       black,
			},
			{
				pos:         tilePos[s][1],
				frame:       knightFrme,
				class:       1,
				tileNumUp:   s,
				tileNumSide: 1,
				black:       black,
			},
			{
				pos:         tilePos[s][6],
				frame:       knightFrme,
				class:       1,
				tileNumUp:   s,
				tileNumSide: 6,
				black:       black,
			},
			{
				pos:         tilePos[s][0],
				frame:       rookFrme,
				class:       2,
				tileNumUp:   s,
				tileNumSide: 0,
				black:       black,
			},
			{
				pos:         tilePos[s][7],
				frame:       rookFrme,
				class:       2,
				tileNumUp:   s,
				tileNumSide: 7,
				black:       black,
			},
			{
				pos:         tilePos[s][2],
				frame:       bishopFrme,
				class:       3,
				tileNumUp:   s,
				tileNumSide: 2,
				black:       black,
			},
			{
				pos:         tilePos[s][5],
				frame:       bishopFrme,
				class:       3,
				tileNumUp:   s,
				tileNumSide: 5,
				black:       black,
			},
			{
				pos:         tilePos[s][4],
				frame:       queenFrme,
				class:       4,
				tileNumUp:   s,
				tileNumSide: 4,
				black:       black,
			},
			{
				pos:         tilePos[s][3],
				frame:       kingFrme,
				class:       5,
				tileNumUp:   s,
				tileNumSide: 3,
				black:       black,
			},
		},
	}
	return set
}
