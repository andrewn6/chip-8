package display

type Display struct {
  buf [64 * 32]byte
}

func NewDisplay() *Display {
  return &Display{}
}

func (d *Display) GetIndFromPos(x, y int) int { 
  return y*64 + x
}

func (d *Display) DrawByte(byte byte, x, y byte) bool {
  erasing := false
  coordX := int(x) % 64
  coordY := int(y) % 32

  b := byte

  for i := 0; i < 8; i++ {
    index := d.GetIndFromPos(coordX, coordY)
    bit := (b & 128) >> 7
    prevValue := d.buf[index]
    d.buf[index] ^= bit

    if prevValue == 1 && d.buf[index] == 0 {
      erasing = true
    }

    coordX++
    b <<= 1

  }

    return erasing
}

func (d *Display) GetDisBuf() []byte {
  return d.buf[:]
}

func (d *Display) Clear() {
  for i := range d.buf { 
    d.buf[i] = 0
  }
}
