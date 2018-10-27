package parse

import (
	"bytes"
	"io/ioutil"
)

// Files returns the content as slice of bytes or an empty slice of bytes on error.
func Files(filenames ...string) []byte {
	var out bytes.Buffer
	for _, file := range filenames {
		raw, err := ioutil.ReadFile(file)
		if err != nil {
			return []byte{}
		}
		out.Write(raw)
	}
	return out.Bytes()
}
