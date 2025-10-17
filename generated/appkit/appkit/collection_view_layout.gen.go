// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [CollectionViewLayout] class.
var CollectionViewLayoutClass objc.Class

func init() {
	CollectionViewLayoutClass = objc.GetClass("NSCollectionViewLayout")
}

type CollectionViewLayout struct {
	objc.ID
}

func CollectionViewLayoutFrom(ptr unsafe.Pointer) CollectionViewLayout {
	return CollectionViewLayout{
		ID: objc.ID(ptr),
	}
}




