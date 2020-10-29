package piece

import "fmt"

type Piece interface {
fmt.Stringer
Coord(n int)(int, error)

}