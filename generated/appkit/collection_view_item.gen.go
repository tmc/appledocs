// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CollectionViewItem] class.
var collectionViewItemClass = _CollectionViewItemClass{objc.GetClass("NSCollectionViewItem")}

type _CollectionViewItemClass struct {
	class objc.Class
}

// The visual representation for a single data element in a collection view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewItem

type CollectionViewItem struct {
	ViewController
}

// CollectionViewItemFrom constructs a [CollectionViewItem] from an unsafe.Pointer.
//
// The visual representation for a single data element in a collection view.
func CollectionViewItemFrom(ptr unsafe.Pointer) CollectionViewItem {
	return CollectionViewItem{
		ViewController: ViewControllerFrom(ptr),
	}
}



