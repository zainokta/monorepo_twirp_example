package internal

import "zainokta/twirpexample/service1/internal/repository/hat"

type Container struct {
	HatRepository hat.Hat
}

func NewContainer() *Container {
	return &Container{
		HatRepository: hat.New(),
	}
}
