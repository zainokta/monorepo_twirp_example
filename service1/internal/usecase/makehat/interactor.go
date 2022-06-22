package makehat

import (
	"context"
	"zainokta/twirpexample/service1/internal/repository/hat"
	"zainokta/twirpexample/service1/rpc/serviceone"
)

type interactor struct {
	hatRepository hat.Hat
}

func New(hatRepository hat.Hat) Inport {
	return &interactor{hatRepository: hatRepository}
}

func (i interactor) MakeHat(ctx context.Context, size *serviceone.Size) (*serviceone.Hat, error) {
	hat := i.hatRepository.Make(size.GetInches())

	return &serviceone.Hat{
		Inches: hat.Inches,
		Color:  hat.Color,
		Name:   hat.Name,
	}, nil
}
