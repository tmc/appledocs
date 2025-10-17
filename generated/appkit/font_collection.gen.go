// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FontCollection] class.
var fontCollectionClass = _FontCollectionClass{objc.GetClass("NSFontCollection")}

type _FontCollectionClass struct {
	class objc.Class
}

// A font collection, which is a group of font descriptors taken together as a single object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontCollection

type FontCollection struct {
	objectivec.Object
}

// FontCollectionFrom constructs a [FontCollection] from an unsafe.Pointer.
//
// A font collection, which is a group of font descriptors taken together as a single object.
func FontCollectionFrom(ptr unsafe.Pointer) FontCollection {
	return FontCollection{objectivec.Object{objc.ID(ptr)}}
}



