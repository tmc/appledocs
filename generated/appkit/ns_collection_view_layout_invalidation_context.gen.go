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
	CollectionViewLayoutInvalidationContextClass     _CollectionViewLayoutInvalidationContextClass
	CollectionViewLayoutInvalidationContextClassOnce sync.Once
)

func getCollectionViewLayoutInvalidationContextClass() _CollectionViewLayoutInvalidationContextClass {
	CollectionViewLayoutInvalidationContextClassOnce.Do(func() {
		CollectionViewLayoutInvalidationContextClass = _CollectionViewLayoutInvalidationContextClass{objc.GetClass("NSCollectionViewLayoutInvalidationContext")}
	})
	return CollectionViewLayoutInvalidationContextClass
}

type _CollectionViewLayoutInvalidationContextClass struct {
	class objc.Class
}

// An interface definition for the [CollectionViewLayoutInvalidationContext] class.
type ICollectionViewLayoutInvalidationContext interface {
	objectivec.IObject
}

// An object that identifies the portions of your layout that need to be updated.
//
// Invalidation contexts are a way to improve the efficiency of layout operations and must be supported explicitly by the layout object. Instead of invalidating the entire layout, you can create an invalidation layout object that specifies only the portions of the layout that changed. You then pass that invalidation context to the method of the layout object. Typically, you ask the layout object to create an invalidation context for you. The class defines methods for creating a supported invalidation context. If you define a custom layout, you can define additional methods for creating invalidation contexts with custom information. Layout objects may also create invalidation contexts in response to specific changes. For example, layout objects automatically create invalidation contexts when you change the collection view’s data source, when you insert or delete items, and when you reload the collection view’s data.
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
