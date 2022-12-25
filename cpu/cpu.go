package main

const CHIP8_WIDTH = 64
const CHIP8_HEIGHT = 32
const CHIP8_RAM = 4096;

type CPU struct {
  Vram [CHIP8_HEIGHT][CHIP8_WIDH]uint8
  VramChanged bool
  Stack [16]uint
  Ram [CHIP8_RAM]uint8
  V [16]uint8
  I uint
  Pc uint
  Sp uint
  DelayTimer uint8
  SoundTimer uint8
  Keypad [16] bool
  KeypadWaiting bool
  KeypadRegister uint
}
