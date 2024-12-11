package chip8sys

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	WIDTH         = 64
	HEIGHT        = 32
	START_ADDRESS = 0x200

	FONT_ADDRESS_START_INDEX = 0x050

	CHIP8_INSTRUCTIONS_PER_SECOND = 700
	TIME_PER_INSTRUCTION          = time.Second / CHIP8_INSTRUCTIONS_PER_SECOND
)

var keyMap = map[sdl.Keycode]byte{
	sdl.K_1: 0x1,
	sdl.K_2: 0x2,
	sdl.K_3: 0x3,
	sdl.K_4: 0xC,
	sdl.K_q: 0x4,
	sdl.K_w: 0x5,
	sdl.K_e: 0x6,
	sdl.K_r: 0xD,
	sdl.K_a: 0x7,
	sdl.K_s: 0x8,
	sdl.K_d: 0x9,
	sdl.K_f: 0xE,
	sdl.K_z: 0xA,
	sdl.K_x: 0x0,
	sdl.K_c: 0xB,
	sdl.K_v: 0xF,
}

type Chip8 struct {
	memory     [4069]byte      // 4kiB of memory
	display    [64 * 32]uint32 // 16b sprite width * 16b sprite height -> 32b uint
	pc         uint16          // program counter - current index in memory
	I          uint16          // stack pointer - current stack
	stack      [16]uint16      // functions
	delayTimer byte            // decrement by 1, 60 times per second
	soundTimer byte            // decrement by 1, 60 times per second - should beep when above 0
	registers  [16]byte        // registers for other use
}

func GetFont() []byte {
	return []byte{
		0xF0, 0x90, 0x90, 0x90, 0xF0, // 0
		0x20, 0x60, 0x20, 0x20, 0x70, // 1
		0xF0, 0x10, 0xF0, 0x80, 0xF0, // 2
		0xF0, 0x10, 0xF0, 0x10, 0xF0, // 3
		0x90, 0x90, 0xF0, 0x10, 0x10, // 4
		0xF0, 0x80, 0xF0, 0x10, 0xF0, // 5
		0xF0, 0x80, 0xF0, 0x90, 0xF0, // 6
		0xF0, 0x10, 0x20, 0x40, 0x40, // 7
		0xF0, 0x90, 0xF0, 0x90, 0xF0, // 8
		0xF0, 0x90, 0xF0, 0x10, 0xF0, // 9
		0xF0, 0x90, 0xF0, 0x90, 0x90, // A
		0xE0, 0x90, 0xE0, 0x90, 0xE0, // B
		0xF0, 0x80, 0x80, 0x80, 0xF0, // C
		0xE0, 0x90, 0x90, 0x90, 0xE0, // D
		0xF0, 0x80, 0xF0, 0x80, 0xF0, // E
		0xF0, 0x80, 0xF0, 0x80, 0x80, // F
	}
}

func NewChip8() (*Chip8, error) {
	chip8 := &Chip8{
		delayTimer: 255,
		soundTimer: 255,
		pc:         START_ADDRESS,
	}
	font := GetFont()

	for i := 0; i < len(font); i++ {
		chip8.memory[i+FONT_ADDRESS_START_INDEX] = font[i] // fill memory with font from 0x050 to 0x1FF (index 80 to 511)
	}

	return chip8, nil
}

func (c *Chip8) GetMemory() []byte {
	return c.memory[:]
}

func (c *Chip8) GetDelayTimer() byte {
	return c.delayTimer
}

func (c *Chip8) GetSoundTimer() byte {
	return c.soundTimer
}

func (c *Chip8) SetDelayTimer(newValue byte) {
	c.delayTimer = newValue
}

func (c *Chip8) SetSoundTimer(newValue byte) {
	c.soundTimer = newValue
}

func (c *Chip8) Load(chip8file string) {
	// Load a program into memory
	file, err := os.Open(chip8file)
	if err != nil {
		panic(err)
	}

	// Get an array of bytes from the file
	br := bufio.NewReader(file)
	readBytes := []byte{}
	for {
		_byte, err := br.ReadByte()
		if err != nil && !errors.Is(err, io.EOF) {
			fmt.Println(err)
			break
		}
		readBytes = append(readBytes, _byte)

		if err != nil {
			break
		}
	}

	for i := 0; i < len(readBytes); i++ {
		c.memory[START_ADDRESS+i] = readBytes[i]
	}
}

func (c *Chip8) Fetch() uint16 {
	// instruction from memory at the current PC (program counter)
	instruction := uint16(c.memory[c.pc])<<8 | uint16(c.memory[c.pc+1])
	c.pc += 2

	return instruction
}

func (c *Chip8) Execute() {
	// the instruction and do what it tells you
	// TODO
	fmt.Println("Execute: not yet implemented")
}
