// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLCaptureDescriptor */


/* debug [class_header]: Header for MTLCaptureDescriptor */
// The class instance for the [CaptureDescriptor] class.
var (
	CaptureDescriptorClass     _CaptureDescriptorClass
	CaptureDescriptorClassOnce sync.Once
)

func getCaptureDescriptorClass() _CaptureDescriptorClass {
	CaptureDescriptorClassOnce.Do(func() {
		CaptureDescriptorClass = _CaptureDescriptorClass{objc.GetClass("MTLCaptureDescriptor")}
	})
	return CaptureDescriptorClass
}

type _CaptureDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CaptureDescriptor */
// An interface definition for the [CaptureDescriptor] class.
type ICaptureDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CaptureDescriptor */
	// properties:
	CaptureObject() objc.ID
	SetCaptureObject(value objc.ID)
	Destination() CaptureDestination
	SetDestination(value CaptureDestination)
	OutputURL() objc.IObject /* cross-framework: NSURL */
	SetOutputURL(value objc.IObject /* cross-framework: NSURL */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CaptureDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CaptureDescriptor */
// Alloc allocates a new instance without initialization.
func (cc _CaptureDescriptorClass) Alloc() CaptureDescriptor {
	rv := objc.Send[CaptureDescriptor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CaptureDescriptorClass) New() CaptureDescriptor {
	rv := objc.Send[CaptureDescriptor](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureDescriptor) Init() CaptureDescriptor {
	rv := objc.Send[CaptureDescriptor](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureDescriptor) Autorelease() CaptureDescriptor {
	rv := objc.Send[CaptureDescriptor](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureDescriptor creates a new CaptureDescriptor instance.
func NewCaptureDescriptor() CaptureDescriptor {
	return getCaptureDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CaptureDescriptor */
// A configuration for a Metal capture session.


// A configuration for a Metal capture session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCaptureDescriptor
type CaptureDescriptor struct {
	objectivec.Object
}

// CaptureDescriptorFrom constructs a [CaptureDescriptor] from an unsafe.Pointer.
//
// A configuration for a Metal capture session.
func CaptureDescriptorFrom(ptr unsafe.Pointer) CaptureDescriptor {
	return CaptureDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CaptureDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CaptureDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CaptureDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CaptureDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CaptureDescriptor */

// The instance whose contents should be captured.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCaptureDescriptor/captureObject
func (c_ CaptureDescriptor) CaptureObject() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("captureObject"))
	return rv
}/* debug [instance_properties/getter]: captureObject */


// The instance whose contents should be captured.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCaptureDescriptor/captureObject
func (c_ CaptureDescriptor) SetCaptureObject(value objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCaptureObject:"), value)
}/* debug [instance_properties/setter]: captureObject */


// The destination for any captured command data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCaptureDescriptor/destination
func (c_ CaptureDescriptor) Destination() CaptureDestination {
	rv := objc.Send[CaptureDestination](c_.ID, objc.Sel("destination"))
	return rv
}/* debug [instance_properties/getter]: destination */


// The destination for any captured command data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCaptureDescriptor/destination
func (c_ CaptureDescriptor) SetDestination(value CaptureDestination) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDestination:"), value)
}/* debug [instance_properties/setter]: destination */


// A URL for a file to write the capture data into.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCaptureDescriptor/outputURL
func (c_ CaptureDescriptor) OutputURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](c_.ID, objc.Sel("outputURL"))
	return rv
}/* debug [instance_properties/getter]: outputURL */


// A URL for a file to write the capture data into.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCaptureDescriptor/outputURL
func (c_ CaptureDescriptor) SetOutputURL(value objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOutputURL:"), value)
}/* debug [instance_properties/setter]: outputURL */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLCaptureDescriptor */



