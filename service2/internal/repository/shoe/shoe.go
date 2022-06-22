package shoe

import "zainokta/twirpexample/service2/internal/entity"

type Shoe interface {
	Make(int32) entity.Shoe
}

type interactor struct{}

func New() Shoe {
	return &interactor{}
}

func (i interactor) Make(size int32) entity.Shoe {
	return entity.Shoe{
		Inches: size,
		Color:  "blue",
		Name:   "shoe",
	}
}
