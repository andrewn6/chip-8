package keyboard

type Keyboard struct {
	keyPressed uint8
}

func NewKeyboard() *Keyboard {
	return &Keyboard{keyPressed: 0}
}

func (k *Keyboard) IsKeyPressed(keyCode uint8) bool {
	if key, ok := k.keyPressed; ok {
		return key == keyCode
	}

	return false
}

func (k *Keyboard) SetKeyPressed(key uint8) {
	k.keyPressed = key
}

func (k *Keyboard) GetKeyPressed() uint8 {
	return k.keyPressed
}
