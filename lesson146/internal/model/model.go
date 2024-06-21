package model

import "github.com/google/uuid"

type Order struct {
	ID       uuid.UUID
	Table    int
	Category string
	Params   HookahParams
}

type HookahParams struct {
	Strength int
	Taste    string
}
