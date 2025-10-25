// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVCaptureSynchronizedData */


/* debug [class_header]: Header for AVCaptureSynchronizedData */
// The class instance for the [CaptureSynchronizedData] class.
var (
	CaptureSynchronizedDataClass     _CaptureSynchronizedDataClass
	CaptureSynchronizedDataClassOnce sync.Once
)

func getCaptureSynchronizedDataClass() _CaptureSynchronizedDataClass {
	CaptureSynchronizedDataClassOnce.Do(func() {
		CaptureSynchronizedDataClass = _CaptureSynchronizedDataClass{objc.GetClass("AVCaptureSynchronizedData")}
	})
	return CaptureSynchronizedDataClass
}

type _CaptureSynchronizedDataClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CaptureSynchronizedData */
// An interface definition for the [CaptureSynchronizedData] class.
type ICaptureSynchronizedData interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CaptureSynchronizedData */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CaptureSynchronizedData */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CaptureSynchronizedData */
// Alloc allocates a new instance without initialization.
func (cc _CaptureSynchronizedDataClass) Alloc() CaptureSynchronizedData {
	rv := objc.Send[CaptureSynchronizedData](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CaptureSynchronizedDataClass) New() CaptureSynchronizedData {
	rv := objc.Send[CaptureSynchronizedData](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureSynchronizedData) Init() CaptureSynchronizedData {
	rv := objc.Send[CaptureSynchronizedData](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureSynchronizedData) Autorelease() CaptureSynchronizedData {
	rv := objc.Send[CaptureSynchronizedData](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureSynchronizedData creates a new CaptureSynchronizedData instance.
func NewCaptureSynchronizedData() CaptureSynchronizedData {
	return getCaptureSynchronizedDataClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CaptureSynchronizedData */
// The abstract superclass for media samples collected using synchronized capture.


// The abstract superclass for media samples collected using synchronized capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSynchronizedData
type CaptureSynchronizedData struct {
	objectivec.Object
}

// CaptureSynchronizedDataFrom constructs a [CaptureSynchronizedData] from an unsafe.Pointer.
//
// The abstract superclass for media samples collected using synchronized capture.
func CaptureSynchronizedDataFrom(ptr unsafe.Pointer) CaptureSynchronizedData {
	return CaptureSynchronizedData{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CaptureSynchronizedData *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CaptureSynchronizedData */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CaptureSynchronizedData */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CaptureSynchronizedData */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CaptureSynchronizedData */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVCaptureSynchronizedData */


