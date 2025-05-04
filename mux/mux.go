package mux

type Mux interface {
	Name() string
	Probe() (string, string, error)
}
