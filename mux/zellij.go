package mux

type Zellij struct{}

func NewZellij() *Zellij {
	return &Zellij{}
}

func (z *Zellij) Name() string {
	return "zellij"
}

// Returns the equivilant of TERM, TERM_PROGRAM hiding behind the MUX
// TERM, TERM_PROGRAM, error
func (z *Zellij) Probe() (string, string, error) {
	// TODO: needed in zellij?
	return "", "", nil
}
