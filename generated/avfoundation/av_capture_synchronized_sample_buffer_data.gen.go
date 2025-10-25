// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class AVCaptureSynchronizedSampleBufferData */


/* debug [class_header]: Header for AVCaptureSynchronizedSampleBufferData */
// The class instance for the [CaptureSynchronizedSampleBufferData] class.
var (
	CaptureSynchronizedSampleBufferDataClass     _CaptureSynchronizedSampleBufferDataClass
	CaptureSynchronizedSampleBufferDataClassOnce sync.Once
)

func getCaptureSynchronizedSampleBufferDataClass() _CaptureSynchronizedSampleBufferDataClass {
	CaptureSynchronizedSampleBufferDataClassOnce.Do(func() {
		CaptureSynchronizedSampleBufferDataClass = _CaptureSynchronizedSampleBufferDataClass{objc.GetClass("AVCaptureSynchronizedSampleBufferData")}
	})
	return CaptureSynchronizedSampleBufferDataClass
}

type _CaptureSynchronizedSampleBufferDataClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CaptureSynchronizedSampleBufferData */
// An interface definition for the [CaptureSynchronizedSampleBufferData] class.
type ICaptureSynchronizedSampleBufferData interface {
	ICaptureSynchronizedData
	
/* debug [class_interface_properties]: Properties for CaptureSynchronizedSampleBufferData */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CaptureSynchronizedSampleBufferData */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CaptureSynchronizedSampleBufferData */
// Alloc allocates a new instance without initialization.
func (cc _CaptureSynchronizedSampleBufferDataClass) Alloc() CaptureSynchronizedSampleBufferData {
	rv := objc.Send[CaptureSynchronizedSampleBufferData](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CaptureSynchronizedSampleBufferDataClass) New() CaptureSynchronizedSampleBufferData {
	rv := objc.Send[CaptureSynchronizedSampleBufferData](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureSynchronizedSampleBufferData) Init() CaptureSynchronizedSampleBufferData {
	rv := objc.Send[CaptureSynchronizedSampleBufferData](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureSynchronizedSampleBufferData) Autorelease() CaptureSynchronizedSampleBufferData {
	rv := objc.Send[CaptureSynchronizedSampleBufferData](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureSynchronizedSampleBufferData creates a new CaptureSynchronizedSampleBufferData instance.
func NewCaptureSynchronizedSampleBufferData() CaptureSynchronizedSampleBufferData {
	return getCaptureSynchronizedSampleBufferDataClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CaptureSynchronizedSampleBufferData */
// A container for video or audio samples collected using synchronized capture.


// A container for video or audio samples collected using synchronized capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSynchronizedSampleBufferData
type CaptureSynchronizedSampleBufferData struct {
	CaptureSynchronizedData
}

// CaptureSynchronizedSampleBufferDataFrom constructs a [CaptureSynchronizedSampleBufferData] from an unsafe.Pointer.
//
// A container for video or audio samples collected using synchronized capture.
func CaptureSynchronizedSampleBufferDataFrom(ptr unsafe.Pointer) CaptureSynchronizedSampleBufferData {
	return CaptureSynchronizedSampleBufferData{
		CaptureSynchronizedData: CaptureSynchronizedDataFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CaptureSynchronizedSampleBufferData *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CaptureSynchronizedSampleBufferData */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CaptureSynchronizedSampleBufferData */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CaptureSynchronizedSampleBufferData */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CaptureSynchronizedSampleBufferData */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVCaptureSynchronizedSampleBufferData */


