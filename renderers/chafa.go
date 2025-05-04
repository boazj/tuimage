package renderers

import "image"

type Chafa struct{}

func (c *Chafa) Name() string {
	return "chafa"
}

func (c *Chafa) ImageShow(image *Image) (image.Rectangle, error) {
}

func (c *Chafa) ImageErase(image *Image) error {
}
