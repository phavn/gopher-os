// Package logo contains logos that can be used with a framebuffer console.
package logo

import "image/color"

// ConsoleLogo defines the logo used by framebuffer consoles. If set to nil
// then no logo will be displayed.
var ConsoleLogo *Image

// Image describes an 8bpp image with
type Image struct {
	// The width and height of the logo in pixels.
	Width  uint32
	Height uint32

	// TransparentIndex defines a color index that will be treated as
	// transparent when drawing the logo.
	TransparentIndex uint8

	// The palette for the logo. The console remaps the palette
	// entries to the end of its own palette.
	Palette []color.RGBA

	// The logo data comprises of Width*Height bytes where each byte
	// represents an index in the logo palette.
	Data []uint8
}
