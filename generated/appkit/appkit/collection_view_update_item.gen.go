// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [CollectionViewUpdateItem] class.
var CollectionViewUpdateItemClass objc.Class

func init() {
	CollectionViewUpdateItemClass = objc.GetClass("NSCollectionViewUpdateItem")
}

type CollectionViewUpdateItem struct {
	objc.ID
}

func CollectionViewUpdateItemFrom(ptr unsafe.Pointer) CollectionViewUpdateItem {
	return CollectionViewUpdateItem{
		ID: objc.ID(ptr),
	}
}



