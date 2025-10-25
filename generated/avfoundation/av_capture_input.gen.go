// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVCaptureInput */


/* debug [class_header]: Header for AVCaptureInput */
// The class instance for the [CaptureInput] class.
var (
	CaptureInputClass     _CaptureInputClass
	CaptureInputClassOnce sync.Once
)

func getCaptureInputClass() _CaptureInputClass {
	CaptureInputClassOnce.Do(func() {
		CaptureInputClass = _CaptureInputClass{objc.GetClass("AVCaptureInput")}
	})
	return CaptureInputClass
}

type _CaptureInputClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CaptureInput */
// An interface definition for the [CaptureInput] class.
type ICaptureInput interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CaptureInput */
	// properties:
	Ports() []CaptureInputPort
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CaptureInput */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CaptureInput */
// Alloc allocates a new instance without initialization.
func (cc _CaptureInputClass) Alloc() CaptureInput {
	rv := objc.Send[CaptureInput](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CaptureInputClass) New() CaptureInput {
	rv := objc.Send[CaptureInput](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureInput) Init() CaptureInput {
	rv := objc.Send[CaptureInput](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureInput) Autorelease() CaptureInput {
	rv := objc.Send[CaptureInput](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureInput creates a new CaptureInput instance.
func NewCaptureInput() CaptureInput {
	return getCaptureInputClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CaptureInput */
// An abstract superclass for objects that provide input data to a capture session.
//
// You create concrete instances of this class, such as , to add inputs to a capture session. An input provides one or more streams of media data. For example, input devices can provide both audio and video data. The framework represents each media stream that an input provides as an object. A capture makes connections between capture inputs and capture outputs using a object. The connection defines the mapping between a set of port objects and an .


// An abstract superclass for objects that provide input data to a capture session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureInput
type CaptureInput struct {
	objectivec.Object
}

// CaptureInputFrom constructs a [CaptureInput] from an unsafe.Pointer.
//
// An abstract superclass for objects that provide input data to a capture session.
func CaptureInputFrom(ptr unsafe.Pointer) CaptureInput {
	return CaptureInput{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CaptureInput *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CaptureInput */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CaptureInput */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CaptureInput */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CaptureInput */

// The ports available on a capture input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureInput/ports
func (c_ CaptureInput) Ports() []CaptureInputPort {
	rv := objc.Send[[]CaptureInputPort](c_.ID, objc.Sel("ports"))
	return rv
}/* debug [instance_properties/getter]: ports */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVCaptureInput */



