package imagegenerator

import (
	"fmt"
	"io"
)

func read(buffer io.Reader, len int) ([]byte, error) {
	h := make([]byte, len)
	r, err := buffer.Read(h)
	if err != nil {
		return nil, err
	}
	if r != len {
		return nil, fmt.Errorf("unable to read: %d", r)
	}
	return h, nil
}

func write(w io.Writer, h []byte) error {
	r, err := w.Write(h)
	if err != nil {
		return err
	}
	if r != len(h) {
		return fmt.Errorf("unable to set write: %d", r)
	}
	return nil
}

func fixHeader(w io.Writer, buffer io.Reader) error {
	//   get 54 bytes
	_, err := read(buffer, 54)
	if err != nil {
		return err
	}
	//   write hard coded 62 bytes
	var fixHeader = []byte{
		0x42, 0x4d, 0x3e, 0x44, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x3e, 0x00, 0x00, 0x00, 0x28, 0x00,
		0x00, 0x00, 0xf8, 0x03, 0x00, 0x00, 0x88, 0x02, 0x00, 0x00, 0x01, 0x00, 0x01, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x13, 0x0b, 0x00, 0x00, 0x13, 0x0b, 0x00, 0x00, 0x02, 0x00,
		0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0xff, 0xff,
	}
	err = write(w, fixHeader)
	if err != nil {
		return err
	}
	return nil
}

func setBit(b *byte, bit int) {
	*b |= (0x80 >> (bit % 8))
}

func fixBody(w io.Writer, buffer io.Reader) error {
	offset := 0
	dataExists := true
	for dataExists {
		var b byte
		max := 8
		for j := 0; j < max; j++ {
			//   get 3 bytes(24 bits)
			d, err := read(buffer, 3)
			if err != nil && err.Error() != "EOF" {
				return err
			}
			if err != nil && err.Error() == "EOF" {
				dataExists = false
				break
			}
			//   convert to single bit
			bit := false
			if d[0] > 0xF0 && d[1] > 0xF0 && d[2] > 0xF0 {
				bit = true
			}
			//   append
			if bit {
				setBit(&b, j)
			}
		}
		if !dataExists {
			break
		}
		err := write(w, []byte{b})
		if err != nil {
			return err
		}
		offset++
		if offset == 127 {
			err := write(w, []byte{0x00})
			if err != nil {
				return err
			}
			offset = 0
		}
	}

	return nil
}

func EncodeMonoChrome(w io.Writer, buffer io.Reader) error {
	// replace header
	err := fixHeader(w, buffer)
	if err != nil {
		return err
	}

	// replace body
	err = fixBody(w, buffer)
	if err != nil {
		return err
	}

	return nil
}
