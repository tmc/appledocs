// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TextLayoutManager] class.
var textLayoutManagerClass = _TextLayoutManagerClass{objc.GetClass("NSTextLayoutManager")}

type _TextLayoutManagerClass struct {
	class objc.Class
}

// An interface definition for the [TextLayoutManager] class.
type ITextLayoutManager interface {
	objectivec.IObject
}

// The primary class that you use to manage text layout and presentation for custom text displays. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager

type TextLayoutManager struct {
	objectivec.Object
}

// TextLayoutManagerFrom constructs a [TextLayoutManager] from an unsafe.Pointer.
//
// The primary class that you use to manage text layout and presentation for custom text displays.
func TextLayoutManagerFrom(ptr unsafe.Pointer) TextLayoutManager {
	return TextLayoutManager{objectivec.Object{objc.ID(ptr)}}
}



