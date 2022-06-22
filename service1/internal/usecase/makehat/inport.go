package makehat

import (
	"context"
	"zainokta/twirpexample/service1/rpc/serviceone"
)

type Inport interface {
	MakeHat(context.Context, *serviceone.Size) (*serviceone.Hat, error)
}
