package makeshoe

import (
	"context"
	"zainokta/twirpexample/service2/rpc/servicetwo"
)

type Inport interface {
	MakeHat(context.Context, *servicetwo.Size) (*servicetwo.Shoe, error)
}
