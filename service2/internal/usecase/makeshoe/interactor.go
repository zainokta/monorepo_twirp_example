package makeshoe

import (
	"context"
	"zainokta/twirpexample/service2/internal/repository/shoe"
	"zainokta/twirpexample/service2/rpc/servicetwo"
)

type interactor struct {
	shoeRepository shoe.Shoe
}

func New(shoeRepository shoe.Shoe) Inport {
	return &interactor{shoeRepository: shoeRepository}
}

func (i interactor) MakeHat(ctx context.Context, size *servicetwo.Size) (*servicetwo.Shoe, error) {
	hat := i.shoeRepository.Make(size.GetInches())

	return &servicetwo.Shoe{
		Inches: hat.Inches,
		Color:  hat.Color,
		Name:   hat.Name,
	}, nil
}
