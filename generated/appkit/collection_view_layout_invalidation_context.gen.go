// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CollectionViewLayoutInvalidationContext] class.
var (
	collectionViewLayoutInvalidationContextClass     _CollectionViewLayoutInvalidationContextClass
	collectionViewLayoutInvalidationContextClassOnce sync.Once
)

func getCollectionViewLayoutInvalidationContextClass() _CollectionViewLayoutInvalidationContextClass {
	collectionViewLayoutInvalidationContextClassOnce.Do(func() {
		collectionViewLayoutInvalidationContextClass = _CollectionViewLayoutInvalidationContextClass{objc.GetClass("NSCollectionViewLayoutInvalidationContext")}
	})
	return collectionViewLayoutInvalidationContextClass
}

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

// Alloc allocates a new instance without initialization.
func (cc _CollectionViewLayoutInvalidationContextClass) Alloc() CollectionViewLayoutInvalidationContext {
	rv := objc.Send[CollectionViewLayoutInvalidationContext](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CollectionViewLayoutInvalidationContextClass) New() CollectionViewLayoutInvalidationContext {
	rv := objc.Send[CollectionViewLayoutInvalidationContext](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CollectionViewLayoutInvalidationContext) Init() CollectionViewLayoutInvalidationContext {
	rv := objc.Send[CollectionViewLayoutInvalidationContext](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CollectionViewLayoutInvalidationContext) Autorelease() CollectionViewLayoutInvalidationContext {
	rv := objc.Send[CollectionViewLayoutInvalidationContext](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCollectionViewLayoutInvalidationContext creates a new CollectionViewLayoutInvalidationContext instance.
func NewCollectionViewLayoutInvalidationContext() CollectionViewLayoutInvalidationContext {
	return getCollectionViewLayoutInvalidationContextClass().New()
}




