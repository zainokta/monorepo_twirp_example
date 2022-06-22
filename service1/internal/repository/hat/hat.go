package hat

import "zainokta/twirpexample/service1/internal/entity"

type Hat interface {
	Make(int32) entity.Hat
}

type interactor struct{}

func New() Hat {
	return &interactor{}
}

func (i interactor) Make(size int32) entity.Hat {
	return entity.Hat{
		Inches: size,
		Color:  "blue",
		Name:   "hat",
	}
}
