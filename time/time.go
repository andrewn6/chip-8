package time

import (
  "time"
)

type Timer struct {
  delayTimer uint8
  timerSetTime time.Time
}

func NewTimer() *Timer {
  return &Timer {
    delayTimer: 0,
    timerSetTime: time.Now(),
  }
}

func (t *Timer) SetTimer(value uint8) {
 t.delayTimer = value
 t.timerSetTime = time.Now()
}

func (t *Timer) GetTime() uint8 {
  diff := time.Since(t.timerSetTime)
  ms := diff / time.Millisecond
  if ms / 16 >= uint64(t.delayTimer) {
    return 0
  }
  return t.delayTimer - uint8(ms/16)
}
