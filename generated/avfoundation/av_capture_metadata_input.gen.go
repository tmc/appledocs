// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVCaptureMetadataInput */


/* debug [class_header]: Header for AVCaptureMetadataInput */
// The class instance for the [CaptureMetadataInput] class.
var (
	CaptureMetadataInputClass     _CaptureMetadataInputClass
	CaptureMetadataInputClassOnce sync.Once
)

func getCaptureMetadataInputClass() _CaptureMetadataInputClass {
	CaptureMetadataInputClassOnce.Do(func() {
		CaptureMetadataInputClass = _CaptureMetadataInputClass{objc.GetClass("AVCaptureMetadataInput")}
	})
	return CaptureMetadataInputClass
}

type _CaptureMetadataInputClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CaptureMetadataInput */
// An interface definition for the [CaptureMetadataInput] class.
type ICaptureMetadataInput interface {
	ICaptureInput
	
/* debug [class_interface_properties]: Properties for CaptureMetadataInput */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CaptureMetadataInput */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CaptureMetadataInput */
// Alloc allocates a new instance without initialization.
func (cc _CaptureMetadataInputClass) Alloc() CaptureMetadataInput {
	rv := objc.Send[CaptureMetadataInput](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CaptureMetadataInputClass) New() CaptureMetadataInput {
	rv := objc.Send[CaptureMetadataInput](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureMetadataInput) Init() CaptureMetadataInput {
	rv := objc.Send[CaptureMetadataInput](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureMetadataInput) Autorelease() CaptureMetadataInput {
	rv := objc.Send[CaptureMetadataInput](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureMetadataInput creates a new CaptureMetadataInput instance.
func NewCaptureMetadataInput() CaptureMetadataInput {
	return getCaptureMetadataInputClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CaptureMetadataInput */
// A capture input for providing timed metadata to a capture session.
//
// This class provides input to an . An instance of can present one and only one connected to an . Provide metadata through the input port by conforming to a and supplying objects in an .


// A capture input for providing timed metadata to a capture session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureMetadataInput
type CaptureMetadataInput struct {
	CaptureInput
}

// CaptureMetadataInputFrom constructs a [CaptureMetadataInput] from an unsafe.Pointer.
//
// A capture input for providing timed metadata to a capture session.
func CaptureMetadataInputFrom(ptr unsafe.Pointer) CaptureMetadataInput {
	return CaptureMetadataInput{
		CaptureInput: CaptureInputFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CaptureMetadataInput */

// Creates capture metadata input to provide timed groups to a capture session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureMetadataInput/init(formatDescription:clock:)
func NewCaptureMetadataInputWithFormatDescriptionClock(desc MetadataFormatDescriptionRef /* not a class type */, clock ClockRef /* not a class type */) CaptureMetadataInput {
	instance := getCaptureMetadataInputClass().Alloc()
	rv := objc.Send[CaptureMetadataInput](instance.ID, objc.Sel("initWithFormatDescription:clock:"), desc, clock)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCaptureMetadataInputWithFormatDescriptionClock */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CaptureMetadataInput */

// Returns a metadata input instance that allows clients to provide timed metadata groups to a capture session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureMetadataInput/metadataInputWithFormatDescription:clock:
func (cc _CaptureMetadataInputClass) MetadataInputWithFormatDescriptionClock(desc MetadataFormatDescriptionRef /* not a class type */, clock ClockRef /* not a class type */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("metadataInputWithFormatDescription:clock:"), desc, clock)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=MetadataInputWithFormatDescriptionClock) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CaptureMetadataInput */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CaptureMetadataInput */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CaptureMetadataInput */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVCaptureMetadataInput */


