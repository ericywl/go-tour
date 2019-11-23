package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(char byte) byte {
	if (char >= 'a' && char <= 'm') || (char >= 'A' && char <= 'M') {
		return char + 13
	}

	if (char >= 'n' && char <= 'z') || (char >= 'N' && char <= 'Z') {
		return char - 13
	}

	return char
}

func (rt rot13Reader) Read(b []byte) (int, error) {
	n, err := rt.r.Read(b)
	for i, char := range b {
		b[i] = rot13(char)
	}

	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
