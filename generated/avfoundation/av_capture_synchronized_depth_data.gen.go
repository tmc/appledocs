// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class AVCaptureSynchronizedDepthData */


/* debug [class_header]: Header for AVCaptureSynchronizedDepthData */
// The class instance for the [CaptureSynchronizedDepthData] class.
var (
	CaptureSynchronizedDepthDataClass     _CaptureSynchronizedDepthDataClass
	CaptureSynchronizedDepthDataClassOnce sync.Once
)

func getCaptureSynchronizedDepthDataClass() _CaptureSynchronizedDepthDataClass {
	CaptureSynchronizedDepthDataClassOnce.Do(func() {
		CaptureSynchronizedDepthDataClass = _CaptureSynchronizedDepthDataClass{objc.GetClass("AVCaptureSynchronizedDepthData")}
	})
	return CaptureSynchronizedDepthDataClass
}

type _CaptureSynchronizedDepthDataClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CaptureSynchronizedDepthData */
// An interface definition for the [CaptureSynchronizedDepthData] class.
type ICaptureSynchronizedDepthData interface {
	ICaptureSynchronizedData
	
/* debug [class_interface_properties]: Properties for CaptureSynchronizedDepthData */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CaptureSynchronizedDepthData */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CaptureSynchronizedDepthData */
// Alloc allocates a new instance without initialization.
func (cc _CaptureSynchronizedDepthDataClass) Alloc() CaptureSynchronizedDepthData {
	rv := objc.Send[CaptureSynchronizedDepthData](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CaptureSynchronizedDepthDataClass) New() CaptureSynchronizedDepthData {
	rv := objc.Send[CaptureSynchronizedDepthData](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureSynchronizedDepthData) Init() CaptureSynchronizedDepthData {
	rv := objc.Send[CaptureSynchronizedDepthData](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureSynchronizedDepthData) Autorelease() CaptureSynchronizedDepthData {
	rv := objc.Send[CaptureSynchronizedDepthData](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureSynchronizedDepthData creates a new CaptureSynchronizedDepthData instance.
func NewCaptureSynchronizedDepthData() CaptureSynchronizedDepthData {
	return getCaptureSynchronizedDepthDataClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CaptureSynchronizedDepthData */
// A container for scene depth information collected using synchronized capture.


// A container for scene depth information collected using synchronized capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSynchronizedDepthData
type CaptureSynchronizedDepthData struct {
	CaptureSynchronizedData
}

// CaptureSynchronizedDepthDataFrom constructs a [CaptureSynchronizedDepthData] from an unsafe.Pointer.
//
// A container for scene depth information collected using synchronized capture.
func CaptureSynchronizedDepthDataFrom(ptr unsafe.Pointer) CaptureSynchronizedDepthData {
	return CaptureSynchronizedDepthData{
		CaptureSynchronizedData: CaptureSynchronizedDataFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CaptureSynchronizedDepthData *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CaptureSynchronizedDepthData */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CaptureSynchronizedDepthData */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CaptureSynchronizedDepthData */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CaptureSynchronizedDepthData */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVCaptureSynchronizedDepthData */


