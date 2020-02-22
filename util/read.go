package util

import (
	"bufio"
	"io"
	"os"

	"github.com/pkg/errors"
)

// ReadBytesFromStdin reads bytes from STDIN
func ReadBytesFromStdin() ([]byte, error) {
	var (
		data   []byte
		reader = bufio.NewReader(os.Stdin)
	)

	for {
		buf, err := reader.ReadBytes('\n')
		if err != nil && err != io.EOF {
			return nil, err
		}

		data = append(data, buf...)

		if err == io.EOF {
			break
		}
	}

	if len(data) == 0 {
		return nil, errors.New("empty content")
	}

	return data, nil
}
