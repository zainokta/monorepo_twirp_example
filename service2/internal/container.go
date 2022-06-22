package internal

import "zainokta/twirpexample/service2/internal/repository/shoe"

type Container struct {
	ShoeRepository shoe.Shoe
}

func NewContainer() *Container {
	return &Container{
		ShoeRepository: shoe.New(),
	}
}
