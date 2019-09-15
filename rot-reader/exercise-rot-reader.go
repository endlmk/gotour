package main

import (
	"strings"
	"os"
	"io"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader)Read(b []byte) (int, error) {
	for i := 0; i < len(b); i++ {
		bb := make([]byte, 1)
		_, err := r.r.Read(bb)
		if err == io.EOF {
			break
		}
		c := rune(bb[0])
		var base byte 
		if 'a' <= c && c <= 'z' {
			base = 'a'
			b[i] = ((bb[0] - base) + 13) % 26 + base
		} else if 'A' <= c && c <= 'Z' {
			base = 'A'
			b[i] = ((bb[0] - base) + 13) % 26 + base
		} else {
			b[i] = bb[0]
		}
	}
	return len(b), io.EOF
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}