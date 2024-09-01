package model

type ModelArgs struct {
	Name   string
	Port   string
	Prompt string
	Temp   string
}

type ModelParams struct {
	Name   string
	Port   string
	Prompt string
	Temp   float64
}
