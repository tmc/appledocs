// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MutableFontCollection] class.
var mutableFontCollectionClass = _MutableFontCollectionClass{objc.GetClass("NSMutableFontCollection")}

type _MutableFontCollectionClass struct {
	class objc.Class
}

// An interface definition for the [MutableFontCollection] class.
type IMutableFontCollection interface {
	IFontCollection
}

// A mutable collection of font descriptors taken together as a single object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableFontCollection

type MutableFontCollection struct {
	FontCollection
}

// MutableFontCollectionFrom constructs a [MutableFontCollection] from an unsafe.Pointer.
//
// A mutable collection of font descriptors taken together as a single object.
func MutableFontCollectionFrom(ptr unsafe.Pointer) MutableFontCollection {
	return MutableFontCollection{
		FontCollection: FontCollectionFrom(ptr),
	}
}



