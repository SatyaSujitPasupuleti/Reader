package main

import "golang.org/x/tour/reader"
import "io"
type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (r MyReader) Read(b []byte) (n int, err error) {
	for i, _ := range b {
		b[i] = 'A'
		if err == io.EOF {
			break
		}
	}

	return len(b), nil

}
func main() {
	reader.Validate(MyReader{})
}