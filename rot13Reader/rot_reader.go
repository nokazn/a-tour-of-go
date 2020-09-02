package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(b byte) byte {
	isUpperCase := 'A' <= b && b <= 'Z'
	isLowerCase := 'a' <= b && b <= 'z'
	if isUpperCase || isLowerCase {
		result := b + 13
		if (isUpperCase && result > 'Z') || (isLowerCase && result > 'z') {
			return result - 26
		}
		return result
	}

	return b
}

func (reader rot13Reader) Read(bytes []byte) (int, error) {
	n, err := reader.r.Read(bytes)
	for i, b := range bytes {
		bytes[i] = rot13(b)
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!\n")
	r := rot13Reader{s}

	io.Copy(os.Stdout, &r)
}
