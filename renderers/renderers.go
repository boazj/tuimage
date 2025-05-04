package renderers

import (
	"image"
)

type Renderer interface {
	Name() string

	ImageShow(image *Image) (image.Rectangle, error)
	ImageErase(image *Image) error
}

var (
	Kgp    Renderer
	KgpOld Renderer
	Iip    Renderer
	Sixel  Renderer

	// Supported by Überzug++
	X11     Renderer
	Wayland Renderer
	Chafa   Renderer
)
