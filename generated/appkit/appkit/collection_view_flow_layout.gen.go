// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [CollectionViewFlowLayout] class.
var CollectionViewFlowLayoutClass objc.Class

func init() {
	CollectionViewFlowLayoutClass = objc.GetClass("NSCollectionViewFlowLayout")
}

type CollectionViewFlowLayout struct {
	objc.ID
}

func CollectionViewFlowLayoutFrom(ptr unsafe.Pointer) CollectionViewFlowLayout {
	return CollectionViewFlowLayout{
		ID: objc.ID(ptr),
	}
}




