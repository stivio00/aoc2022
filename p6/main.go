package main

import (
	"fmt"
	"os"
)

const (
	pageSize            uint64 = 1024 * 4 // standard page 4Kb
	circularBufferSize  uint8  = 4
	circularBufferSize2 uint8  = 14
)

type CircularBuffer struct {
	_data       []byte
	_currentPos uint16
	_size       uint8
}

// Creates a new circular buffer.
func NewCB(size uint8) *CircularBuffer {
	return &CircularBuffer{
		_data:       make([]byte, int(size)),
		_currentPos: uint16(size - uint8(1)),
		_size:       size,
	}
}

func main() {
	var buffer []byte = make([]byte, pageSize)
	var cbuffer CircularBuffer = *NewCB(circularBufferSize)
	var pos int = 0 //file cursor

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err.Error())
	}

	_, err = file.Read(buffer)
	if err != nil {
		panic(err.Error())
	}

	// part 1
	for pos < len(buffer) {
		pos += 1
		if !cbuffer.Append(buffer[pos]) {
			continue
		}

		if cbuffer.CheckIfAllAreDiff() {
			fmt.Println("Part 1 Found at ", pos+1, "With value:", string(buffer[pos-3:pos+1]))
			break
		}
	}

	// part 2
	//pos = 0
	cbuffer = *NewCB(circularBufferSize2)
	for pos < len(buffer) {
		pos += 1
		if !cbuffer.Append(buffer[pos]) {
			continue
		}

		if cbuffer.CheckIfAllAreDiff() {
			fmt.Println("Part 2 Found at ", pos+1, "With value:", string(cbuffer._data))
			break
		}

	}

}

// Insert a new byte into the circular buffer.
// Returns true if the last byte is different as from the inserted byte.
func (c *CircularBuffer) Append(newByte byte) bool {
	lastPost := c._currentPos
	c._currentPos += 1
	c._currentPos %= uint16(c._size)
	c._data[c._currentPos] = newByte
	return c._data[lastPost] != c._data[c._currentPos]
}

// Returns true is all the bytes are different.
func (c *CircularBuffer) CheckIfAllAreDiff() bool {
	//TODO implement loop

	for i := uint8(0); i < c._size; i++ {
		for j := i + 1; j < c._size; j++ {
			if c._data[i] == c._data[j] {
				return false
			}
		}
	}

	return true
}
