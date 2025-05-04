package tuimage

import (
	"image"
	"os"
	"runtime"
	"strings"

	"github.com/boazj/tuimage/mux"
	"github.com/boazj/tuimage/renderers"
	"github.com/tiendc/gofn"
)

type Provider struct {
	os       string
	wsl      bool
	brand    Terminal
	mux      mux.Mux
	renderer renderers.Renderer
	// TODO: lock
}

func NewProvider() (*Provider, error) {
	brand, mux, err := FromEnv()
	if err != nil {
		return nil, err // TODO: error handling
	}
	if brand == None {
		// TODO: csi process
	}
	e := &Provider{
		os:    runtime.GOOS,
		wsl:   isWsl(),
		mux:   mux,
		brand: brand,
	}
	renderer, err := e.infer()
	if err != nil {
		return nil, err // TODO: error handling
	}
	e.renderer = renderer
	return e, nil
}

func (p *Provider) Image(path string, maxArea image.Rectangle) *renderers.Image {
	return renderers.NewImage(path, maxArea, p.renderer)
}

func isWsl() bool {
	if runtime.GOOS == "linux" {
		_, err := os.Lstat("/proc/sys/fs/binfmt_misc/WSLInterop")
		if err == nil {
			return true
		}
	}
	return false
}

func (p *Provider) infer() (renderers.Renderer, error) {
	if p.brand == Microsoft {
		// TODO: strange, windows -> IIP but windows terminal -> sixel?
		return renderers.Sixel, nil
	}
	if p.wsl && p.brand == WezTerm {
		// TODO: aggressive, wez supports more then that
		return renderers.KgpOld, nil
	}

	supported, ok := DefaultRenderers[p.brand]
	if !ok {
		// TODO: need to infer support from CSI
		return nil, nil
	}
	if p.os == "windows" {
		// TODO: move to windows
		supported = gofn.FilterIN(supported, renderers.Iip)
	}
	switch p.mux {
	case Zellij:
		// TODO: move to zellij
		supported = gofn.FilterIN(supported, renderers.Sixel)
	case Tmux:
		// TODO: move to tmux
		supported = gofn.FilterIN(supported, renderers.KgpOld)
	}
	if len(supported) > 0 {
		return supported[0], nil
	}

	// fallback to ueberzug
	hasWaylandDompositor := isEnvExists("SWAYSOCK") || isEnvExists("HYPRLAND_INSTANCE_SIGNATURE") || isEnvExists("WAYFIRE_SOCKET")
	session, ok := os.LookupEnv("XDG_SESSION_TYPE")
	if ok {
		switch {
		case session == "x11":
			return renderers.X11, nil
		case session == "wayland" && hasWaylandDompositor:
			return renderers.Wayland, nil
		default:
			return renderers.Chafa, nil
		}
	}
	if isEnvExists("WAYLAND_DISPLAY") && hasWaylandDompositor {
		return renderers.Wayland, nil
	}

	display, ok := os.LookupEnv("DISPLAY")
	if ok && display != "" && strings.Contains(display, "/org.xquartz") {
		return renderers.X11, nil
	}

	// fallback to chafa
	return renderers.Chafa, nil
}
