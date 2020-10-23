package main

import (
	"io"
	"os"
	"strings"
)

type spaceEraser struct {
	r io.Reader
}

func main() {
	s := strings.NewReader("H e l l o w o r l d!")
	r := spaceEraser{s}
	io.Copy(os.Stdout, &r)
}


func (r spaceEraser) Read(p []byte) (int, error) {
	n, err := r.r.Read(p)
	j := 0
	i := 0
	for  i < n  {
		if p[i] != 32 {
			p[j] = p[i]
			j++
		}
		i += 1
	}
	return j, err

}


// Implement a type that satisfies the io.Reader interface and reads from another io.Reader,
// modifying the stream by removing the spaces.
// Expected output: "Helloworld!"
