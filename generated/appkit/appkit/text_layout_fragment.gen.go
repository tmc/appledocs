// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TextLayoutFragment] class.
var TextLayoutFragmentClass objc.Class

func init() {
	TextLayoutFragmentClass = objc.GetClass("NSTextLayoutFragment")
}

type TextLayoutFragment struct {
	objc.ID
}

func TextLayoutFragmentFrom(ptr unsafe.Pointer) TextLayoutFragment {
	return TextLayoutFragment{
		ID: objc.ID(ptr),
	}
}



