// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [CollectionViewLayoutInvalidationContext] class.
var CollectionViewLayoutInvalidationContextClass objc.Class

func init() {
	CollectionViewLayoutInvalidationContextClass = objc.GetClass("NSCollectionViewLayoutInvalidationContext")
}

type CollectionViewLayoutInvalidationContext struct {
	objc.ID
}

func CollectionViewLayoutInvalidationContextFrom(ptr unsafe.Pointer) CollectionViewLayoutInvalidationContext {
	return CollectionViewLayoutInvalidationContext{
		ID: objc.ID(ptr),
	}
}



