package lang

import (
	"fmt"
	"os"
)

func readElf(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	fi, err := f.Stat()
	if err != nil {
		return nil, err
	}

	len := fi.Size()

	buf := make([]byte, len)

	rlen, err := f.Read(buf)
	if err != nil {
		return nil, err
	}

	if rlen != int(len) {
		return nil, fmt.Errorf("readElf: read %d bytes, expected %d", rlen, len)
	}
	return buf, nil
}
