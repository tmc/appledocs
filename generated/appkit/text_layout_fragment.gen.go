// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TextLayoutFragment] class.
var textLayoutFragmentClass = _TextLayoutFragmentClass{objc.GetClass("NSTextLayoutFragment")}

type _TextLayoutFragmentClass struct {
	class objc.Class
}

// A class that represents the layout fragment typically corresponding to a rendering surface, such as a layer or view subclass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment

type TextLayoutFragment struct {
	objectivec.Object
}

// TextLayoutFragmentFrom constructs a [TextLayoutFragment] from an unsafe.Pointer.
//
// A class that represents the layout fragment typically corresponding to a rendering surface, such as a layer or view subclass.
func TextLayoutFragmentFrom(ptr unsafe.Pointer) TextLayoutFragment {
	return TextLayoutFragment{objectivec.Object{objc.ID(ptr)}}
}



