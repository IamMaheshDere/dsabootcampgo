package models

import "context"

type Tree interface {
	Populate(ctx context.Context)
	Display(ctx context.Context)
	PreetyDisplay(ctx context.Context)
}
