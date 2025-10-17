// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CollectionViewUpdateItem] class.
var collectionViewUpdateItemClass = _CollectionViewUpdateItemClass{objc.GetClass("NSCollectionViewUpdateItem")}

type _CollectionViewUpdateItemClass struct {
	class objc.Class
}

// An interface definition for the [CollectionViewUpdateItem] class.
type ICollectionViewUpdateItem interface {
	objectivec.IObject
}

// A description of a single change to make to an item in a collection view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewUpdateItem

type CollectionViewUpdateItem struct {
	objectivec.Object
}

// CollectionViewUpdateItemFrom constructs a [CollectionViewUpdateItem] from an unsafe.Pointer.
//
// A description of a single change to make to an item in a collection view.
func CollectionViewUpdateItemFrom(ptr unsafe.Pointer) CollectionViewUpdateItem {
	return CollectionViewUpdateItem{objectivec.Object{objc.ID(ptr)}}
}



