package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

var moving bool
var movingPiece *piece

func updateMovement(win *pixelgl.Window, batch *pixel.Batch) {
	if win.Pressed(pixelgl.MouseButtonLeft) {
		if !moving {
			for i := 0; i < len(board); i++ {
				for j := 0; j < len(board[i]); j++ {
					if board[i][j].occupied &&
						win.MousePosition().X <= board[i][j].max.X &&
						win.MousePosition().Y <= board[i][j].max.Y &&
						win.MousePosition().X >= board[i][j].min.X &&
						win.MousePosition().Y >= board[i][j].min.Y {
						moving = true
						movingPiece = board[i][j].occupint
						board[i][j].occupint = nil
						board[i][j].occupied = false
					}
				}
			}
		} else {
			movingPiece.pos = win.MousePosition()
			movingPiece.moveUpdate(batch)
		}
	} else {
		found := false
		if moving {
			for i := 0; i < len(board); i++ {
				for j := 0; j < len(board[i]); j++ {
					if win.MousePosition().X <= board[i][j].max.X &&
						win.MousePosition().Y <= board[i][j].max.Y &&
						win.MousePosition().X >= board[i][j].min.X &&
						win.MousePosition().Y >= board[i][j].min.Y {
						movingPiece.tryMove(&board[i][j], &board[movingPiece.tileNumUp][movingPiece.tileNumSide], i, j, &found)
					}
				}
			}

			if !found {
				movingPiece.pos = board[movingPiece.tileNumUp][movingPiece.tileNumSide].pos
				board[movingPiece.tileNumUp][movingPiece.tileNumSide].occupied = true
				board[movingPiece.tileNumUp][movingPiece.tileNumSide].occupint = movingPiece
			}

			moving = false
			movingPiece = nil
			updatePieces = true
		}
	}
}

func (p *piece) tryMove(square, oldSquare *square, i, j int, found *bool) {

	if p.class == 0 { //Start pawn rules
		if square.occupied && square.occupint.black != p.black {
			if square.side == oldSquare.side-1 || square.side == oldSquare.side+1 {
				if p.black {
					if square.up == oldSquare.up+1 {
						square.occupint.dead = true
						p.move(square, oldSquare, i, j, found)
					}
				} else {
					if square.up == oldSquare.up-1 {
						square.occupint.dead = true
						p.move(square, oldSquare, i, j, found)
					}
				}
			}
		} else if !square.occupied {
			if p.moves == 0 {
				if p.black {
					if square.side == oldSquare.side && square.up > oldSquare.up && square.up <= oldSquare.up+2 {
						p.move(square, oldSquare, i, j, found)
					}
				} else {
					if square.side == oldSquare.side && square.up < oldSquare.up && square.up >= oldSquare.up-2 {
						p.move(square, oldSquare, i, j, found)
					}
				}
			} else {
				if p.black {
					if square.side == oldSquare.side && square.up > oldSquare.up && square.up <= oldSquare.up+1 {
						p.move(square, oldSquare, i, j, found)
					}
				} else {
					if square.side == oldSquare.side && square.up < oldSquare.up && square.up >= oldSquare.up-1 {
						p.move(square, oldSquare, i, j, found)
					}
				}
			}
		}
	} //End pawn rules

}

func (p *piece) move(square, oldSquare *square, i, j int, found *bool) {
	p.tileNumUp = i
	p.tileNumSide = j
	p.pos = square.pos
	p.moves++
	square.occupied = true
	square.occupint = movingPiece
	*found = true
}

func (set *pieceSet) updatePossibleMoves() {
	for i := 0; i < len(set.pieces); i++ {
		if set.pieces[i].class == 0 { //Pawn check

		}
	}
}

func (p *piece) updatePawnMovers() {
	p.possibleMoves = nil //Resetting possible moves
	if p.moves == 0 {
		if p.black {
			if !board[p.tileNumUp+1][p.tileNumSide].occupied { //Up 1 check
				p.possibleMoves = append(p.possibleMoves, &board[p.tileNumUp+1][p.tileNumSide])
			}

			if !board[p.tileNumUp+2][p.tileNumSide].occupied { //Up 2 check
				p.possibleMoves = append(p.possibleMoves, &board[p.tileNumUp+2][p.tileNumSide])
			}

			if p.tileNumUp != 7 && board[p.tileNumUp+1][p.tileNumSide+1].occupied { //Diag right check

			}
		}

	}
}
