package board

import (
	"github.com/jglouis/4eail40_2020/exercises/chess/model/coord"
	"github.com/jglouis/4eail40_2020/exercises/chess/model/piece"
)

// Classic 8x8 Chess board
type Classic [8][8]piece.Piece

func (c *Classic) String() string {
	panic("not implemented") // TODO: Implement

}

// PieceAt retrievves piece at give coordinates.
// Returns nil if no piece was found.
func (c *Classic) PieceAt(at coord.ChessCoordinates) piece.Piece {
	//panic("not implemented") // TODO: Implement

		x, _ := at.Coord(0)
		y, _ := at.Coord(1)
		piece := c[x][y]

		if x =< 8 && y <= 8 {
			return piece
		} else {
			return nil
		}
}

// MovePiece moves a piece from given coordinates to
// given coordinates.
// Returns an error if destination was occupied.
func (c *Classic) MovePiece(from coord.ChessCoordinates, to coord.ChessCoordinates) error {
	//panic("not implemented") // TODO: Implement
	if c.PieceAt(from) != nil {
		x, _ := from.Coord(0)
		y, _ := from.Coord(1)
		p := c[x][y]
		return c.PlacePieceAt(p, to)

	} else {
		return fmt.Errorf("No piece ")
	}

}

// PlacePieceAt places a given piece at given location.
// Returns an error if destination was occupied.
func (c *Classic) PlacePieceAt(p piece.Piece, at coord.ChessCoordinates) error {

		if c.PieceAt(at) == nil {
		x, _ := at.Coord(0)
		y, _ := at.Coord(1)
		c[x][y] = p
		return nil
		}
		if c.PieceAt(at) != nil{
			return fmt.Errorf("occupied")
		}

}
