// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSCollectionViewFlowLayoutInvalidationContext */


/* debug [class_header]: Header for NSCollectionViewFlowLayoutInvalidationContext */
// The class instance for the [CollectionViewFlowLayoutInvalidationContext] class.
var (
	CollectionViewFlowLayoutInvalidationContextClass     _CollectionViewFlowLayoutInvalidationContextClass
	CollectionViewFlowLayoutInvalidationContextClassOnce sync.Once
)

func getCollectionViewFlowLayoutInvalidationContextClass() _CollectionViewFlowLayoutInvalidationContextClass {
	CollectionViewFlowLayoutInvalidationContextClassOnce.Do(func() {
		CollectionViewFlowLayoutInvalidationContextClass = _CollectionViewFlowLayoutInvalidationContextClass{objc.GetClass("NSCollectionViewFlowLayoutInvalidationContext")}
	})
	return CollectionViewFlowLayoutInvalidationContextClass
}

type _CollectionViewFlowLayoutInvalidationContextClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CollectionViewFlowLayoutInvalidationContext */
// An interface definition for the [CollectionViewFlowLayoutInvalidationContext] class.
type ICollectionViewFlowLayoutInvalidationContext interface {
	ICollectionViewLayoutInvalidationContext
	
/* debug [class_interface_properties]: Properties for CollectionViewFlowLayoutInvalidationContext */
	// properties:
	InvalidateFlowLayoutAttributes() bool
	SetInvalidateFlowLayoutAttributes(value bool)
	InvalidateFlowLayoutDelegateMetrics() bool
	SetInvalidateFlowLayoutDelegateMetrics(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CollectionViewFlowLayoutInvalidationContext */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CollectionViewFlowLayoutInvalidationContext */
// Alloc allocates a new instance without initialization.
func (cc _CollectionViewFlowLayoutInvalidationContextClass) Alloc() CollectionViewFlowLayoutInvalidationContext {
	rv := objc.Send[CollectionViewFlowLayoutInvalidationContext](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CollectionViewFlowLayoutInvalidationContextClass) New() CollectionViewFlowLayoutInvalidationContext {
	rv := objc.Send[CollectionViewFlowLayoutInvalidationContext](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CollectionViewFlowLayoutInvalidationContext) Init() CollectionViewFlowLayoutInvalidationContext {
	rv := objc.Send[CollectionViewFlowLayoutInvalidationContext](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CollectionViewFlowLayoutInvalidationContext) Autorelease() CollectionViewFlowLayoutInvalidationContext {
	rv := objc.Send[CollectionViewFlowLayoutInvalidationContext](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCollectionViewFlowLayoutInvalidationContext creates a new CollectionViewFlowLayoutInvalidationContext instance.
func NewCollectionViewFlowLayoutInvalidationContext() CollectionViewFlowLayoutInvalidationContext {
	return getCollectionViewFlowLayoutInvalidationContextClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CollectionViewFlowLayoutInvalidationContext */
// An object that identifies the portions of a flow layout object that need to be updated.
//
// Layout objects use invalidation contexts to optimize the layout process and avoid unnecessary work. You use this class to specify whether the object should fetch new size information from its delegate. You can also prevent the flow layout object from updating its layout information altogether. When you want to invalidate your flow layout object, call the method of your layout object and instantiate the resulting class. (The implementation of that method in returns this class.) After instantiating this class, set the properties to appropriate values and pass the object to the method of the layout object.


// An object that identifies the portions of a flow layout object that need to be updated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayoutInvalidationContext
type CollectionViewFlowLayoutInvalidationContext struct {
	CollectionViewLayoutInvalidationContext
}

// CollectionViewFlowLayoutInvalidationContextFrom constructs a [CollectionViewFlowLayoutInvalidationContext] from an unsafe.Pointer.
//
// An object that identifies the portions of a flow layout object that need to be updated.
func CollectionViewFlowLayoutInvalidationContextFrom(ptr unsafe.Pointer) CollectionViewFlowLayoutInvalidationContext {
	return CollectionViewFlowLayoutInvalidationContext{
		CollectionViewLayoutInvalidationContext: CollectionViewLayoutInvalidationContextFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CollectionViewFlowLayoutInvalidationContext *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CollectionViewFlowLayoutInvalidationContext */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CollectionViewFlowLayoutInvalidationContext */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CollectionViewFlowLayoutInvalidationContext */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CollectionViewFlowLayoutInvalidationContext */

// A Boolean value indicating whether the flow layout object should invalidate its current attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayoutInvalidationContext/invalidateFlowLayoutAttributes
func (c_ CollectionViewFlowLayoutInvalidationContext) InvalidateFlowLayoutAttributes() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("invalidateFlowLayoutAttributes"))
	return rv
}/* debug [instance_properties/getter]: invalidateFlowLayoutAttributes */


// A Boolean value indicating whether the flow layout object should invalidate its current attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayoutInvalidationContext/invalidateFlowLayoutAttributes
func (c_ CollectionViewFlowLayoutInvalidationContext) SetInvalidateFlowLayoutAttributes(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setInvalidateFlowLayoutAttributes:"), value)
}/* debug [instance_properties/setter]: invalidateFlowLayoutAttributes */


// A Boolean value indicating whether the flow layout object should fetch new size information from its delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayoutInvalidationContext/invalidateFlowLayoutDelegateMetrics
func (c_ CollectionViewFlowLayoutInvalidationContext) InvalidateFlowLayoutDelegateMetrics() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("invalidateFlowLayoutDelegateMetrics"))
	return rv
}/* debug [instance_properties/getter]: invalidateFlowLayoutDelegateMetrics */


// A Boolean value indicating whether the flow layout object should fetch new size information from its delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayoutInvalidationContext/invalidateFlowLayoutDelegateMetrics
func (c_ CollectionViewFlowLayoutInvalidationContext) SetInvalidateFlowLayoutDelegateMetrics(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setInvalidateFlowLayoutDelegateMetrics:"), value)
}/* debug [instance_properties/setter]: invalidateFlowLayoutDelegateMetrics */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSCollectionViewFlowLayoutInvalidationContext */



