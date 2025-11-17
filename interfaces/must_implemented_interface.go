package main

// Interface -.
type Interface interface {
	Example()
}

type CustomType struct {
}

var _ Interface = (*CustomType)(nil)
