package console

import "image/color"

// ScrollDir defines a scroll direction.
type ScrollDir uint8

// The supported list of scroll directions for the console Scroll() calls.
const (
	ScrollDirUp ScrollDir = iota
	ScrollDirDown
)

// The Device interface is implemented by objects that can function as system
// consoles.
type Device interface {
	// Dimensions returns the width and height of the console in characters.
	Dimensions() (uint32, uint32)

	// DefaultColors returns the default foreground and background colors
	// used by this console.
	DefaultColors() (fg, bg uint8)

	// Fill sets the contents of the specified rectangular region to the
	// requested color. Both x and y coordinates are 1-based (top-left
	// corner has coordinates 1,1).
	Fill(x, y, width, height uint32, fg, bg uint8)

	// Scroll the console contents to the specified direction. The caller
	// is responsible for updating (e.g. clear or replace) the contents of
	// the region that was scrolled.
	Scroll(dir ScrollDir, lines uint32)

	// Write a char to the specified location. Both x and y coordinates are
	// 1-based (top-left corner has coordinates 1,1).
	Write(ch byte, fg, bg uint8, x, y uint32)

	// Palette returns the active color palette for this console.
	Palette() color.Palette

	// SetPaletteColor updates the color definition for the specified
	// palette index. Passing a color index greated than the number of
	// supported colors should be a no-op.
	SetPaletteColor(uint8, color.RGBA)
}
