package keyboard

import (
	"ebinten-menus/controller"
	types "ebinten-menus/types"
)

type keyboard interface {
	Add(Key *KeyboardButton) types.ID
	SetActive(Active controller.Status)
	IsActive() controller.Status
	GetID() types.ID
}

type KeyboardButton struct {
	MenuID types.ID
	Active controller.Status
	Prev   *KeyboardButton
	Next   *KeyboardButton
}

func (K *KeyboardButton) GetID() types.ID {
	return K.MenuID
}

func (K *KeyboardButton) Add(Key *KeyboardButton) types.ID {
	if K.MenuID != Key.MenuID {
		return 0
	}
	if K.Next == nil {
		K.Next = Key
		Key.Prev = K
	} else {
		K.Next.Add(Key)
	}
	return K.MenuID
}

func (E *KeyboardButton) SetActive(Active controller.Status) {
	E.Active = Active
}

func (E *KeyboardButton) IsActive() controller.Status {
	return E.Active
}
