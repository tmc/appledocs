// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TextSelectionNavigation] class.
var textSelectionNavigationClass = _TextSelectionNavigationClass{objc.GetClass("NSTextSelectionNavigation")}

type _TextSelectionNavigationClass struct {
	class objc.Class
}

// An interface you use to expose methods for obtaining results from actions performed on text selections. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation

type TextSelectionNavigation struct {
	objectivec.Object
}

// TextSelectionNavigationFrom constructs a [TextSelectionNavigation] from an unsafe.Pointer.
//
// An interface you use to expose methods for obtaining results from actions performed on text selections.
func TextSelectionNavigationFrom(ptr unsafe.Pointer) TextSelectionNavigation {
	return TextSelectionNavigation{objectivec.Object{objc.ID(ptr)}}
}



