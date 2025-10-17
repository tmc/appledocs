// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CollectionViewLayoutInvalidationContext] class.
var collectionViewLayoutInvalidationContextClass = _CollectionViewLayoutInvalidationContextClass{objc.GetClass("NSCollectionViewLayoutInvalidationContext")}

type _CollectionViewLayoutInvalidationContextClass struct {
	class objc.Class
}

// An interface definition for the [CollectionViewLayoutInvalidationContext] class.
type ICollectionViewLayoutInvalidationContext interface {
	objectivec.IObject
}

// An object that identifies the portions of your layout that need to be updated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutInvalidationContext

type CollectionViewLayoutInvalidationContext struct {
	objectivec.Object
}

// CollectionViewLayoutInvalidationContextFrom constructs a [CollectionViewLayoutInvalidationContext] from an unsafe.Pointer.
//
// An object that identifies the portions of your layout that need to be updated.
func CollectionViewLayoutInvalidationContextFrom(ptr unsafe.Pointer) CollectionViewLayoutInvalidationContext {
	return CollectionViewLayoutInvalidationContext{objectivec.Object{objc.ID(ptr)}}
}



