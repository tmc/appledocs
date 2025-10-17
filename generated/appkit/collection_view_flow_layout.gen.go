// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CollectionViewFlowLayout] class.
var collectionViewFlowLayoutClass = _CollectionViewFlowLayoutClass{objc.GetClass("NSCollectionViewFlowLayout")}

type _CollectionViewFlowLayoutClass struct {
	class objc.Class
}

// A layout that organizes items into a flexible and configurable arrangement. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout

type CollectionViewFlowLayout struct {
	CollectionViewLayout
}

// CollectionViewFlowLayoutFrom constructs a [CollectionViewFlowLayout] from an unsafe.Pointer.
//
// A layout that organizes items into a flexible and configurable arrangement.
func CollectionViewFlowLayoutFrom(ptr unsafe.Pointer) CollectionViewFlowLayout {
	return CollectionViewFlowLayout{
		CollectionViewLayout: CollectionViewLayoutFrom(ptr),
	}
}



