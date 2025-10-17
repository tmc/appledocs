// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CollectionViewLayout] class.
var collectionViewLayoutClass = _CollectionViewLayoutClass{objc.GetClass("NSCollectionViewLayout")}

type _CollectionViewLayoutClass struct {
	class objc.Class
}

// An abstract base class that you subclass and use to generate layout information for a collection view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout

type CollectionViewLayout struct {
	objectivec.Object
}

// CollectionViewLayoutFrom constructs a [CollectionViewLayout] from an unsafe.Pointer.
//
// An abstract base class that you subclass and use to generate layout information for a collection view.
func CollectionViewLayoutFrom(ptr unsafe.Pointer) CollectionViewLayout {
	return CollectionViewLayout{objectivec.Object{objc.ID(ptr)}}
}



