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

func initPieces() {
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

	whiteSet = initSet(pieceFrames, false)
	blackSet = initSet(pieceFrames, true)

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

func initSet(pieceFrames [2][6]pixel.Rect, black bool) pieceSet {
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
				pos:         board[ps][0].pos,
				frame:       pawnFrme,
				class:       0,
				tileNumUp:   ps,
				tileNumSide: 0,
				black:       black,
			},
			{
				pos:         board[ps][1].pos,
				frame:       pawnFrme,
				class:       0,
				tileNumUp:   ps,
				tileNumSide: 1,
				black:       black,
			},
			{
				pos:         board[ps][2].pos,
				frame:       pawnFrme,
				class:       0,
				tileNumUp:   ps,
				tileNumSide: 2,
				black:       black,
			},
			{
				pos:         board[ps][3].pos,
				frame:       pawnFrme,
				class:       0,
				tileNumUp:   ps,
				tileNumSide: 3,
				black:       black,
			},
			{
				pos:         board[ps][4].pos,
				frame:       pawnFrme,
				class:       0,
				tileNumUp:   ps,
				tileNumSide: 4,
				black:       black,
			},
			{
				pos:         board[ps][5].pos,
				frame:       pawnFrme,
				class:       0,
				tileNumUp:   ps,
				tileNumSide: 5,
				black:       black,
			},
			{
				pos:         board[ps][6].pos,
				frame:       pawnFrme,
				class:       0,
				tileNumUp:   ps,
				tileNumSide: 6,
				black:       black,
			},
			{
				pos:         board[ps][7].pos,
				frame:       pawnFrme,
				class:       0,
				tileNumUp:   ps,
				tileNumSide: 7,
				black:       black,
			},
			{
				pos:         board[s][1].pos,
				frame:       knightFrme,
				class:       1,
				tileNumUp:   s,
				tileNumSide: 1,
				black:       black,
			},
			{
				pos:         board[s][6].pos,
				frame:       knightFrme,
				class:       1,
				tileNumUp:   s,
				tileNumSide: 6,
				black:       black,
			},
			{
				pos:         board[s][0].pos,
				frame:       rookFrme,
				class:       2,
				tileNumUp:   s,
				tileNumSide: 0,
				black:       black,
			},
			{
				pos:         board[s][7].pos,
				frame:       rookFrme,
				class:       2,
				tileNumUp:   s,
				tileNumSide: 7,
				black:       black,
			},
			{
				pos:         board[s][2].pos,
				frame:       bishopFrme,
				class:       3,
				tileNumUp:   s,
				tileNumSide: 2,
				black:       black,
			},
			{
				pos:         board[s][5].pos,
				frame:       bishopFrme,
				class:       3,
				tileNumUp:   s,
				tileNumSide: 5,
				black:       black,
			},
			{
				pos:         board[s][4].pos,
				frame:       queenFrme,
				class:       4,
				tileNumUp:   s,
				tileNumSide: 4,
				black:       black,
			},
			{
				pos:         board[s][3].pos,
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
