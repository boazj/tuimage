package renderers

import (
	"fmt"
	"image"
)

type RenderState int

const (
	ImageNotRendered RenderState = iota
	ImageRendered
	ImageError
)

type Image struct {
	renderer    Renderer
	renderState RenderState

	path string
	// The allocated area for rendering this image, the process will try and occupy as much of this area as possible
	maxArea image.Rectangle
	// The actual area occupied by the image, might not be equal to the max area due to aspect ratio gaps
	area image.Rectangle

	// Cached downscaled version of the image befitting the allocated area
	downscaled any // TODO: type
}

func NewImage(path string, maxArea image.Rectangle, renderer Renderer) *Image {
	// TODO: should the rendere come from somewhere else?
	// TODO: validate and normalize max area
	// TODO: check file exists
	// TODO: check image type supported
	m := &Image{
		renderer:    renderer,
		path:        path,
		maxArea:     maxArea,
		renderState: ImageNotRendered,
	}
	return m
}

func (m *Image) Render() (image.Rectangle, error) {
	if m.renderState != ImageNotRendered {
		return image.Rectangle{}, fmt.Errorf("image is already rendered, rendering again will leave a ghost image on the screen")
	}
	rect, err := m.renderer.ImageShow(m)
	m.area = rect
	m.renderState = ImageRendered
	if err != nil {
		m.renderState = ImageError
		return image.Rectangle{}, err // TODO: error handling
	}
	return rect, nil
}

func (m *Image) Erase() error {
	if m.renderState == ImageRendered {
		return fmt.Errorf("image is not rendered, erasing the allocated are will have unexpected results")
	}
	err := m.renderer.ImageErase(m)
	if err != nil {
		m.renderState = ImageError
		// TODO: error handling
	} else {
		m.renderState = ImageNotRendered
	}
	return err
}

func (m *Image) IsShown() bool {
	return m.renderState == ImageRendered
}

func (m *Image) IsError() bool {
	return m.renderState == ImageError
}

func (m *Image) GetAllocatedArea() image.Rectangle {
	return m.maxArea
}

func (m *Image) GetRenderedArea() (image.Rectangle, error) {
	if m.IsShown() {
		return m.area, nil
	}
	return image.Rectangle{}, fmt.Errorf("image not rendered and thus has no rendered area")
}

func (m *Image) RendererType() string {
	return m.renderer.Name()
}

func (m *Image) Downscale() {
	// TODO:
}

func (m *Image) MaxPixel(rect image.Rectangle) (uint32, uint32) {
	// TODO:
	return 0, 0
}

func (m *Image) PixelArea(width uint32, height uint32, rect image.Rectangle) image.Rectangle {
	// TODO:
	return image.Rectangle{}
}
