// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class AVCaptureSynchronizedMetadataObjectData */


/* debug [class_header]: Header for AVCaptureSynchronizedMetadataObjectData */
// The class instance for the [CaptureSynchronizedMetadataObjectData] class.
var (
	CaptureSynchronizedMetadataObjectDataClass     _CaptureSynchronizedMetadataObjectDataClass
	CaptureSynchronizedMetadataObjectDataClassOnce sync.Once
)

func getCaptureSynchronizedMetadataObjectDataClass() _CaptureSynchronizedMetadataObjectDataClass {
	CaptureSynchronizedMetadataObjectDataClassOnce.Do(func() {
		CaptureSynchronizedMetadataObjectDataClass = _CaptureSynchronizedMetadataObjectDataClass{objc.GetClass("AVCaptureSynchronizedMetadataObjectData")}
	})
	return CaptureSynchronizedMetadataObjectDataClass
}

type _CaptureSynchronizedMetadataObjectDataClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CaptureSynchronizedMetadataObjectData */
// An interface definition for the [CaptureSynchronizedMetadataObjectData] class.
type ICaptureSynchronizedMetadataObjectData interface {
	ICaptureSynchronizedData
	
/* debug [class_interface_properties]: Properties for CaptureSynchronizedMetadataObjectData */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CaptureSynchronizedMetadataObjectData */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CaptureSynchronizedMetadataObjectData */
// Alloc allocates a new instance without initialization.
func (cc _CaptureSynchronizedMetadataObjectDataClass) Alloc() CaptureSynchronizedMetadataObjectData {
	rv := objc.Send[CaptureSynchronizedMetadataObjectData](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CaptureSynchronizedMetadataObjectDataClass) New() CaptureSynchronizedMetadataObjectData {
	rv := objc.Send[CaptureSynchronizedMetadataObjectData](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureSynchronizedMetadataObjectData) Init() CaptureSynchronizedMetadataObjectData {
	rv := objc.Send[CaptureSynchronizedMetadataObjectData](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureSynchronizedMetadataObjectData) Autorelease() CaptureSynchronizedMetadataObjectData {
	rv := objc.Send[CaptureSynchronizedMetadataObjectData](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureSynchronizedMetadataObjectData creates a new CaptureSynchronizedMetadataObjectData instance.
func NewCaptureSynchronizedMetadataObjectData() CaptureSynchronizedMetadataObjectData {
	return getCaptureSynchronizedMetadataObjectDataClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CaptureSynchronizedMetadataObjectData */
// A container for metadata objects collected using synchronized capture.


// A container for metadata objects collected using synchronized capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSynchronizedMetadataObjectData
type CaptureSynchronizedMetadataObjectData struct {
	CaptureSynchronizedData
}

// CaptureSynchronizedMetadataObjectDataFrom constructs a [CaptureSynchronizedMetadataObjectData] from an unsafe.Pointer.
//
// A container for metadata objects collected using synchronized capture.
func CaptureSynchronizedMetadataObjectDataFrom(ptr unsafe.Pointer) CaptureSynchronizedMetadataObjectData {
	return CaptureSynchronizedMetadataObjectData{
		CaptureSynchronizedData: CaptureSynchronizedDataFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CaptureSynchronizedMetadataObjectData *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CaptureSynchronizedMetadataObjectData */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CaptureSynchronizedMetadataObjectData */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CaptureSynchronizedMetadataObjectData */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CaptureSynchronizedMetadataObjectData */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVCaptureSynchronizedMetadataObjectData */


