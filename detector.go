package tuimage

import (
	"fmt"
	"os"
	"strings"

	"github.com/boazj/tuimage/mux"
	"github.com/boazj/tuimage/renderers"
)

type (
	Terminal string
)

const (
	None      Terminal = "NONE"
	Kitty     Terminal = "kitty"     // https://sw.kovidgoyal.net/kitty/
	Konsole   Terminal = "konsole"   // https://konsole.kde.org/
	Iterm2    Terminal = "iterm2"    // https://iterm2.com/
	WezTerm   Terminal = "wezterm"   // https://wezterm.org/
	Foot      Terminal = "foot"      // https://codeberg.org/dnkl/foot
	Ghostty   Terminal = "ghostty"   // https://ghostty.org/
	Microsoft Terminal = "microsoft" // TODO:
	Warp      Terminal = "wrap"      // https://www.warp.dev/
	Rio       Terminal = "rio"       // https://rioterm.com/
	BlackBox  Terminal = "blackbox"  // https://github.com/yonasBSD/blackbox-terminal
	VSCode    Terminal = "vscode"    // TODO:
	Tabby     Terminal = "tabby"     // https://tabby.sh/
	Hyper     Terminal = "hyper"     // https://hyper.is/
	Mintty    Terminal = "mintty"    // https://mintty.github.io/
	VTerm     Terminal = "libvterm"  // TODO:
	Apple     Terminal = "apple"     // TODO:
	Urxvt     Terminal = "urxvt"     // TODO:
	Bobcat    Terminal = "bobcat"    // TODO:
)

var (
	NoneMux mux.Mux = nil
	Tmux    mux.Mux = mux.NewTmux()   // https://github.com/tmux/tmux/wiki
	Zellij  mux.Mux = mux.NewZellij() // https://zellij.dev/
)

var terminalCsiHints = map[string]Terminal{
	"kitty":    Kitty,
	"Konsole":  Konsole,
	"iTerm2":   Iterm2,
	"WezTerm":  WezTerm,
	"foot":     Foot,
	"ghostty":  Ghostty,
	"libvterm": VTerm,
	"Bobcat":   Bobcat,
}

var muxCsiHinds = map[string]mux.Mux{
	"tmux ": Tmux,
}

var terminalEnvHints = map[string]Terminal{
	"KITTY_WINDOW_ID":        Kitty,
	"KONSOLE_VERSION":        Konsole,
	"ITERM_SESSION_ID":       Iterm2,
	"WEZTERM_EXECUTABLE":     WezTerm,
	"GHOSTTY_RESOURCES_DIR":  Ghostty,
	"WT_Session":             Microsoft,
	"WARP_HONOR_PS1":         Warp,
	"VSCODE_INJECTION":       VSCode,
	"TABBY_CONFIG_DIRECTORY": Tabby,
}

var muxEnvHints = map[string]mux.Mux{
	"TMUX":                Tmux,
	"TMUX_PANE":           Tmux,
	"ZELLIJ":              Zellij,
	"ZELLIJ_SESSION_NAME": Zellij,
}

var terminalEnvTermHints = map[string]Terminal{
	"xterm-kitty":           Kitty,
	"foot":                  Foot,
	"foot-extra":            Foot,
	"xterm-ghostty":         Ghostty,
	"rio":                   Rio,
	"rxvt-unicode-256color": Urxvt,
}

var terminalEnvProgramHints = map[string]Terminal{
	"iTerm.app":      Iterm2,
	"WezTerm":        WezTerm,
	"ghostty":        Ghostty,
	"WarpTerminal":   Warp,
	"rio":            Rio,
	"BlackBox":       BlackBox,
	"vscode":         VSCode,
	"Tabby":          Tabby,
	"Hyper":          Hyper,
	"mintty":         Mintty,
	"Apple_Terminal": Apple,
}

var muxEnvProgramHints = map[string]mux.Mux{
	"tmux": Tmux,
}

// TODO: needed? don't think so
var DefaultMuxRenderers = map[mux.Mux][]renderers.Renderer{
	Tmux:   {},
	Zellij: {},
}

var DefaultRenderers = map[Terminal][]renderers.Renderer{
	Kitty:     {renderers.Kgp},
	Konsole:   {renderers.KgpOld},
	Iterm2:    {renderers.Iip, renderers.Sixel},
	WezTerm:   {renderers.Iip, renderers.Sixel},
	Foot:      {renderers.Sixel},
	Ghostty:   {renderers.Kgp},
	Microsoft: {renderers.Sixel},
	Warp:      {renderers.Iip, renderers.KgpOld},
	Rio:       {renderers.Iip, renderers.Sixel},
	BlackBox:  {renderers.Sixel},
	VSCode:    {renderers.Iip, renderers.Sixel},
	Tabby:     {renderers.Iip, renderers.Sixel},
	Hyper:     {renderers.Iip, renderers.Sixel},
	Mintty:    {renderers.Iip},
	VTerm:     {},
	Apple:     {},
	Urxvt:     {},
	Bobcat:    {renderers.Iip, renderers.Sixel},
}

func muxFromEnv(term string, program string) mux.Mux {
	mux, ok := muxEnvProgramHints[term]
	if !ok {
		for k, v := range muxEnvHints {
			if isEnvExists(k) {
				mux = v
				ok = true
				break
			}
		}
	}
	if ok {
		return mux
	}
	return NoneMux
}

func FromCsi(csi string) (Terminal, error) {
	for k, v := range terminalCsiHints {
		if strings.Contains(csi, k) {
			return v, nil
		}
	}
	return "", fmt.Errorf("cant recognize terminal emulator via csi")
}

// Decides on the Emulator and Multiplexer brands using (mostly) the env data
// Can diverge from env data when uncovering data behind the active Multiplexer
func FromEnv() (Terminal, mux.Mux, error) {
	term := getEnvTerm()
	program := getEnvProgram()

	// probe mux to get underlying emu
	mux := muxFromEnv(term, program)
	if mux != NoneMux {
		t, p, err := mux.Probe()
		term = t
		program = p
		if err != nil {
			// TODO: should a failure to probe means a return? or just continue?
			return None, NoneMux, err
		}
	}

	if emu, ok := terminalEnvTermHints[term]; ok {
		return emu, mux, nil
	}

	if emu, ok := terminalEnvProgramHints[program]; ok {
		return emu, mux, nil
	}

	for k, v := range terminalEnvHints {
		if isEnvExists(k) {
			return v, mux, nil
		}
	}

	return None, NoneMux, fmt.Errorf("cant recognize terminal emulator via env")
}

func getEnvTerm() string {
	term, ok := os.LookupEnv("OVERRIDE_TERM")
	if !ok {
		term = os.Getenv("TERM")
	}
	return term
}

func getEnvProgram() string {
	program, ok := os.LookupEnv("OVERRIDE_TERM_PROGRAM")
	if !ok {
		program = os.Getenv("TERM_PROGRAM")
	}
	return program
}

// TODO: util?
func isEnvExists(name string) bool {
	_, exists := os.LookupEnv(name)
	return exists
}
