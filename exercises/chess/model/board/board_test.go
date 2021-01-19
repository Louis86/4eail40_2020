package board

import (
	"reflect"
	"testing"

	"github.com/jglouis/4eail40_2020/exercises/chess/model/coord"
	"github.com/jglouis/4eail40_2020/exercises/chess/model/piece"
	"github.com/jglouis/4eail40_2020/exercises/chess/model/player"
)

// mockPiece piece pour le test
type mock struct {
}

func (r mock) String() string {
	panic("implement me") //TODO: create function
}

func (r mock) Moves() map[coord.ChessCoordinates]bool {
	panic("implement me") //TODO: create function
}
func (r mock) Color() player.Color {
	panic("implement me") //TODO: create function
}

func TestClassic_MovePiece(t *testing.T) {
	class := Classic{}
	p := mock{}
	p1 := mock{}

	coordin := coord.NewCartesian(2, 1)
	x, _ := coordin.Coord(0)
	y, _ := coordin.Coord(1)
	class[x][y] = p

	coordin1 := coord.NewCartesian(4, 5)
	x1, _ := coordin1.Coord(0)
	y1, _ := coordin1.Coord(1)
	class[x1][y1] = p1

	type args struct {
		from coord.ChessCoordinates
		to   coord.ChessCoordinates
	}

	tests := []struct {
		name    string
		c       Classic
		args    args
		wantErr bool // true if there is an error, false if there is not.
	}{
		{
			"no piece to move",
			class,
			args{from: coord.NewCartesian(2, 1), to: coordin},
			true,
		},
		{
			"occupied at to",
			class,
			args{from: coordin1, to: coordin},
			true,
		},
		{
			"move accepted",
			class,
			args{from: coordin, to: coord.NewCartesian(3, 2)},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.MovePiece(tt.args.from, tt.args.to); (err != nil) != tt.wantErr {
				t.Errorf("Classic.MovePiece() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClassic_PieceAt(t *testing.T) {
	class := Classic{}
	p := mock{}
	coordin := coord.NewCartesian(7, 0)
	x, _ := coordin.Coord(0)
	y, _ := coordin.Coord(1)
	class[x][y] = p
	type args struct {
		at coord.ChessCoordinates
	}
	tests := []struct {
		name string
		c    Classic
		args args
		want piece.Piece
	}{
		{
			"test H0",
			class,
			args{at: coordin},
			p,
		},
		{
			"test C3",
			class,
			args{at: coord.NewCartesian(2, 3)},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.PieceAt(tt.args.at); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PieceAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClassic_PlacePieceAt(t *testing.T) {
	class := Classic{}
	p := mock{}
	coordin := coord.NewCartesian(7, 0)
	x, _ := coordin.Coord(0)
	y, _ := coordin.Coord(1)
	class[x][y] = p
	type args struct {
		p  piece.Piece
		at coord.ChessCoordinates
	}
	tests := []struct {
		name    string
		c       Classic
		args    args
		wantErr bool // true if there is an error, false if there is not.
	}{
		{
			"Test pass",
			class,
			args{p: p, at: coord.NewCartesian(3, 2)},
			false,
		},
		{
			"Test failed",
			class,
			args{p: p, at: coordin},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.PlacePieceAt(tt.args.p, tt.args.at); (err != nil) != tt.wantErr {
				t.Errorf("PlacePieceAt() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
