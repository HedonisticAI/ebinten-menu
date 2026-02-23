package keyboard

import (
	types "ebinten-menus/types"
)

type KeyboardButton struct {
	MenuID types.ID
	Active bool
	Prev   *KeyboardButton
	Next   *KeyboardButton
}

func (K *KeyboardButton) Add(Key *KeyboardButton) {
	if K.Next == nil {
		K.Next = Key
		Key.Prev = K
	} else {
		K.Next.Add(Key)
	}
}

func (E *KeyboardButton) SetActive(Active bool) {
	E.Active = Active
}

func (E *KeyboardButton) IsActive() bool {
	return E.Active
}
