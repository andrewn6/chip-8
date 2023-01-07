package chip

import (
  "github.com/anddddrew/chip-8/display"
  "github.com/anddddrew/chip-8/keyboard"
  "github.com/anddddrew/chip-8/timer"
)

type Chip8 struct {
  mem[4096]uint8
  vx [16]uint8
  stack []uint16
  i uint16
  pc uint16
  // TODO: add timer, keyboard, display.
}

func NewChip8() *Chip8 {
  res := &Chip8 {
    mem: [4096]uint8{},
  }
  return res
}
