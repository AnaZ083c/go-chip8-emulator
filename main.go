package main

import (
	"chip8/chip8sys"
	"time"

	sdl "github.com/veandco/go-sdl2/sdl"
)

const (
	WIDTH  = 800
	HEIGHT = 600
)

func main() {
	chip8, _ := chip8sys.NewChip8()

	// Init SDL2
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	// Init Window
	window, err := sdl.CreateWindow(
		"CHIP-8 Emulator",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		WIDTH, HEIGHT,
		sdl.WINDOW_SHOWN,
	)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	// Main loop
	running := true
	for running {
		cycleStart := time.Now()

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case sdl.QuitEvent:
				println("Quit")
				running = false
				break
			}
		}

		// Execute one CHIP-8 instruction
		chip8.Execute()

		elapsed := time.Since(cycleStart)
		if elapsed < chip8sys.TIME_PER_INSTRUCTION {
			time.Sleep(chip8sys.TIME_PER_INSTRUCTION - elapsed)
		}

		if chip8.GetDelayTimer() > 0 {
			chip8.SetDelayTimer(chip8.GetDelayTimer() - 1)
		}

		if chip8.GetSoundTimer() > 0 {
			chip8.SetSoundTimer(chip8.GetSoundTimer() - 1)
		}
	}
}
