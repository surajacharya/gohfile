package hfile

import (
	"encoding/binary"
	"fmt"
)

func MockKeyInt(i int) []byte {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(i))
	return buf
}

func MockValueInt(i int) []byte {
	return []byte(fmt.Sprintf("value-for-%d", i))
}

func GenerateMockHfile(path string, keyCount, blockSize int, compress, verbose, progress bool) error {
	w, err := NewLocalWriter(path, compress, blockSize, verbose)
	if err != nil {
		return err
	}
	return WriteMockIntPairs(w, keyCount, progress)
}

func WriteMockIntPairs(w *Writer, keyCount int, progress bool) error {
	for i := 0; i < keyCount; i++ {
		if progress && i%10000 == 0 {
			fmt.Printf("\r %d %.02f%%", i, (float64(i)*100.0)/float64(keyCount))
		}
		if err := w.Write(MockKeyInt(i), MockValueInt(i)); err != nil {
			return err
		}
	}

	if progress {
		fmt.Println()
	}

	w.Close()
	return nil
}
