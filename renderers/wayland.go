package renderers

import "image"

type Wayland struct{}

func (w *Wayland) Name() string {
}

func (w *Wayland) ImageShow(image *Image) (image.Rectangle, error) {
}

func (w *Wayland) ImageErase(image *Image) error {
}
