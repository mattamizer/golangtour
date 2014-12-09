package main

import "code.google.com/p/go-tour/reader"
import "io"

type MyReader struct {
	reader io.Reader
}

// TODO: Figure out why this doesn't do the thing. Probably do this after I'm
//       not on medication.
func (r MyReader) Read(b []byte) (n int, err error) {
	for {
		r.reader.Read(b)
	}
}

func main() {
	reader.Validate(MyReader{})
}
