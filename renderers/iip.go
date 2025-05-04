package renderers

import "image"

type Iip struct{}

func (i *Iip) Name() string {
	return "iip"
}

func (i *Iip) ImageShow(image *Image) (image.Rectangle, error) {
}

func (i *Iip) ImageErase(image *Image) error {
}
