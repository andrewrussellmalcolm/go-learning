package main

import (
	"io"
	"os"
	"strings"

)

type Rot13Reader struct {
	reader io.Reader

	i int
}


func (r Rot13Reader) Read(b []byte) (int, error) {

	c:=make([]byte,len(b))
	n,err := r.reader.Read(c)

	if err == io.EOF {
		return 0, io.EOF
	} else {
		return n, err
	}
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := Rot13Reader{s,100}
	io.Copy(os.Stdout, &r)
}
