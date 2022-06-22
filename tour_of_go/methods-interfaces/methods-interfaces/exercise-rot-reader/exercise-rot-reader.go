package main

import (
	"io"
	"os"
	"strings"
	"unicode"
)

type rot13Reader struct {
	r io.Reader
}

const (
	alphabet      = "abcdefghijklmnopqrstuvwxyz"
	encryptionKey = 13
	table         = alphabet + alphabet
)

func (r *rot13Reader) Read(p []byte) (int, error) {
	n, err := r.r.Read(p)
	if err != nil {
		return n, err
	}

	for i := range p {
		p[i] = r.rotate(p[i], encryptionKey)
	}
	return n, nil
}

func (r *rot13Reader) rotate(b byte, key int) byte {
	ch := rune(b)
	if !unicode.IsLetter(ch) {
		return b
	}

	key = key % len(alphabet)

	isUpper := false
	if unicode.IsUpper(ch) {
		isUpper = true
		ch = unicode.ToLower(ch)
	}

	idx := byte(ch) - 'a' + byte(key)
	rotatedCh := rune(table[idx])
	if isUpper {
		rotatedCh = unicode.ToUpper(rotatedCh)
	}
	return byte(rotatedCh)
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
