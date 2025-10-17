// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [CollectionViewItem] class.
var CollectionViewItemClass objc.Class

func init() {
	CollectionViewItemClass = objc.GetClass("NSCollectionViewItem")
}

type CollectionViewItem struct {
	objc.ID
}

func CollectionViewItemFrom(ptr unsafe.Pointer) CollectionViewItem {
	return CollectionViewItem{
		ID: objc.ID(ptr),
	}
}



