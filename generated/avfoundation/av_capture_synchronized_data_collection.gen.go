// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVCaptureSynchronizedDataCollection */


/* debug [class_header]: Header for AVCaptureSynchronizedDataCollection */
// The class instance for the [CaptureSynchronizedDataCollection] class.
var (
	CaptureSynchronizedDataCollectionClass     _CaptureSynchronizedDataCollectionClass
	CaptureSynchronizedDataCollectionClassOnce sync.Once
)

func getCaptureSynchronizedDataCollectionClass() _CaptureSynchronizedDataCollectionClass {
	CaptureSynchronizedDataCollectionClassOnce.Do(func() {
		CaptureSynchronizedDataCollectionClass = _CaptureSynchronizedDataCollectionClass{objc.GetClass("AVCaptureSynchronizedDataCollection")}
	})
	return CaptureSynchronizedDataCollectionClass
}

type _CaptureSynchronizedDataCollectionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CaptureSynchronizedDataCollection */
// An interface definition for the [CaptureSynchronizedDataCollection] class.
type ICaptureSynchronizedDataCollection interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CaptureSynchronizedDataCollection */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CaptureSynchronizedDataCollection */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CaptureSynchronizedDataCollection */
// Alloc allocates a new instance without initialization.
func (cc _CaptureSynchronizedDataCollectionClass) Alloc() CaptureSynchronizedDataCollection {
	rv := objc.Send[CaptureSynchronizedDataCollection](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CaptureSynchronizedDataCollectionClass) New() CaptureSynchronizedDataCollection {
	rv := objc.Send[CaptureSynchronizedDataCollection](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureSynchronizedDataCollection) Init() CaptureSynchronizedDataCollection {
	rv := objc.Send[CaptureSynchronizedDataCollection](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureSynchronizedDataCollection) Autorelease() CaptureSynchronizedDataCollection {
	rv := objc.Send[CaptureSynchronizedDataCollection](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureSynchronizedDataCollection creates a new CaptureSynchronizedDataCollection instance.
func NewCaptureSynchronizedDataCollection() CaptureSynchronizedDataCollection {
	return getCaptureSynchronizedDataCollectionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CaptureSynchronizedDataCollection */
// A set of data samples collected simultaneously from multiple capture outputs.


// A set of data samples collected simultaneously from multiple capture outputs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSynchronizedDataCollection
type CaptureSynchronizedDataCollection struct {
	objectivec.Object
}

// CaptureSynchronizedDataCollectionFrom constructs a [CaptureSynchronizedDataCollection] from an unsafe.Pointer.
//
// A set of data samples collected simultaneously from multiple capture outputs.
func CaptureSynchronizedDataCollectionFrom(ptr unsafe.Pointer) CaptureSynchronizedDataCollection {
	return CaptureSynchronizedDataCollection{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CaptureSynchronizedDataCollection *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CaptureSynchronizedDataCollection */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CaptureSynchronizedDataCollection */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CaptureSynchronizedDataCollection */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CaptureSynchronizedDataCollection */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVCaptureSynchronizedDataCollection */


